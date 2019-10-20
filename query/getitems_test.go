package query

import (
	"testing"
)

func TestGetItems(t *testing.T) {
	testCases := []struct {
		q   *GetItems
		str string
	}{
		{q: NewGetItems("foo.bar", "mytag-20", "Associates"), str: `{"Operation":"GetItems","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewGetItems("foo.bar", "mytag-20", "Associates").ASINs([]string{"4900900028"}), str: `{"Operation":"GetItems","ItemIds":["4900900028"],"ItemIdType":"ASIN","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetItems.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestRequestFiltersInGetItems(t *testing.T) {
	testCases := []struct {
		q   *GetItems
		str string
	}{
		{q: NewGetItems("", "", ""), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Actor, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Artist, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Availability, "Available"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Author, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Brand, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(BrowseNodeID, "123"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Condition, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Any"), str: `{"Operation":"GetItems","Condition":"Any"}`},
		{q: NewGetItems("", "", "").Request(Condition, "New"), str: `{"Operation":"GetItems","Condition":"New"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Used"), str: `{"Operation":"GetItems","Condition":"Used"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Collectible"), str: `{"Operation":"GetItems","Condition":"Collectible"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Refurbished"), str: `{"Operation":"GetItems","Condition":"Refurbished"}`},
		{q: NewGetItems("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"Operation":"GetItems","CurrencyOfPreference":"foo"}`},
		{q: NewGetItems("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(ItemIds, "4900900028"), str: `{"Operation":"GetItems","ItemIds":["4900900028"]}`},
		{q: NewGetItems("", "", "").Request(ItemIdType, "ASIN"), str: `{"Operation":"GetItems","ItemIdType":"ASIN"}`},
		{q: NewGetItems("", "", "").Request(ItemCount, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(ItemPage, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(ItemPage, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Keywords, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"Operation":"GetItems","LanguagesOfPreference":["foo"]}`},
		{q: NewGetItems("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"Operation":"GetItems","LanguagesOfPreference":["foo","bar"]}`},
		{q: NewGetItems("", "", "").Request(Marketplace, "foo.bar"), str: `{"Operation":"GetItems","Marketplace":"foo.bar"}`},
		{q: NewGetItems("", "", "").Request(MaxPrice, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Merchant, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Merchant, "All"), str: `{"Operation":"GetItems","Merchant":"All"}`},
		{q: NewGetItems("", "", "").Request(Merchant, "Amazon"), str: `{"Operation":"GetItems","Merchant":"Amazon"}`},
		{q: NewGetItems("", "", "").Request(MinPrice, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(MinReviewsRating, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(MinSavingPercent, 1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(OfferCount, -1), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(OfferCount, 0), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(OfferCount, 1), str: `{"Operation":"GetItems","OfferCount":1}`},
		{q: NewGetItems("", "", "").Request(OfferCount, 123), str: `{"Operation":"GetItems","OfferCount":123}`},
		{q: NewGetItems("", "", "").Request(PartnerTag, "foo"), str: `{"Operation":"GetItems","PartnerTag":"foo"}`},
		{q: NewGetItems("", "", "").Request(PartnerType, "foo"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(PartnerType, "Associates"), str: `{"Operation":"GetItems","PartnerType":"Associates"}`},
		{q: NewGetItems("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(SearchIndex, "All"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{"Operation":"GetItems"}`},
		{q: NewGetItems("", "", "").Request(Title, "foo"), str: `{"Operation":"GetItems"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetItems.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestResourcesInGetItems(t *testing.T) {
	testCases := []struct {
		q   *GetItems
		str string
	}{
		{q: NewGetItems("", "", "").EnableBrowseNodeInfo(), str: `{"Operation":"GetItems","Resources":["BrowseNodeInfo.BrowseNodes","BrowseNodeInfo.BrowseNodes.Ancestor","BrowseNodeInfo.BrowseNodes.SalesRank","BrowseNodeInfo.WebsiteSalesRank"]}`},
		{q: NewGetItems("", "", "").EnableImages(), str: `{"Operation":"GetItems","Resources":["Images.Primary.Small","Images.Primary.Medium","Images.Primary.Large","Images.Variants.Small","Images.Variants.Medium","Images.Variants.Large"]}`},
		{q: NewGetItems("", "", "").EnableItemInfo(), str: `{"Operation":"GetItems","Resources":["ItemInfo.ByLineInfo","ItemInfo.ContentInfo","ItemInfo.ContentRating","ItemInfo.Classifications","ItemInfo.ExternalIds","ItemInfo.Features","ItemInfo.ManufactureInfo","ItemInfo.ProductInfo","ItemInfo.TechnicalInfo","ItemInfo.Title","ItemInfo.TradeInInfo"]}`},
		{q: NewGetItems("", "", "").EnableOffers(), str: `{"Operation":"GetItems","Resources":["Offers.Listings.Availability.MaxOrderQuantity","Offers.Listings.Availability.Message","Offers.Listings.Availability.MinOrderQuantity","Offers.Listings.Availability.Type","Offers.Listings.Condition","Offers.Listings.Condition.SubCondition","Offers.Listings.DeliveryInfo.IsAmazonFulfilled","Offers.Listings.DeliveryInfo.IsFreeShippingEligible","Offers.Listings.DeliveryInfo.IsPrimeEligible","Offers.Listings.DeliveryInfo.ShippingCharges","Offers.Listings.IsBuyBoxWinner","Offers.Listings.LoyaltyPoints.Points","Offers.Listings.MerchantInfo","Offers.Listings.Price","Offers.Listings.ProgramEligibility.IsPrimeExclusive","Offers.Listings.ProgramEligibility.IsPrimePantry","Offers.Listings.Promotions","Offers.Listings.SavingBasis","Offers.Summaries.HighestPrice","Offers.Summaries.LowestPrice","Offers.Summaries.OfferCount"]}`},
		{q: NewGetItems("", "", "").EnableParentASIN(), str: `{"Operation":"GetItems","Resources":["ParentASIN"]}`},
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
