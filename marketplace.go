package paapi5

// Marketplace is interface class of locale information.
type Marketplace interface {
	String() string
	HostName() string
	Region() string
	Language() string
}

// MarketplaceEnum is enumeration of locale information.
type MarketplaceEnum int

const (
	LocaleUnknown            MarketplaceEnum = iota //Unknown local
	LocaleAustralia                                 //Australia
	LocaleBrazil                                    //Brazil
	LocaleCanada                                    //Canada
	LocaleEgypt                                     //Egypt
	LocaleFrance                                    //France
	LocaleGermany                                   //Germany
	LocaleIndia                                     //India
	LocaleItaly                                     //Italy
	LocaleJapan                                     //Japan
	LocaleMexico                                    //Mexico
	LocaleNetherlands                               //Netherlands
	LocalePoland                                    //Poland
	LocaleSingapore                                 //Singapore
	LocaleSaudiArabia                               //SaudiArabia
	LocaleSpain                                     //Spain
	LocaleSweden                                    //Sweden
	LocaleTurkey                                    //Turkey
	LocaleUnitedArabEmirates                        //United Arab Emirates
	LocaleUnitedKingdom                             //United Kingdom
	LocaleUnitedStates                              //United States
	DefaultMarketplace       = LocaleUnitedStates
)

var marketplaceMap = map[MarketplaceEnum]string{
	LocaleAustralia:          "www.amazon.com.au", //Australia
	LocaleBrazil:             "www.amazon.com.br", //Brazil
	LocaleCanada:             "www.amazon.ca",     //Canada
	LocaleEgypt:              "www.amazon.eg",     //Egypt
	LocaleFrance:             "www.amazon.fr",     //France
	LocaleGermany:            "www.amazon.de",     //Germany
	LocaleIndia:              "www.amazon.in",     //India
	LocaleItaly:              "www.amazon.it",     //Italy
	LocaleJapan:              "www.amazon.co.jp",  //Japan
	LocaleMexico:             "www.amazon.com.mx", //Mexico
	LocaleNetherlands:        "www.amazon.nl",     //Netherlands
	LocalePoland:             "www.amazon.pl",     //Poland
	LocaleSingapore:          "www.amazon.sg",     //Singapore
	LocaleSaudiArabia:        "www.amazon.sa",     //SaudiArabia
	LocaleSpain:              "www.amazon.es",     //Spain
	LocaleSweden:             "www.amazon.se",     //Sweden
	LocaleTurkey:             "www.amazon.com.tr", //Turkey
	LocaleUnitedArabEmirates: "www.amazon.ae",     //United Arab Emirates
	LocaleUnitedKingdom:      "www.amazon.co.uk",  //United Kingdom
	LocaleUnitedStates:       "www.amazon.com",    //United States
}

var hostMap = map[MarketplaceEnum]string{
	LocaleAustralia:          "webservices.amazon.com.au", //Australia
	LocaleBrazil:             "webservices.amazon.com.br", //Brazil
	LocaleCanada:             "webservices.amazon.ca",     //Canada
	LocaleEgypt:              "webservices.amazon.eg",     //Egypt
	LocaleFrance:             "webservices.amazon.fr",     //France
	LocaleGermany:            "webservices.amazon.de",     //Germany
	LocaleIndia:              "webservices.amazon.in",     //India
	LocaleItaly:              "webservices.amazon.it",     //Italy
	LocaleJapan:              "webservices.amazon.co.jp",  //Japan
	LocaleMexico:             "webservices.amazon.com.mx", //Mexico
	LocaleNetherlands:        "webservices.amazon.nl",     //Netherlands
	LocalePoland:             "webservices.amazon.pl",     //Poland
	LocaleSingapore:          "webservices.amazon.sg",     //Singapore
	LocaleSaudiArabia:        "webservices.amazon.sa",     //SaudiArabia
	LocaleSpain:              "webservices.amazon.es",     //Spain
	LocaleSweden:             "webservices.amazon.se",     //Sweden
	LocaleTurkey:             "webservices.amazon.com.tr", //Turkey
	LocaleUnitedArabEmirates: "webservices.amazon.ae",     //United Arab Emirates
	LocaleUnitedKingdom:      "webservices.amazon.co.uk",  //United Kingdom
	LocaleUnitedStates:       "webservices.amazon.com",    //United States
}

var regionMap = map[MarketplaceEnum]string{
	LocaleAustralia:          "us-west-2", //Australia
	LocaleBrazil:             "us-east-1", //Brazil
	LocaleCanada:             "us-east-1", //Canada
	LocaleEgypt:              "us-west-1", //Egypt
	LocaleFrance:             "eu-west-1", //France
	LocaleGermany:            "eu-west-1", //Germany
	LocaleIndia:              "eu-west-1", //India
	LocaleItaly:              "eu-west-1", //Italy
	LocaleJapan:              "us-west-2", //Japan
	LocaleMexico:             "us-east-1", //Mexico
	LocaleNetherlands:        "eu-west-1", //Netherlands
	LocalePoland:             "eu-west-1", //Poland
	LocaleSingapore:          "us-west-2", //Singapore
	LocaleSaudiArabia:        "eu-west-1", //SaudiArabia
	LocaleSpain:              "eu-west-1", //Spain
	LocaleSweden:             "eu-west-1", //Sweden
	LocaleTurkey:             "eu-west-1", //Turkey
	LocaleUnitedArabEmirates: "eu-west-1", //United Arab Emirates
	LocaleUnitedKingdom:      "eu-west-1", //United Kingdom
	LocaleUnitedStates:       "us-east-1", //United States
}

var languageMap = map[MarketplaceEnum]string{
	LocaleAustralia:          "en_AU", //Australia
	LocaleBrazil:             "pt_BR", //Brazil
	LocaleCanada:             "en_CA", //Canada
	LocaleEgypt:              "ar_EG", //Egypt
	LocaleFrance:             "fr_FR", //France
	LocaleGermany:            "de_DE", //Germany
	LocaleIndia:              "en_IN", //India
	LocaleItaly:              "it_IT", //Italy
	LocaleJapan:              "ja_JP", //Japan
	LocaleMexico:             "es_MX", //Mexico
	LocaleNetherlands:        "nl_NL", //Netherlands
	LocalePoland:             "pl_PL", //Poland
	LocaleSingapore:          "en_SG", //Singapore
	LocaleSaudiArabia:        "en_AE", //SaudiArabia
	LocaleSpain:              "es_ES", //Spain
	LocaleSweden:             "sv_SE", //Sweden
	LocaleTurkey:             "tr_TR", //Turkey
	LocaleUnitedArabEmirates: "en_AE", //United Arab Emirates
	LocaleUnitedKingdom:      "en_GB", //United Kingdom
	LocaleUnitedStates:       "en_US", //United States
}

// MarketplaceOf function returns Marketplace instance from service domain.
func MarketplaceOf(s string) Marketplace {
	for k, v := range marketplaceMap {
		if s == v {
			return k
		}
	}
	return LocaleUnknown
}

// String returns marketplace name of Marketplace.
func (m MarketplaceEnum) String() string {
	if s, ok := marketplaceMap[m]; ok {
		return s
	}
	return marketplaceMap[DefaultMarketplace]
}

// HostName returns hostname of Marketplace.
func (m MarketplaceEnum) HostName() string {
	if s, ok := hostMap[m]; ok {
		return s
	}
	return hostMap[LocaleUnitedStates]
}

// Region returns region name of Marketplace.
func (m MarketplaceEnum) Region() string {
	if s, ok := regionMap[m]; ok {
		return s
	}
	return regionMap[DefaultMarketplace]
}

// Language returns language name of Marketplace.
func (m MarketplaceEnum) Language() string {
	if s, ok := languageMap[m]; ok {
		return s
	}
	return languageMap[DefaultMarketplace]
}

/* Copyright 2019-2021 Spiegel
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
