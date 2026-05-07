package paapi5

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/goark/errs"
)

const (
	// oauthScopeCognito is the scope for v2.x (Cognito /oauth2/token).
	oauthScopeCognito = "creatorsapi/default"
	// oauthScopeLWA is the scope for v3.x (Login with Amazon /auth/o2/token).
	oauthScopeLWA = "creatorsapi::default"
	// oauthGrantType is the OAuth2 grant type used for the client_credentials flow.
	oauthGrantType = "client_credentials"
	// oauthTokenLeewaySeconds keeps a small buffer before the announced expiration
	// to avoid handing out a token that expires mid-request.
	oauthTokenLeewaySeconds = 30
	// maxTokenBodyReadBytes caps how much of the token endpoint response body
	// we read into memory, regardless of Content-Length.
	maxTokenBodyReadBytes = 8 * 1024
	// maxTokenBodyContextBytes caps how much of the body we attach to error
	// contexts so we don't propagate arbitrary endpoint output into logs.
	maxTokenBodyContextBytes = 256
)

// tokenManager handles the OAuth2 client_credentials flow against a Cognito
// token endpoint. Tokens are cached in memory and reused across requests until
// they are within oauthTokenLeewaySeconds of expiring.
type tokenManager struct {
	httpClient   *http.Client
	endpoint     string
	clientID     string
	clientSecret string
	lwa          bool

	mu          sync.Mutex
	accessToken string
	expiresAt   time.Time
}

// newTokenManager constructs a tokenManager. httpClient may be nil, in which
// case http.DefaultClient is used.
func newTokenManager(httpClient *http.Client, endpoint, clientID, clientSecret string, lwa bool) *tokenManager {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &tokenManager{
		httpClient:   httpClient,
		endpoint:     endpoint,
		clientID:     clientID,
		clientSecret: clientSecret,
		lwa:          lwa,
	}
}

// Token returns a valid OAuth2 access token, refreshing it if the cached one
// is missing or near expiration.
func (t *tokenManager) Token(ctx context.Context) (string, error) {
	if t == nil {
		return "", errs.Wrap(ErrNullPointer)
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.accessToken != "" && time.Now().Before(t.expiresAt) {
		return t.accessToken, nil
	}
	if err := t.refreshLocked(ctx); err != nil {
		t.accessToken = ""
		t.expiresAt = time.Time{}
		return "", err
	}
	return t.accessToken, nil
}

// tokenResponse mirrors the relevant subset of a Cognito token endpoint reply.
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// refreshLocked POSTs the client_credentials grant to the configured token
// endpoint and stores the resulting access token under the existing lock.
//
// The token POST is issued via *http.Client directly (not the fetch wrapper)
// so that we can inspect the HTTP status code and capture a bounded slice of
// the response body for error diagnostics on non-2xx responses.
func (t *tokenManager) refreshLocked(ctx context.Context) error {
	if len(strings.TrimSpace(t.endpoint)) == 0 {
		return errs.Wrap(ErrNullPointer, errs.WithContext("reason", "empty OAuth2 token endpoint"))
	}
	form := url.Values{}
	form.Set("grant_type", oauthGrantType)
	if t.lwa {
		form.Set("scope", oauthScopeLWA)
	} else {
		form.Set("client_id", t.clientID)
		form.Set("client_secret", t.clientSecret)
		form.Set("scope", oauthScopeCognito)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("endpoint", t.endpoint))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	if t.lwa {
		basic := base64.StdEncoding.EncodeToString([]byte(t.clientID + ":" + t.clientSecret))
		req.Header.Set("Authorization", "Basic "+basic)
	}
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("endpoint", t.endpoint))
	}
	defer func() { _ = resp.Body.Close() }()
	body, readErr := io.ReadAll(io.LimitReader(resp.Body, maxTokenBodyReadBytes))
	if readErr != nil {
		return errs.Wrap(readErr,
			errs.WithContext("endpoint", t.endpoint),
			errs.WithContext("status", resp.StatusCode),
		)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errs.Wrap(
			fmt.Errorf("%w: HTTP %d", ErrHTTPStatus, resp.StatusCode),
			errs.WithContext("endpoint", t.endpoint),
			errs.WithContext("status", resp.StatusCode),
			errs.WithContext("body", truncateForLog(body, maxTokenBodyContextBytes)),
		)
	}
	tr := tokenResponse{}
	if err := json.Unmarshal(body, &tr); err != nil {
		return errs.Wrap(err,
			errs.WithContext("endpoint", t.endpoint),
			errs.WithContext("status", resp.StatusCode),
			errs.WithContext("body", truncateForLog(body, maxTokenBodyContextBytes)),
		)
	}
	if tr.AccessToken == "" {
		return errs.Wrap(ErrNoData,
			errs.WithContext("endpoint", t.endpoint),
			errs.WithContext("status", resp.StatusCode),
			errs.WithContext("body", truncateForLog(body, maxTokenBodyContextBytes)),
		)
	}
	expiresIn := tr.ExpiresIn
	if expiresIn <= 0 {
		expiresIn = 3600
	}
	leeway := oauthTokenLeewaySeconds
	if leeway >= expiresIn {
		leeway = expiresIn / 2
	}
	t.accessToken = tr.AccessToken
	t.expiresAt = time.Now().Add(time.Duration(expiresIn-leeway) * time.Second)
	return nil
}

// truncateForLog returns a string copy of b clipped to max bytes, with a
// trailing marker if truncation occurred. Used to keep arbitrary token
// endpoint output out of logs at full length.
func truncateForLog(b []byte, max int) string {
	if len(b) <= max {
		return string(b)
	}
	return string(b[:max]) + "...(truncated)"
}

// authorizationHeader returns the Creators API catalog Authorization header.
// v3.x (Login with Amazon) tokens use `Bearer <token>` only; v2.x Cognito
// tokens append `, Version <version>`.
func authorizationHeader(token, version string, lwa bool) string {
	if lwa {
		return "Bearer " + token
	}
	b := strings.Builder{}
	b.WriteString("Bearer ")
	b.WriteString(token)
	if len(version) > 0 {
		b.WriteString(", Version ")
		b.WriteString(version)
	}
	return b.String()
}

/* Copyright 2026 Spiegel and contributors
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
