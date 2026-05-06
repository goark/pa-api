package query

import (
	"testing"
)

func TestGetItems(t *testing.T) {
	testCases := []struct {
		q   *GetItems
		str string
	}{
		// Marketplace and PartnerType are no longer transmitted in the body
		// (the marketplace travels in the `x-marketplace` header and
		// PartnerType is implicit). Only PartnerTag stays in the payload.
		{q: NewGetItems("foo.bar", "mytag-20", "Associates"), str: `{"partnerTag":"mytag-20"}`},
		{q: NewGetItems("foo.bar", "mytag-20", "Associates").ASINs([]string{"4900900028"}), str: `{"itemIds":["4900900028"],"itemIdType":"ASIN","partnerTag":"mytag-20"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetItems.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestRequestFiltersInGetItems(t *testing.T) {
	testCases := []struct {
		q   *GetItems
		str string
	}{
		{q: NewGetItems("", "", ""), str: `{}`},
		{q: NewGetItems("", "", "").Request(Actor, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Artist, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Availability, "Available"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Author, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Brand, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(BrowseNodeID, "123"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Condition, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Condition, "Any"), str: `{"condition":"Any"}`},
		{q: NewGetItems("", "", "").Request(Condition, "New"), str: `{"condition":"New"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Used"), str: `{"condition":"Used"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Collectible"), str: `{"condition":"Collectible"}`},
		{q: NewGetItems("", "", "").Request(Condition, "Refurbished"), str: `{"condition":"Refurbished"}`},
		{q: NewGetItems("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"currencyOfPreference":"foo"}`},
		{q: NewGetItems("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{}`},
		{q: NewGetItems("", "", "").Request(ItemIds, "4900900028"), str: `{"itemIds":["4900900028"]}`},
		{q: NewGetItems("", "", "").Request(ItemIdType, "ASIN"), str: `{"itemIdType":"ASIN"}`},
		{q: NewGetItems("", "", "").Request(ItemCount, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(ItemPage, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(Keywords, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(BrowseNodeIds, "123"), str: `{}`},
		{q: NewGetItems("", "", "").Request(BrowseNodeIds, []string{"123", "456"}), str: `{}`},
		{q: NewGetItems("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"languagesOfPreference":["foo"]}`},
		{q: NewGetItems("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"languagesOfPreference":["foo","bar"]}`},
		// Marketplace, Merchant, OfferCount, PartnerType are no-ops in the
		// Creators API, so they should never appear in the body.
		{q: NewGetItems("", "", "").Request(Marketplace, "foo.bar"), str: `{}`},
		{q: NewGetItems("", "", "").Request(MaxPrice, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(Merchant, "All"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Merchant, "Amazon"), str: `{}`},
		{q: NewGetItems("", "", "").Request(MinPrice, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(MinReviewsRating, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(MinSavingPercent, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(OfferCount, -1), str: `{}`},
		{q: NewGetItems("", "", "").Request(OfferCount, 0), str: `{}`},
		{q: NewGetItems("", "", "").Request(OfferCount, 1), str: `{}`},
		{q: NewGetItems("", "", "").Request(OfferCount, 123), str: `{}`},
		{q: NewGetItems("", "", "").Request(PartnerTag, "foo"), str: `{"partnerTag":"foo"}`},
		{q: NewGetItems("", "", "").Request(PartnerType, "foo"), str: `{}`},
		{q: NewGetItems("", "", "").Request(PartnerType, "Associates"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{}`},
		{q: NewGetItems("", "", "").Request(SearchIndex, "All"), str: `{}`},
		{q: NewGetItems("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{}`},
		{q: NewGetItems("", "", "").Request(Title, "foo"), str: `{}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetItems.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestResourcesInGetItems(t *testing.T) {
	testCases := []struct {
		q   *GetItems
		str string
	}{
		{q: NewGetItems("", "", "").EnableBrowseNodeInfo(), str: `{"resources":["browseNodeInfo.browseNodes","browseNodeInfo.browseNodes.ancestor","browseNodeInfo.browseNodes.salesRank","browseNodeInfo.websiteSalesRank"]}`},
		{q: NewGetItems("", "", "").EnableImages(), str: `{"resources":["images.primary.small","images.primary.medium","images.primary.large","images.variants.small","images.variants.medium","images.variants.large"]}`},
		{q: NewGetItems("", "", "").EnableItemInfo(), str: `{"resources":["itemInfo.byLineInfo","itemInfo.contentInfo","itemInfo.contentRating","itemInfo.classifications","itemInfo.externalIds","itemInfo.features","itemInfo.manufactureInfo","itemInfo.productInfo","itemInfo.technicalInfo","itemInfo.title","itemInfo.tradeInInfo"]}`},
		// EnableOffers (V1) is now an alias for EnableOffersV2.
		{q: NewGetItems("", "", "").EnableOffers(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: NewGetItems("", "", "").EnableOffersV2(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: NewGetItems("", "", "").EnableParentASIN(), str: `{"resources":["parentASIN"]}`},
		{q: NewGetItems("", "", "").EnableCustomerReviews(), str: `{"resources":["customerReviews.count","customerReviews.starRating"]}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is %q, want %q", str, tc.str)
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
