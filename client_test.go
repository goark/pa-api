package paapi5

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// stubQuery is a minimal Query implementation for client tests.
type stubQuery struct {
	op      Operation
	payload []byte
	err     error
}

func (s stubQuery) Operation() Operation     { return s.op }
func (s stubQuery) Payload() ([]byte, error) { return s.payload, s.err }

// newServers stands up two httptest.Servers: one acting as the Cognito
// token endpoint, the other as the Creators API. Returns both plus a
// configured Server pointing at them.
func newServers(t *testing.T, tokenHandler, apiHandler http.HandlerFunc) (*httptest.Server, *httptest.Server, *Server) {
	t.Helper()
	tokenSrv := httptest.NewServer(tokenHandler)
	apiSrv := httptest.NewServer(apiHandler)
	t.Cleanup(tokenSrv.Close)
	t.Cleanup(apiSrv.Close)
	apiURL := strings.TrimPrefix(apiSrv.URL, "http://")
	sv := New(
		WithMarketplace(LocaleUnitedStates),
		WithServerScheme("http"),
		WithServerHost(apiURL),
		WithServerAuthEndpoint(tokenSrv.URL),
	)
	return tokenSrv, apiSrv, sv
}

func TestClientBasics(t *testing.T) {
	c := New().CreateClient("mytag-20", "credID", "credSecret")
	if got, want := c.Marketplace(), DefaultMarketplace.String(); got != want {
		t.Errorf("Client.Marketplace() = %q, want %q", got, want)
	}
	if got, want := c.PartnerTag(), "mytag-20"; got != want {
		t.Errorf("Client.PartnerTag() = %q, want %q", got, want)
	}
	if got, want := c.PartnerType(), defaultPartnerType; got != want {
		t.Errorf("Client.PartnerType() = %q, want %q", got, want)
	}
	cc, ok := c.(*client)
	if !ok {
		t.Fatalf("Client is not *client: %T", c)
	}
	if got, want := cc.version, CredentialVersionNA; got != want {
		t.Errorf("client.version = %q, want %q", got, want)
	}
}

func TestClientRequestSendsExpectedHeadersAndBody(t *testing.T) {
	var tokenCalls int32
	tokenHandler := func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&tokenCalls, 1)
		if got, want := r.Method, http.MethodPost; got != want {
			t.Errorf("token request method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Content-Type"), "application/x-www-form-urlencoded"; got != want {
			t.Errorf("token Content-Type = %q, want %q", got, want)
		}
		body, _ := io.ReadAll(r.Body)
		form := string(body)
		for _, want := range []string{
			"grant_type=client_credentials",
			"client_id=credID",
			"client_secret=credSecret",
			"scope=creatorsapi%2Fdefault",
		} {
			if !strings.Contains(form, want) {
				t.Errorf("token request body %q missing %q", form, want)
			}
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "tok-abc",
			"expires_in":   3600,
			"token_type":   "Bearer",
		})
	}

	var apiCalls int32
	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&apiCalls, 1)
		if got, want := r.Method, http.MethodPost; got != want {
			t.Errorf("api request method = %q, want %q", got, want)
		}
		if got, want := r.URL.Path, "/catalog/v1/getItems"; got != want {
			t.Errorf("api path = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer tok-abc, Version 2.1"; got != want {
			t.Errorf("Authorization header = %q, want %q", got, want)
		}
		if got, want := r.Header.Get(marketplaceHeader), "www.amazon.com"; got != want {
			t.Errorf("x-marketplace header = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Content-Type"), defaultContentType; got != want {
			t.Errorf("Content-Type header = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Accept"), defaultAccept; got != want {
			t.Errorf("Accept header = %q, want %q", got, want)
		}
		body, _ := io.ReadAll(r.Body)
		if got, want := string(body), `{"hello":"world"}`; got != want {
			t.Errorf("api body = %q, want %q", got, want)
		}
		_, _ = w.Write([]byte(`{"itemsResult":{}}`))
	}

	_, _, sv := newServers(t, tokenHandler, apiHandler)
	c := sv.CreateClient("mytag-20", "credID", "credSecret")

	q := stubQuery{op: GetItems, payload: []byte(`{"hello":"world"}`)}
	body, err := c.RequestContext(context.Background(), q)
	if err != nil {
		t.Fatalf("RequestContext: %v", err)
	}
	if got, want := string(body), `{"itemsResult":{}}`; got != want {
		t.Errorf("response body = %q, want %q", got, want)
	}

	if got, want := atomic.LoadInt32(&tokenCalls), int32(1); got != want {
		t.Errorf("token endpoint hit %d times, want %d", got, want)
	}
	if got, want := atomic.LoadInt32(&apiCalls), int32(1); got != want {
		t.Errorf("api endpoint hit %d times, want %d", got, want)
	}

	// Second call should reuse the cached token (no second token POST).
	if _, err := c.RequestContext(context.Background(), q); err != nil {
		t.Fatalf("second RequestContext: %v", err)
	}
	if got, want := atomic.LoadInt32(&tokenCalls), int32(1); got != want {
		t.Errorf("token cached miss: hit %d times, want %d", got, want)
	}
	if got, want := atomic.LoadInt32(&apiCalls), int32(2); got != want {
		t.Errorf("api endpoint hit %d times, want %d", got, want)
	}
}

func TestClientTokenRefreshAfterExpiry(t *testing.T) {
	var tokenCalls int32
	tokenHandler := func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&tokenCalls, 1)
		_ = json.NewEncoder(w).Encode(map[string]any{
			// expires_in <= leeway forces leeway = 0, so the cached
			// token expires effectively immediately.
			"access_token": "tok-short",
			"expires_in":   1,
		})
	}
	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("{}"))
	}
	_, _, sv := newServers(t, tokenHandler, apiHandler)
	c := sv.CreateClient("tag", "id", "secret")

	q := stubQuery{op: GetItems, payload: []byte("{}")}
	if _, err := c.RequestContext(context.Background(), q); err != nil {
		t.Fatalf("first call: %v", err)
	}
	// Sleep past the (leeway-adjusted) expiry to force a refresh.
	time.Sleep(1100 * time.Millisecond)
	if _, err := c.RequestContext(context.Background(), q); err != nil {
		t.Fatalf("second call: %v", err)
	}
	if got, want := atomic.LoadInt32(&tokenCalls), int32(2); got != want {
		t.Errorf("expected 2 token refreshes, got %d", got)
	}
}

func TestClientTokenError(t *testing.T) {
	tokenHandler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error":"invalid_client"}`, http.StatusUnauthorized)
	}
	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("api endpoint should not be called when token fails")
	}
	_, _, sv := newServers(t, tokenHandler, apiHandler)
	c := sv.CreateClient("tag", "id", "secret")

	q := stubQuery{op: GetItems, payload: []byte("{}")}
	_, err := c.RequestContext(context.Background(), q)
	if err == nil {
		t.Fatal("expected token error, got nil")
	}
}

func TestClientPayloadError(t *testing.T) {
	c := New().CreateClient("tag", "id", "secret")
	wantErr := errors.New("payload boom")
	q := stubQuery{op: GetItems, err: wantErr}
	_, err := c.RequestContext(context.Background(), q)
	if err == nil {
		t.Fatal("expected payload error, got nil")
	}
	if !errors.Is(err, wantErr) {
		t.Errorf("error chain missing payload error: %v", err)
	}
}

func TestAuthorizationHeader(t *testing.T) {
	if got, want := authorizationHeader("abc", "2.1"), "Bearer abc, Version 2.1"; got != want {
		t.Errorf("authorizationHeader = %q, want %q", got, want)
	}
	if got, want := authorizationHeader("abc", ""), "Bearer abc"; got != want {
		t.Errorf("authorizationHeader (no version) = %q, want %q", got, want)
	}
}

/* Copyright 2019,2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
