package query

import (
	"testing"
)

func TestGetVariations(t *testing.T) {
	testCases := []struct {
		q   *GetVariations
		str string
	}{
		{q: NewGetVariations("foo.bar", "mytag-20", "Associates"), str: `{"partnerTag":"mytag-20"}`},
		{q: NewGetVariations("foo.bar", "mytag-20", "Associates").ASIN("4900900028"), str: `{"asin":"4900900028","partnerTag":"mytag-20"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetVariations.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestRequestFiltersInGetVariations(t *testing.T) {
	testCases := []struct {
		q   *GetVariations
		str string
	}{
		{q: NewGetVariations("", "", ""), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Actor, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Artist, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(ASIN, "4900900028"), str: `{"asin":"4900900028"}`},
		{q: NewGetVariations("", "", "").Request(Availability, "Available"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Author, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Brand, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(BrowseNodeID, "123"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Condition, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Any"), str: `{"condition":"Any"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "New"), str: `{"condition":"New"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Used"), str: `{"condition":"Used"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Collectible"), str: `{"condition":"Collectible"}`},
		{q: NewGetVariations("", "", "").Request(Condition, "Refurbished"), str: `{"condition":"Refurbished"}`},
		{q: NewGetVariations("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"currencyOfPreference":"foo"}`},
		{q: NewGetVariations("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(ItemIds, "4900900028"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(ItemIdType, "ASIN"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(ItemCount, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(ItemPage, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Keywords, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(BrowseNodeIds, "123"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(BrowseNodeIds, []string{"123", "456"}), str: `{}`},
		{q: NewGetVariations("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"languagesOfPreference":["foo"]}`},
		{q: NewGetVariations("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"languagesOfPreference":["foo","bar"]}`},
		{q: NewGetVariations("", "", "").Request(Marketplace, "foo.bar"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(MaxPrice, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Merchant, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Merchant, "All"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Merchant, "Amazon"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(MinPrice, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(MinReviewsRating, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(MinSavingPercent, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, -1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, 0), str: `{}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, 1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(OfferCount, 123), str: `{}`},
		{q: NewGetVariations("", "", "").Request(PartnerTag, "foo"), str: `{"partnerTag":"foo"}`},
		{q: NewGetVariations("", "", "").Request(PartnerType, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(PartnerType, "Associates"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{}`},
		{q: NewGetVariations("", "", "").Request(SearchIndex, "All"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(Title, "foo"), str: `{}`},
		{q: NewGetVariations("", "", "").Request(VariationCount, 5), str: `{"variationCount":5}`},
		{q: NewGetVariations("", "", "").Request(VariationCount, 11), str: `{}`},
		{q: NewGetVariations("", "", "").Request(VariationCount, -1), str: `{}`},
		{q: NewGetVariations("", "", "").Request(VariationPage, 2), str: `{"variationPage":2}`},
		{q: NewGetVariations("", "", "").Request(VariationPage, 0), str: `{}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("GetVariations.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestResourcesInGetVariations(t *testing.T) {
	testCases := []struct {
		q   *GetVariations
		str string
	}{
		{q: NewGetVariations("", "", "").EnableBrowseNodeInfo(), str: `{"resources":["browseNodeInfo.browseNodes","browseNodeInfo.browseNodes.ancestor","browseNodeInfo.websiteSalesRank"]}`},
		{q: NewGetVariations("", "", "").EnableImages(), str: `{"resources":["images.primary.small","images.primary.medium","images.primary.large","images.variants.small","images.variants.medium","images.variants.large"]}`},
		{q: NewGetVariations("", "", "").EnableItemInfo(), str: `{"resources":["itemInfo.byLineInfo","itemInfo.contentInfo","itemInfo.contentRating","itemInfo.classifications","itemInfo.externalIds","itemInfo.features","itemInfo.manufactureInfo","itemInfo.productInfo","itemInfo.technicalInfo","itemInfo.title","itemInfo.tradeInInfo"]}`},
		{q: NewGetVariations("", "", "").EnableOffers(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: NewGetVariations("", "", "").EnableOffersV2(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: NewGetVariations("", "", "").EnableVariationSummary(), str: `{"resources":["variationSummary.price.highestPrice","variationSummary.price.lowestPrice","variationSummary.variationDimension"]}`},
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
