package query

import (
	"testing"
)

func TestGetVariations(t *testing.T) {
	testCases := []struct {
		q   *GetVariations
		str string
	}{
		{q: NewGetVariations("foo.bar", "mytag-20", "Associates"), str: `{"Operation":"GetVariations","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewGetVariations("foo.bar", "mytag-20", "Associates").ASIN("4900900028"), str: `{"Operation":"GetVariations","ASIN":"4900900028","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetVariations.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestRequestFiltersInGetVariations(t *testing.T) {
	testCases := []struct {
		q   *GetVariations
		str string
	}{
		{q: NewGetVariations("", "", ""), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Actor, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Artist, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(ASIN, "4900900028"), str: `{"Operation":"GetVariations","ASIN":"4900900028"}`},
		{q: NewGetVariations("", "", "").Request(Availability, "Available"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Author, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Brand, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(BrowseNodeID, "123"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Any"), str: `{"Operation":"GetVariations","Condition":"Any"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "New"), str: `{"Operation":"GetVariations","Condition":"New"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Used"), str: `{"Operation":"GetVariations","Condition":"Used"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Collectible"), str: `{"Operation":"GetVariations","Condition":"Collectible"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Refurbished"), str: `{"Operation":"GetVariations","Condition":"Refurbished"}`},
		{q: NewGetVariations("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"Operation":"GetVariations","CurrencyOfPreference":"foo"}`},
		{q: NewGetVariations("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(ItemIds, "4900900028"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(ItemIdType, "ASIN"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(ItemCount, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(ItemPage, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(ItemPage, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Keywords, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(BrowseNodeIds, "123"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(BrowseNodeIds, []string{"123", "456"}), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"Operation":"GetVariations","LanguagesOfPreference":["foo"]}`},
		{q: NewGetVariations("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"Operation":"GetVariations","LanguagesOfPreference":["foo","bar"]}`},
		{q: NewGetVariations("", "", "").Request(Marketplace, "foo.bar"), str: `{"Operation":"GetVariations","Marketplace":"foo.bar"}`},
		{q: NewGetVariations("", "", "").Request(MaxPrice, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Merchant, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Merchant, "All"), str: `{"Operation":"GetVariations","Merchant":"All"}`},
		{q: NewGetVariations("", "", "").Request(Merchant, "Amazon"), str: `{"Operation":"GetVariations","Merchant":"Amazon"}`},
		{q: NewGetVariations("", "", "").Request(MinPrice, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(MinReviewsRating, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(MinSavingPercent, 1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, -1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, 0), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, 1), str: `{"Operation":"GetVariations","OfferCount":1}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, 123), str: `{"Operation":"GetVariations","OfferCount":123}`},
		{q: NewGetVariations("", "", "").Request(PartnerTag, "foo"), str: `{"Operation":"GetVariations","PartnerTag":"foo"}`},
		{q: NewGetVariations("", "", "").Request(PartnerType, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(PartnerType, "Associates"), str: `{"Operation":"GetVariations","PartnerType":"Associates"}`},
		{q: NewGetVariations("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(SearchIndex, "All"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(Title, "foo"), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(VariationCount, 5), str: `{"Operation":"GetVariations","VariationCount":5}`},
		{q: NewGetVariations("", "", "").Request(VariationCount, 11), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(VariationCount, -1), str: `{"Operation":"GetVariations"}`},
		{q: NewGetVariations("", "", "").Request(VariationPage, 2), str: `{"Operation":"GetVariations","VariationPage":2}`},
		{q: NewGetVariations("", "", "").Request(VariationPage, 0), str: `{"Operation":"GetVariations"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetVariations.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestResourcesInGetVariations(t *testing.T) {
	testCases := []struct {
		q   *GetVariations
		str string
	}{
		{q: NewGetVariations("", "", "").EnableBrowseNodeInfo(), str: `{"Operation":"GetVariations","Resources":["BrowseNodeInfo.BrowseNodes","BrowseNodeInfo.BrowseNodes.Ancestor","BrowseNodeInfo.BrowseNodes.SalesRank","BrowseNodeInfo.WebsiteSalesRank"]}`},
		{q: NewGetVariations("", "", "").EnableImages(), str: `{"Operation":"GetVariations","Resources":["Images.Primary.Small","Images.Primary.Medium","Images.Primary.Large","Images.Variants.Small","Images.Variants.Medium","Images.Variants.Large"]}`},
		{q: NewGetVariations("", "", "").EnableItemInfo(), str: `{"Operation":"GetVariations","Resources":["ItemInfo.ByLineInfo","ItemInfo.ContentInfo","ItemInfo.ContentRating","ItemInfo.Classifications","ItemInfo.ExternalIds","ItemInfo.Features","ItemInfo.ManufactureInfo","ItemInfo.ProductInfo","ItemInfo.TechnicalInfo","ItemInfo.Title","ItemInfo.TradeInInfo"]}`},
		{q: NewGetVariations("", "", "").EnableOffers(), str: `{"Operation":"GetVariations","Resources":["Offers.Listings.Availability.MaxOrderQuantity","Offers.Listings.Availability.Message","Offers.Listings.Availability.MinOrderQuantity","Offers.Listings.Availability.Type","Offers.Listings.Condition","Offers.Listings.Condition.SubCondition","Offers.Listings.DeliveryInfo.IsAmazonFulfilled","Offers.Listings.DeliveryInfo.IsFreeShippingEligible","Offers.Listings.DeliveryInfo.IsPrimeEligible","Offers.Listings.DeliveryInfo.ShippingCharges","Offers.Listings.IsBuyBoxWinner","Offers.Listings.LoyaltyPoints.Points","Offers.Listings.MerchantInfo","Offers.Listings.Price","Offers.Listings.ProgramEligibility.IsPrimeExclusive","Offers.Listings.ProgramEligibility.IsPrimePantry","Offers.Listings.Promotions","Offers.Listings.SavingBasis","Offers.Summaries.HighestPrice","Offers.Summaries.LowestPrice","Offers.Summaries.OfferCount"]}`},
		{q: NewGetVariations("", "", "").EnableVariationSummary(), str: `{"Operation":"GetVariations","Resources":["VariationSummary.Price.HighestPrice","VariationSummary.Price.LowestPrice","VariationSummary.VariationDimension"]}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

/* Copyright 2019-2022 Spiegel and contributors
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
