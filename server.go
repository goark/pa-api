package paapi5

import (
	"context"
	"net/http"
	"net/url"

	"github.com/goark/fetch"
)

const (
	defaultScheme      = "https"
	defaultAccept      = "application/json"
	defaultContentType = "application/json; charset=utf-8"
	// defaultHost is the single API host used for every marketplace under
	// the Amazon Creators API. The marketplace is selected via the
	// `x-marketplace` request header.
	defaultHost = "creatorsapi.amazon"
)

// authEndpointMap maps a Creators API credential version to the OAuth2
// (Cognito) token endpoint URL that issues tokens for that version. These
// are public Amazon endpoint URLs, not credentials — the gosec G101
// hardcoded-credentials check fires on the `oauth2/token` substring.
//
//nolint:gosec // G101: these are public endpoint URLs, not credentials.
var authEndpointMap = map[string]string{
	CredentialVersionNA: "https://creatorsapi.auth.us-east-1.amazoncognito.com/oauth2/token",
	CredentialVersionEU: "https://creatorsapi.auth.eu-south-2.amazoncognito.com/oauth2/token",
	CredentialVersionFE: "https://creatorsapi.auth.us-west-2.amazoncognito.com/oauth2/token",
}

// AuthEndpointFor returns the default OAuth2 token endpoint URL for the
// given credential version. Returns the empty string for unknown versions.
func AuthEndpointFor(version string) string {
	return authEndpointMap[version]
}

// Server type is a configuration of the Amazon Creators API service.
type Server struct {
	scheme       string
	host         string
	marketplace  Marketplace
	language     string
	authEndpoint string
}

// ServerOptFunc type is self-referential function type for New functions. (functional options pattern)
type ServerOptFunc func(*Server)

// New function returns a Server instance with options.
func New(opts ...ServerOptFunc) *Server {
	server := &Server{scheme: defaultScheme, host: defaultHost, marketplace: DefaultMarketplace, language: ""}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// WithMarketplace function returns a ServerOptFunc that sets the Marketplace.
func WithMarketplace(marketplace Marketplace) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.marketplace = marketplace
		}
	}
}

// WithLanguage function returns a ServerOptFunc that sets the desired
// response language. The Creators API does not honour an Accept-Language
// header; the value is forwarded for back-compat but callers should also
// use the LanguagesOfPreference body field for the same purpose.
func WithLanguage(language string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.language = language
		}
	}
}

// WithServerHost overrides the API service host. Useful for tests pointing
// at httptest.Server. Pass an empty string to fall back to the default.
func WithServerHost(host string) ServerOptFunc {
	return func(s *Server) {
		if s != nil && len(host) > 0 {
			s.host = host
		}
	}
}

// WithServerScheme overrides the URL scheme used by the API host. Useful
// for tests that need plain HTTP.
func WithServerScheme(scheme string) ServerOptFunc {
	return func(s *Server) {
		if s != nil && len(scheme) > 0 {
			s.scheme = scheme
		}
	}
}

// WithServerAuthEndpoint overrides the OAuth2 token endpoint resolved from
// the credential version. Primarily useful for tests.
func WithServerAuthEndpoint(endpoint string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.authEndpoint = endpoint
		}
	}
}

// URL method returns url of service server information for the Creators API.
func (s *Server) URL(path string) *url.URL {
	if s == nil {
		s = New()
	}
	return &url.URL{Scheme: s.scheme, Host: s.host, Path: path}
}

// Marketplace method returns the marketplace name.
func (s *Server) Marketplace() string {
	if s == nil {
		s = New()
	}
	return s.marketplace.String()
}

// HostName method returns the API host.
func (s *Server) HostName() string {
	if s == nil {
		s = New()
	}
	if len(s.host) > 0 {
		return s.host
	}
	return defaultHost
}

// Region method returns the historical AWS region associated with the
// configured marketplace.
//
// Deprecated: the Creators API does not use AWS SigV4 signing; the value is
// retained for back-compat callers only.
func (s *Server) Region() string {
	if s == nil {
		s = New()
	}
	return s.marketplace.Region()
}

// Accept method returns the Accept header value used for API calls.
func (s *Server) Accept() string {
	return defaultAccept
}

// AcceptLanguage method returns the Accept-Language parameter for API calls.
func (s *Server) AcceptLanguage() string {
	if s == nil {
		s = New()
	}
	if len(s.language) > 0 {
		return s.language
	}
	return s.marketplace.Language()
}

// ContentType method returns the Content-Type header value for API calls.
func (s *Server) ContentType() string {
	return defaultContentType
}

// CredentialVersion returns the Creators API credential version that
// matches the configured marketplace's region group.
func (s *Server) CredentialVersion() string {
	if s == nil {
		s = New()
	}
	return s.marketplace.CredentialVersion()
}

// AuthEndpoint returns the configured (or version-derived) OAuth2 token
// endpoint URL.
func (s *Server) AuthEndpoint() string {
	if s == nil {
		s = New()
	}
	if len(s.authEndpoint) > 0 {
		return s.authEndpoint
	}
	return AuthEndpointFor(s.CredentialVersion())
}

// ClientOptFunc type is self-referential function type for Server.CreateClient method. (functional options pattern)
type ClientOptFunc func(*client)

// CreateClient method returns a Client instance with the supplied
// associate (partner) tag and Amazon Creators API credential pair.
//
// credentialID and credentialSecret correspond to the Credential ID and
// Credential Secret issued in Associates Central > Tools > Creators API.
func (s *Server) CreateClient(associateTag, credentialID, credentialSecret string, opts ...ClientOptFunc) Client {
	if s == nil {
		s = New()
	}
	cli := &client{
		server:           s,
		httpClient:       nil,
		partnerTag:       associateTag,
		credentialID:     credentialID,
		credentialSecret: credentialSecret,
		version:          s.CredentialVersion(),
		authEndpoint:     s.AuthEndpoint(),
	}
	for _, opt := range opts {
		opt(cli)
	}
	if cli.httpClient == nil {
		cli.httpClient = fetch.New()
	}
	if len(cli.authEndpoint) == 0 {
		cli.authEndpoint = AuthEndpointFor(cli.version)
	}
	cli.auth = newTokenManager(cli.httpClient, cli.authEndpoint, cli.credentialID, cli.credentialSecret)
	return cli
}

// WithContext is retained as a no-op for backward compatibility. Pass a
// context to RequestContext instead.
//
// Deprecated: this option does nothing.
func WithContext(ctx context.Context) ClientOptFunc {
	return func(c *client) {}
}

// WithHttpClient function returns a ClientOptFunc that configures the
// underlying http.Client used for both API and OAuth2 token requests.
func WithHttpClient(hc *http.Client) ClientOptFunc {
	return func(c *client) {
		if c != nil {
			c.httpClient = fetch.New(fetch.WithHTTPClient(hc))
		}
	}
}

// WithCredentialVersion overrides the credential version derived from the
// marketplace. Use this when your Creators API credential set was issued
// for a different region than the configured marketplace would imply.
func WithCredentialVersion(version string) ClientOptFunc {
	return func(c *client) {
		if c != nil && len(version) > 0 {
			c.version = version
		}
	}
}

// WithAuthEndpoint overrides the OAuth2 token endpoint. Defaults to the
// Cognito endpoint resolved from the credential version.
func WithAuthEndpoint(endpoint string) ClientOptFunc {
	return func(c *client) {
		if c != nil && len(endpoint) > 0 {
			c.authEndpoint = endpoint
		}
	}
}

// DefaultClient function returns a default Client instance using the
// supplied associate tag and Creators API credential pair.
func DefaultClient(associateTag, credentialID, credentialSecret string) Client {
	return New().CreateClient(associateTag, credentialID, credentialSecret)
}

/* Copyright 2019-2021 Spiegel and contributors
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
