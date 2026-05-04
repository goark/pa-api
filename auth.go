package paapi5

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/goark/errs"
	"github.com/goark/fetch"
)

const (
	// oauthScope is the scope requested when fetching an access token.
	oauthScope = "creatorsapi/default"
	// oauthGrantType is the OAuth2 grant type used for the client_credentials flow.
	oauthGrantType = "client_credentials"
	// oauthTokenLeewaySeconds keeps a small buffer before the announced expiration
	// to avoid handing out a token that expires mid-request.
	oauthTokenLeewaySeconds = 30
)

// tokenManager handles the OAuth2 client_credentials flow against a Cognito
// token endpoint. Tokens are cached in memory and reused across requests until
// they are within oauthTokenLeewaySeconds of expiring.
type tokenManager struct {
	httpClient   fetch.Client
	endpoint     string
	clientID     string
	clientSecret string

	mu          sync.Mutex
	accessToken string
	expiresAt   time.Time
}

// newTokenManager constructs a tokenManager. httpClient may be nil, in which
// case fetch.New() is used.
func newTokenManager(httpClient fetch.Client, endpoint, clientID, clientSecret string) *tokenManager {
	if httpClient == nil {
		httpClient = fetch.New()
	}
	return &tokenManager{
		httpClient:   httpClient,
		endpoint:     endpoint,
		clientID:     clientID,
		clientSecret: clientSecret,
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
func (t *tokenManager) refreshLocked(ctx context.Context) error {
	if len(strings.TrimSpace(t.endpoint)) == 0 {
		return errs.Wrap(ErrNullPointer, errs.WithContext("reason", "empty OAuth2 token endpoint"))
	}
	u, err := url.Parse(t.endpoint)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("endpoint", t.endpoint))
	}
	form := url.Values{}
	form.Set("grant_type", oauthGrantType)
	form.Set("client_id", t.clientID)
	form.Set("client_secret", t.clientSecret)
	form.Set("scope", oauthScope)
	resp, err := t.httpClient.PostWithContext(
		ctx,
		u,
		strings.NewReader(form.Encode()),
		fetch.WithRequestHeaderSet("Content-Type", "application/x-www-form-urlencoded"),
		fetch.WithRequestHeaderSet("Accept", "application/json"),
	)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("endpoint", t.endpoint))
	}
	body, err := resp.DumpBodyAndClose()
	if err != nil {
		return errs.Wrap(err, errs.WithContext("endpoint", t.endpoint))
	}
	tr := tokenResponse{}
	if err := json.Unmarshal(body, &tr); err != nil {
		return errs.Wrap(err, errs.WithContext("endpoint", t.endpoint), errs.WithContext("body", string(body)))
	}
	if tr.AccessToken == "" {
		return errs.Wrap(ErrNoData, errs.WithContext("endpoint", t.endpoint), errs.WithContext("body", string(body)))
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

// authorizationHeader returns the value of the Authorization header expected
// by the Creators API: `Bearer <token>, Version <version>`.
func authorizationHeader(token, version string) string {
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
