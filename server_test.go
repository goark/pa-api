package paapi5

import "testing"

func TestServer(t *testing.T) {
	testCases := []struct {
		sv             *Server
		marketplace    string
		hostName       string
		region         string
		accept         string
		acceptLanguage string
		contentType    string
		version        string
		authEndpoint   string
		url            string
	}{
		{
			sv:             (*Server)(nil),
			marketplace:    "www.amazon.com",
			hostName:       defaultHost,
			region:         "us-east-1",
			accept:         defaultAccept,
			acceptLanguage: "en_US",
			contentType:    defaultContentType,
			version:        CredentialVersionNAv3,
			authEndpoint:   "https://api.amazon.com/auth/o2/token",
			url:            "https://creatorsapi.amazon/catalog/v1/getItems",
		},
		{
			sv:             New(WithMarketplace(LocaleJapan)),
			marketplace:    "www.amazon.co.jp",
			hostName:       defaultHost,
			region:         "us-west-2",
			accept:         defaultAccept,
			acceptLanguage: "ja_JP",
			contentType:    defaultContentType,
			version:        CredentialVersionFEv3,
			authEndpoint:   "https://api.amazon.co.jp/auth/o2/token",
			url:            "https://creatorsapi.amazon/catalog/v1/getItems",
		},
		{
			sv:             New(WithMarketplace(LocaleGermany)),
			marketplace:    "www.amazon.de",
			hostName:       defaultHost,
			region:         "eu-west-1",
			accept:         defaultAccept,
			acceptLanguage: "de_DE",
			contentType:    defaultContentType,
			version:        CredentialVersionEUv3,
			authEndpoint:   "https://api.amazon.co.uk/auth/o2/token",
			url:            "https://creatorsapi.amazon/catalog/v1/getItems",
		},
	}
	for _, tc := range testCases {
		if tc.sv.Marketplace() != tc.marketplace {
			t.Errorf("Server.Marketplace() is %q, want %q", tc.sv.Marketplace(), tc.marketplace)
		}
		if tc.sv.HostName() != tc.hostName {
			t.Errorf("Server.HostName() is %q, want %q", tc.sv.HostName(), tc.hostName)
		}
		if tc.sv.Region() != tc.region {
			t.Errorf("Server.Region() is %q, want %q", tc.sv.Region(), tc.region)
		}
		if tc.sv.Accept() != tc.accept {
			t.Errorf("Server.Accept() is %q, want %q", tc.sv.Accept(), tc.accept)
		}
		if tc.sv.AcceptLanguage() != tc.acceptLanguage {
			t.Errorf("Server.AcceptLanguage() is %q, want %q", tc.sv.AcceptLanguage(), tc.acceptLanguage)
		}
		if tc.sv.ContentType() != tc.contentType {
			t.Errorf("Server.ContentType() is %q, want %q", tc.sv.ContentType(), tc.contentType)
		}
		if tc.sv.CredentialVersion() != tc.version {
			t.Errorf("Server.CredentialVersion() is %q, want %q", tc.sv.CredentialVersion(), tc.version)
		}
		if tc.sv.AuthEndpoint() != tc.authEndpoint {
			t.Errorf("Server.AuthEndpoint() is %q, want %q", tc.sv.AuthEndpoint(), tc.authEndpoint)
		}
		url := tc.sv.URL(GetItems.Path()).String()
		if url != tc.url {
			t.Errorf("Server.URL() is %q, want %q", url, tc.url)
		}
	}
}

func TestServerLanguageOverride(t *testing.T) {
	sv := New(WithMarketplace(LocaleUnitedStates), WithLanguage("fr_CA"))
	if got, want := sv.AcceptLanguage(), "fr_CA"; got != want {
		t.Errorf("Server.AcceptLanguage() with override is %q, want %q", got, want)
	}
}

func TestServerHostOverride(t *testing.T) {
	sv := New(
		WithServerHost("api.example.test"),
		WithServerScheme("http"),
		WithServerAuthEndpoint("http://auth.example.test/oauth2/token"),
	)
	if got, want := sv.URL(GetItems.Path()).String(), "http://api.example.test/catalog/v1/getItems"; got != want {
		t.Errorf("Server.URL() with override is %q, want %q", got, want)
	}
	if got, want := sv.AuthEndpoint(), "http://auth.example.test/oauth2/token"; got != want {
		t.Errorf("Server.AuthEndpoint() with override is %q, want %q", got, want)
	}
}

func TestWithCredentialVersionIgnoresUnsupportedValue(t *testing.T) {
	c := New().CreateClient("tag", "id", "secret", WithCredentialVersion("9.9"))
	cc, ok := c.(*client)
	if !ok {
		t.Fatalf("Client is not *client: %T", c)
	}
	if got, want := cc.version, CredentialVersionNAv3; got != want {
		t.Errorf("client.version = %q, want %q", got, want)
	}
	if got, want := cc.authEndpoint, AuthEndpointFor(CredentialVersionNAv3); got != want {
		t.Errorf("client.authEndpoint = %q, want %q", got, want)
	}
}

/* Copyright 2019 Spiegel and contributors
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
