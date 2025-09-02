package query

import (
	"errors"
	"testing"

	paapi5 "github.com/goark/pa-api"
)

func TestNilQuery(t *testing.T) {
	testCases := []struct {
		q   *Query
		op  paapi5.Operation
		err error
		jsn string
	}{
		{q: nil, op: paapi5.NullOperation, err: paapi5.ErrNullPointer, jsn: ""},
		{q: (*Query)(nil), op: paapi5.NullOperation, err: paapi5.ErrNullPointer, jsn: ""},
		{q: (*Query)(nil).With(), op: paapi5.NullOperation, err: nil, jsn: `{"Operation":""}`},
		{q: New(paapi5.GetItems), op: paapi5.GetItems, err: nil, jsn: `{"Operation":"GetItems"}`},
	}

	for _, tc := range testCases {
		if op := tc.q.Operation(); op != tc.op {
			t.Errorf("Query.Operation() is \"%v\", want \"%v\"", op, tc.op)

		}
		if b, err := tc.q.Payload(); !errors.Is(err, tc.err) {
			t.Errorf("Query.Payload() is \"%v\", want \"%v\"", err, tc.err)
		} else if string(b) != tc.jsn {
			t.Errorf("Query.Payload() is \"%v\", want \"%v\"", string(b), tc.jsn)
		}
	}
}

func TestRequestFilters(t *testing.T) {
	empty := (*Query)(nil)
	testCases := []struct {
		q   *Query
		str string
	}{
		{q: empty, str: ""},
		{q: empty.With(), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{RequestFilter(0): "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{Actor: "foo"}), str: `{"Operation":"","Actor":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{Artist: "foo"}), str: `{"Operation":"","Artist":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{ASIN: "4900900028"}), str: `{"Operation":"","ASIN":"4900900028"}`},
		{q: empty.With().RequestFilters(RequestMap{Availability: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{Availability: "Available"}), str: `{"Operation":"","Availability":"Available"}`},
		{q: empty.With().RequestFilters(RequestMap{Availability: "IncludeOutOfStock"}), str: `{"Operation":"","Availability":"IncludeOutOfStock"}`},
		{q: empty.With().RequestFilters(RequestMap{Author: "foo"}), str: `{"Operation":"","Author":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{Brand: "foo"}), str: `{"Operation":"","Brand":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeID: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeID: "123"}), str: `{"Operation":"","BrowseNodeId":"123"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Any"}), str: `{"Operation":"","Condition":"Any"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "New"}), str: `{"Operation":"","Condition":"New"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Used"}), str: `{"Operation":"","Condition":"Used"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Collectible"}), str: `{"Operation":"","Condition":"Collectible"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Refurbished"}), str: `{"Operation":"","Condition":"Refurbished"}`},
		{q: empty.With().RequestFilters(RequestMap{CurrencyOfPreference: "foo"}), str: `{"Operation":"","CurrencyOfPreference":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{DeliveryFlags: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{DeliveryFlags: "AmazonGlobal"}), str: `{"Operation":"","DeliveryFlags":["AmazonGlobal"]}`},
		{q: empty.With().RequestFilters(RequestMap{DeliveryFlags: []string{"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"}}), str: `{"Operation":"","DeliveryFlags":["AmazonGlobal","FreeShipping","FulfilledByAmazon","Prime"]}`},
		{q: empty.With().RequestFilters(RequestMap{ItemIds: "4900900028", ItemIdType: "ASIN"}), str: `{"Operation":"","ItemIds":["4900900028"],"ItemIdType":"ASIN"}`},
		{q: empty.With().RequestFilters(RequestMap{ItemIds: "4900900028"}, RequestMap{ItemIdType: "ASIN"}), str: `{"Operation":"","ItemIds":["4900900028"],"ItemIdType":"ASIN"}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 1}), str: `{"Operation":"","ItemCount":1}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 10}), str: `{"Operation":"","ItemCount":10}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 11}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 1}), str: `{"Operation":"","ItemPage":1}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 10}), str: `{"Operation":"","ItemPage":10}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 11}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{Keywords: "foo"}), str: `{"Operation":"","Keywords":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeIds: "123"}), str: `{"Operation":"","BrowseNodeIds":["123"]}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeIds: []string{"123", "456"}}), str: `{"Operation":"","BrowseNodeIds":["123","456"]}`},
		{q: empty.With().RequestFilters(RequestMap{LanguagesOfPreference: "foo"}), str: `{"Operation":"","LanguagesOfPreference":["foo"]}`},
		{q: empty.With().RequestFilters(RequestMap{LanguagesOfPreference: []string{"foo", "bar"}}), str: `{"Operation":"","LanguagesOfPreference":["foo","bar"]}`},
		{q: empty.With().RequestFilters(RequestMap{Marketplace: "foo.bar"}), str: `{"Operation":"","Marketplace":"foo.bar"}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: 1}), str: `{"Operation":"","MaxPrice":1}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: 123}), str: `{"Operation":"","MaxPrice":123}`},
		{q: empty.With().RequestFilters(RequestMap{Merchant: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{Merchant: "All"}), str: `{"Operation":"","Merchant":"All"}`},
		{q: empty.With().RequestFilters(RequestMap{Merchant: "Amazon"}), str: `{"Operation":"","Merchant":"Amazon"}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: 1}), str: `{"Operation":"","MinPrice":1}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: 123}), str: `{"Operation":"","MinPrice":123}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 1}), str: `{"Operation":"","MinReviewsRating":1}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 4}), str: `{"Operation":"","MinReviewsRating":4}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 5}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 1}), str: `{"Operation":"","MinSavingPercent":1}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 99}), str: `{"Operation":"","MinSavingPercent":99}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 100}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{OfferCount: -1}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{OfferCount: 0}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{OfferCount: 1}), str: `{"Operation":"","OfferCount":1}`},
		{q: empty.With().RequestFilters(RequestMap{OfferCount: 123}), str: `{"Operation":"","OfferCount":123}`},
		{q: empty.With().RequestFilters(RequestMap{PartnerTag: "foo"}), str: `{"Operation":"","PartnerTag":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{PartnerType: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{PartnerType: "Associates"}), str: `{"Operation":"","PartnerType":"Associates"}`},
		{q: empty.With().RequestFilters(RequestMap{Properties: map[string]string{"foo": "bar"}}), str: `{"Operation":"","Properties":{"foo":"bar"}}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "All"}), str: `{"Operation":"","SearchIndex":"All"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "AmazonVideo"}), str: `{"Operation":"","SearchIndex":"AmazonVideo"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Apparel"}), str: `{"Operation":"","SearchIndex":"Apparel"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Appliances"}), str: `{"Operation":"","SearchIndex":"Appliances"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "ArtsAndCrafts"}), str: `{"Operation":"","SearchIndex":"ArtsAndCrafts"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Automotive"}), str: `{"Operation":"","SearchIndex":"Automotive"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Baby"}), str: `{"Operation":"","SearchIndex":"Baby"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Beauty"}), str: `{"Operation":"","SearchIndex":"Beauty"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Books"}), str: `{"Operation":"","SearchIndex":"Books"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Classical"}), str: `{"Operation":"","SearchIndex":"Classical"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Collectibles"}), str: `{"Operation":"","SearchIndex":"Collectibles"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Computers"}), str: `{"Operation":"","SearchIndex":"Computers"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "DigitalMusic"}), str: `{"Operation":"","SearchIndex":"DigitalMusic"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Electronics"}), str: `{"Operation":"","SearchIndex":"Electronics"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "EverythingElse"}), str: `{"Operation":"","SearchIndex":"EverythingElse"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Fashion"}), str: `{"Operation":"","SearchIndex":"Fashion"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "FashionBaby"}), str: `{"Operation":"","SearchIndex":"FashionBaby"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "FashionBoys"}), str: `{"Operation":"","SearchIndex":"FashionBoys"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "FashionGirls"}), str: `{"Operation":"","SearchIndex":"FashionGirls"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "FashionMen"}), str: `{"Operation":"","SearchIndex":"FashionMen"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "FashionWomen"}), str: `{"Operation":"","SearchIndex":"FashionWomen"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "GardenAndOutdoor"}), str: `{"Operation":"","SearchIndex":"GardenAndOutdoor"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "GiftCards"}), str: `{"Operation":"","SearchIndex":"GiftCards"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "GroceryAndGourmetFood"}), str: `{"Operation":"","SearchIndex":"GroceryAndGourmetFood"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Handmade"}), str: `{"Operation":"","SearchIndex":"Handmade"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "HealthPersonalCare"}), str: `{"Operation":"","SearchIndex":"HealthPersonalCare"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "HomeAndKitchen"}), str: `{"Operation":"","SearchIndex":"HomeAndKitchen"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Industrial"}), str: `{"Operation":"","SearchIndex":"Industrial"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Jewelry"}), str: `{"Operation":"","SearchIndex":"Jewelry"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "KindleStore"}), str: `{"Operation":"","SearchIndex":"KindleStore"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "LocalServices"}), str: `{"Operation":"","SearchIndex":"LocalServices"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Luggage"}), str: `{"Operation":"","SearchIndex":"Luggage"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "LuxuryBeauty"}), str: `{"Operation":"","SearchIndex":"LuxuryBeauty"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Magazines"}), str: `{"Operation":"","SearchIndex":"Magazines"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "MobileAndAccessories"}), str: `{"Operation":"","SearchIndex":"MobileAndAccessories"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "MobileApps"}), str: `{"Operation":"","SearchIndex":"MobileApps"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "MoviesAndTV"}), str: `{"Operation":"","SearchIndex":"MoviesAndTV"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Music"}), str: `{"Operation":"","SearchIndex":"Music"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "MusicalInstruments"}), str: `{"Operation":"","SearchIndex":"MusicalInstruments"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "OfficeProducts"}), str: `{"Operation":"","SearchIndex":"OfficeProducts"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "PetSupplies"}), str: `{"Operation":"","SearchIndex":"PetSupplies"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Photo"}), str: `{"Operation":"","SearchIndex":"Photo"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Shoes"}), str: `{"Operation":"","SearchIndex":"Shoes"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Software"}), str: `{"Operation":"","SearchIndex":"Software"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "SportsAndOutdoors"}), str: `{"Operation":"","SearchIndex":"SportsAndOutdoors"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "ToolsAndHomeImprovement"}), str: `{"Operation":"","SearchIndex":"ToolsAndHomeImprovement"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "ToysAndGames"}), str: `{"Operation":"","SearchIndex":"ToysAndGames"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "VHS"}), str: `{"Operation":"","SearchIndex":"VHS"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "VideoGames"}), str: `{"Operation":"","SearchIndex":"VideoGames"}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "Watches"}), str: `{"Operation":"","SearchIndex":"Watches"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "foo"}), str: `{"Operation":""}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "AvgCustomerReviews"}), str: `{"Operation":"","SortBy":"AvgCustomerReviews"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Featured"}), str: `{"Operation":"","SortBy":"Featured"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "NewestArrivals"}), str: `{"Operation":"","SortBy":"NewestArrivals"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Price:HighToLow"}), str: `{"Operation":"","SortBy":"Price:HighToLow"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Price:LowToHigh"}), str: `{"Operation":"","SortBy":"Price:LowToHigh"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Relevance"}), str: `{"Operation":"","SortBy":"Relevance"}`},
		{q: empty.With().RequestFilters(RequestMap{Title: "foo"}), str: `{"Operation":"","Title":"foo"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestResources(t *testing.T) {
	empty := (*Query)(nil)
	testCases := []struct {
		q   *Query
		str string
	}{
		{q: empty.With().BrowseNodeInfo(), str: `{"Operation":"","Resources":["BrowseNodeInfo.BrowseNodes","BrowseNodeInfo.BrowseNodes.Ancestor","BrowseNodeInfo.BrowseNodes.SalesRank","BrowseNodeInfo.WebsiteSalesRank"]}`},
		{q: empty.With().Images(), str: `{"Operation":"","Resources":["Images.Primary.Small","Images.Primary.Medium","Images.Primary.Large","Images.Variants.Small","Images.Variants.Medium","Images.Variants.Large"]}`},
		{q: empty.With().ItemInfo(), str: `{"Operation":"","Resources":["ItemInfo.ByLineInfo","ItemInfo.ContentInfo","ItemInfo.ContentRating","ItemInfo.Classifications","ItemInfo.ExternalIds","ItemInfo.Features","ItemInfo.ManufactureInfo","ItemInfo.ProductInfo","ItemInfo.TechnicalInfo","ItemInfo.Title","ItemInfo.TradeInInfo"]}`},
		{q: empty.With().Offers(), str: `{"Operation":"","Resources":["Offers.Listings.Availability.MaxOrderQuantity","Offers.Listings.Availability.Message","Offers.Listings.Availability.MinOrderQuantity","Offers.Listings.Availability.Type","Offers.Listings.Condition","Offers.Listings.Condition.SubCondition","Offers.Listings.DeliveryInfo.IsAmazonFulfilled","Offers.Listings.DeliveryInfo.IsFreeShippingEligible","Offers.Listings.DeliveryInfo.IsPrimeEligible","Offers.Listings.DeliveryInfo.ShippingCharges","Offers.Listings.IsBuyBoxWinner","Offers.Listings.LoyaltyPoints.Points","Offers.Listings.MerchantInfo","Offers.Listings.Price","Offers.Listings.ProgramEligibility.IsPrimeExclusive","Offers.Listings.ProgramEligibility.IsPrimePantry","Offers.Listings.Promotions","Offers.Listings.SavingBasis","Offers.Summaries.HighestPrice","Offers.Summaries.LowestPrice","Offers.Summaries.OfferCount"]}`},
		{q: empty.With().OffersV2(), str: `{"Operation":"","Resources":["OffersV2.Listings.Availability","OffersV2.Listings.Condition","OffersV2.Listings.DealDetails","OffersV2.Listings.IsBuyBoxWinner","OffersV2.Listings.LoyaltyPoints","OffersV2.Listings.MerchantInfo","OffersV2.Listings.Price","OffersV2.Listings.Type"]}`},
		{q: empty.With().SearchRefinements(), str: `{"Operation":"","Resources":["SearchRefinements"]}`},
		{q: empty.With().ParentASIN(), str: `{"Operation":"","Resources":["ParentASIN"]}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is \"%v\", want \"%v\"", str, tc.str)
		}
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
