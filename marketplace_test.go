package paapi5

import "testing"

func TestMarketplace(t *testing.T) {
	testCases := []struct {
		name        string
		marketplace Marketplace
		str         string
		hostName    string
		region      string
		language    string
		version     string
	}{
		{name: "www.amazon.com.au", marketplace: LocaleAustralia, str: "www.amazon.com.au", hostName: defaultHost, region: "us-west-2", language: "en_AU", version: CredentialVersionFE},
		{name: "www.amazon.com.br", marketplace: LocaleBrazil, str: "www.amazon.com.br", hostName: defaultHost, region: "us-east-1", language: "pt_BR", version: CredentialVersionNA},
		{name: "www.amazon.ca", marketplace: LocaleCanada, str: "www.amazon.ca", hostName: defaultHost, region: "us-east-1", language: "en_CA", version: CredentialVersionNA},
		{name: "www.amazon.eg", marketplace: LocaleEgypt, str: "www.amazon.eg", hostName: defaultHost, region: "us-west-1", language: "ar_EG", version: CredentialVersionEU},
		{name: "www.amazon.fr", marketplace: LocaleFrance, str: "www.amazon.fr", hostName: defaultHost, region: "eu-west-1", language: "fr_FR", version: CredentialVersionEU},
		{name: "www.amazon.de", marketplace: LocaleGermany, str: "www.amazon.de", hostName: defaultHost, region: "eu-west-1", language: "de_DE", version: CredentialVersionEU},
		{name: "www.amazon.in", marketplace: LocaleIndia, str: "www.amazon.in", hostName: defaultHost, region: "eu-west-1", language: "en_IN", version: CredentialVersionEU},
		{name: "www.amazon.ie", marketplace: LocaleIreland, str: "www.amazon.ie", hostName: defaultHost, region: "eu-west-1", language: "en_IE", version: CredentialVersionEU},
		{name: "www.amazon.it", marketplace: LocaleItaly, str: "www.amazon.it", hostName: defaultHost, region: "eu-west-1", language: "it_IT", version: CredentialVersionEU},
		{name: "www.amazon.co.jp", marketplace: LocaleJapan, str: "www.amazon.co.jp", hostName: defaultHost, region: "us-west-2", language: "ja_JP", version: CredentialVersionFE},
		{name: "www.amazon.com.mx", marketplace: LocaleMexico, str: "www.amazon.com.mx", hostName: defaultHost, region: "us-east-1", language: "es_MX", version: CredentialVersionNA},
		{name: "www.amazon.nl", marketplace: LocaleNetherlands, str: "www.amazon.nl", hostName: defaultHost, region: "eu-west-1", language: "nl_NL", version: CredentialVersionEU},
		{name: "www.amazon.pl", marketplace: LocalePoland, str: "www.amazon.pl", hostName: defaultHost, region: "eu-west-1", language: "pl_PL", version: CredentialVersionEU},
		{name: "www.amazon.sg", marketplace: LocaleSingapore, str: "www.amazon.sg", hostName: defaultHost, region: "us-west-2", language: "en_SG", version: CredentialVersionFE},
		{name: "www.amazon.sa", marketplace: LocaleSaudiArabia, str: "www.amazon.sa", hostName: defaultHost, region: "eu-west-1", language: "en_AE", version: CredentialVersionEU},
		{name: "www.amazon.es", marketplace: LocaleSpain, str: "www.amazon.es", hostName: defaultHost, region: "eu-west-1", language: "es_ES", version: CredentialVersionEU},
		{name: "www.amazon.se", marketplace: LocaleSweden, str: "www.amazon.se", hostName: defaultHost, region: "eu-west-1", language: "sv_SE", version: CredentialVersionEU},
		{name: "www.amazon.com.tr", marketplace: LocaleTurkey, str: "www.amazon.com.tr", hostName: defaultHost, region: "eu-west-1", language: "tr_TR", version: CredentialVersionEU},
		{name: "www.amazon.ae", marketplace: LocaleUnitedArabEmirates, str: "www.amazon.ae", hostName: defaultHost, region: "eu-west-1", language: "en_AE", version: CredentialVersionEU},
		{name: "www.amazon.co.uk", marketplace: LocaleUnitedKingdom, str: "www.amazon.co.uk", hostName: defaultHost, region: "eu-west-1", language: "en_GB", version: CredentialVersionEU},
		{name: "www.amazon.com", marketplace: LocaleUnitedStates, str: "www.amazon.com", hostName: defaultHost, region: "us-east-1", language: "en_US", version: CredentialVersionNA},
		{name: "foo.bar", marketplace: LocaleUnknown, str: "www.amazon.com", hostName: defaultHost, region: "us-east-1", language: "en_US", version: CredentialVersionNA},
	}
	for _, tc := range testCases {
		m := MarketplaceOf(tc.name)
		if m != tc.marketplace {
			t.Errorf("%q is %v, want %v", tc.name, m, tc.marketplace)
		}
		if m.String() != tc.str {
			t.Errorf("Marketplace.String() is %q, want %q", m.String(), tc.str)
		}
		if m.HostName() != tc.hostName {
			t.Errorf("Marketplace.HostName() is %q, want %q", m.HostName(), tc.hostName)
		}
		if m.Region() != tc.region {
			t.Errorf("Marketplace.Region() is %q, want %q", m.Region(), tc.region)
		}
		if m.Language() != tc.language {
			t.Errorf("Marketplace.Language() is %q, want %q", m.Language(), tc.language)
		}
		// CredentialVersion is reported via the optional credentialVersioner
		// interface (satisfied by MarketplaceEnum), not the Marketplace
		// interface itself, so it must be type-asserted or fetched via the
		// package-level credentialVersionOf helper.
		if got := credentialVersionOf(m); got != tc.version {
			t.Errorf("credentialVersionOf(%q) is %q, want %q", tc.name, got, tc.version)
		}
	}
}

// fakeMarketplace omits CredentialVersion entirely; credentialVersionOf must
// fall back to the default NA version for these implementations.
type fakeMarketplace struct{}

func (fakeMarketplace) String() string   { return "fake.amazon" }
func (fakeMarketplace) HostName() string { return "fake.amazon" }
func (fakeMarketplace) Region() string   { return "us-east-1" }
func (fakeMarketplace) Language() string { return "en_US" }

func TestCredentialVersionOfFallsBackForExternalMarketplaces(t *testing.T) {
	if got, want := credentialVersionOf(fakeMarketplace{}), CredentialVersionNA; got != want {
		t.Errorf("credentialVersionOf(fakeMarketplace) = %q, want %q (default NA fallback)", got, want)
	}
}

func TestAuthEndpointFor(t *testing.T) {
	testCases := []struct {
		version  string
		endpoint string
	}{
		{version: CredentialVersionNA, endpoint: "https://creatorsapi.auth.us-east-1.amazoncognito.com/oauth2/token"},
		{version: CredentialVersionEU, endpoint: "https://creatorsapi.auth.eu-south-2.amazoncognito.com/oauth2/token"},
		{version: CredentialVersionFE, endpoint: "https://creatorsapi.auth.us-west-2.amazoncognito.com/oauth2/token"},
		{version: "9.9", endpoint: ""},
	}
	for _, tc := range testCases {
		if got := AuthEndpointFor(tc.version); got != tc.endpoint {
			t.Errorf("AuthEndpointFor(%q) is %q, want %q", tc.version, got, tc.endpoint)
		}
	}
}

/* Copyright 2019 Spiegel
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
