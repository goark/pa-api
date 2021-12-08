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
	}{
		{name: "www.amazon.com.au", marketplace: LocaleAustralia, str: "www.amazon.com.au", hostName: "webservices.amazon.com.au", region: "us-west-2", language: "en_AU"},
		{name: "www.amazon.com.br", marketplace: LocaleBrazil, str: "www.amazon.com.br", hostName: "webservices.amazon.com.br", region: "us-east-1", language: "pt_BR"},
		{name: "www.amazon.ca", marketplace: LocaleCanada, str: "www.amazon.ca", hostName: "webservices.amazon.ca", region: "us-east-1", language: "en_CA"},
		{name: "www.amazon.eg", marketplace: LocaleEgypt, str: "www.amazon.eg", hostName: "webservices.amazon.eg", region: "us-west-1", language: "ar_EG"},
		{name: "www.amazon.fr", marketplace: LocaleFrance, str: "www.amazon.fr", hostName: "webservices.amazon.fr", region: "eu-west-1", language: "fr_FR"},
		{name: "www.amazon.de", marketplace: LocaleGermany, str: "www.amazon.de", hostName: "webservices.amazon.de", region: "eu-west-1", language: "de_DE"},
		{name: "www.amazon.in", marketplace: LocaleIndia, str: "www.amazon.in", hostName: "webservices.amazon.in", region: "eu-west-1", language: "en_IN"},
		{name: "www.amazon.it", marketplace: LocaleItaly, str: "www.amazon.it", hostName: "webservices.amazon.it", region: "eu-west-1", language: "it_IT"},
		{name: "www.amazon.co.jp", marketplace: LocaleJapan, str: "www.amazon.co.jp", hostName: "webservices.amazon.co.jp", region: "us-west-2", language: "ja_JP"},
		{name: "www.amazon.com.mx", marketplace: LocaleMexico, str: "www.amazon.com.mx", hostName: "webservices.amazon.com.mx", region: "us-east-1", language: "es_MX"},
		{name: "www.amazon.nl", marketplace: LocaleNetherlands, str: "www.amazon.nl", hostName: "webservices.amazon.nl", region: "eu-west-1", language: "nl_NL"},
		{name: "www.amazon.pl", marketplace: LocalePoland, str: "www.amazon.pl", hostName: "webservices.amazon.pl", region: "eu-west-1", language: "pl_PL"},
		{name: "www.amazon.sg", marketplace: LocaleSingapore, str: "www.amazon.sg", hostName: "webservices.amazon.sg", region: "us-west-2", language: "en_SG"},
		{name: "www.amazon.sa", marketplace: LocaleSaudiArabia, str: "www.amazon.sa", hostName: "webservices.amazon.sa", region: "eu-west-1", language: "en_AE"},
		{name: "www.amazon.es", marketplace: LocaleSpain, str: "www.amazon.es", hostName: "webservices.amazon.es", region: "eu-west-1", language: "es_ES"},
		{name: "www.amazon.se", marketplace: LocaleSweden, str: "www.amazon.se", hostName: "webservices.amazon.se", region: "eu-west-1", language: "sv_SE"},
		{name: "www.amazon.com.tr", marketplace: LocaleTurkey, str: "www.amazon.com.tr", hostName: "webservices.amazon.com.tr", region: "eu-west-1", language: "tr_TR"},
		{name: "www.amazon.ae", marketplace: LocaleUnitedArabEmirates, str: "www.amazon.ae", hostName: "webservices.amazon.ae", region: "eu-west-1", language: "en_AE"},
		{name: "www.amazon.co.uk", marketplace: LocaleUnitedKingdom, str: "www.amazon.co.uk", hostName: "webservices.amazon.co.uk", region: "eu-west-1", language: "en_GB"},
		{name: "www.amazon.com", marketplace: LocaleUnitedStates, str: "www.amazon.com", hostName: "webservices.amazon.com", region: "us-east-1", language: "en_US"},
		{name: "foo.bar", marketplace: LocaleUnknown, str: "www.amazon.com", hostName: "webservices.amazon.com", region: "us-east-1", language: "en_US"},
	}
	for _, tc := range testCases {
		m := MarketplaceOf(tc.name)
		if m != tc.marketplace {
			t.Errorf("\"%v\" is \"%v\", want \"%v\"", tc.name, m, tc.marketplace)
		}
		if m.String() != tc.str {
			t.Errorf("Marketplace.String() is \"%v\", want \"%v\"", m.String(), tc.str)
		}
		if m.HostName() != tc.hostName {
			t.Errorf("Marketplace.HostName() is \"%v\", want \"%v\"", m.HostName(), tc.hostName)
		}
		if m.Region() != tc.region {
			t.Errorf("Marketplace.Region() is \"%v\", want \"%v\"", m.Region(), tc.region)
		}
		if m.Language() != tc.language {
			t.Errorf("Marketplace.Language() is \"%v\", want \"%v\"", m.Language(), tc.language)
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
