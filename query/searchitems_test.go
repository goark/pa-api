package query

import "testing"

func TestSearchItems(t *testing.T) {
	testCases := []struct {
		q   *SearchItems
		str string
	}{
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates"), str: `{"partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(ItemIds, "foo"), str: `{"partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Actor, "foo"), str: `{"actor":"foo","partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Artist, "foo"), str: `{"artist":"foo","partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Author, "foo"), str: `{"author":"foo","partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Brand, "foo"), str: `{"brand":"foo","partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Keywords, "foo"), str: `{"keywords":"foo","partnerTag":"mytag-20"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Title, "foo"), str: `{"partnerTag":"mytag-20","title":"foo"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("SearchItems.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestRequestInSearchItems(t *testing.T) {
	testCases := []struct {
		q   *SearchItems
		str string
	}{
		{q: NewSearchItems("", "", ""), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Actor, "foo"), str: `{"actor":"foo"}`},
		{q: NewSearchItems("", "", "").Request(Artist, "foo"), str: `{"artist":"foo"}`},
		{q: NewSearchItems("", "", "").Request(Availability, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Availability, "Available"), str: `{"availability":"Available"}`},
		{q: NewSearchItems("", "", "").Request(Availability, "IncludeOutOfStock"), str: `{"availability":"IncludeOutOfStock"}`},
		{q: NewSearchItems("", "", "").Request(Author, "foo"), str: `{"author":"foo"}`},
		{q: NewSearchItems("", "", "").Request(Brand, "foo"), str: `{"brand":"foo"}`},
		{q: NewSearchItems("", "", "").Request(BrowseNodeID, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(BrowseNodeID, "123"), str: `{"browseNodeId":"123"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Any"), str: `{"condition":"Any"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "New"), str: `{"condition":"New"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Used"), str: `{"condition":"Used"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Collectible"), str: `{"condition":"Collectible"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Refurbished"), str: `{"condition":"Refurbished"}`},
		{q: NewSearchItems("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"currencyOfPreference":"foo"}`},
		{q: NewSearchItems("", "", "").Request(DeliveryFlags, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{"deliveryFlags":["AmazonGlobal"]}`},
		{q: NewSearchItems("", "", "").Request(DeliveryFlags, []string{"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"}), str: `{"deliveryFlags":["AmazonGlobal","FreeShipping","FulfilledByAmazon","Prime"]}`},
		{q: NewSearchItems("", "", "").Request(ItemIds, "4900900028"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemIdType, "ASIN"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 1), str: `{"itemCount":1}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 10), str: `{"itemCount":10}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 11), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 1), str: `{"itemPage":1}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 10), str: `{"itemPage":10}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 11), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Keywords, "foo"), str: `{"keywords":"foo"}`},
		{q: NewSearchItems("", "", "").Request(BrowseNodeIds, "123"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(BrowseNodeIds, []string{"123", "456"}), str: `{}`},
		{q: NewSearchItems("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"languagesOfPreference":["foo"]}`},
		{q: NewSearchItems("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"languagesOfPreference":["foo","bar"]}`},
		{q: NewSearchItems("", "", "").Request(Marketplace, "foo.bar"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, 1), str: `{"maxPrice":1}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, 123), str: `{"maxPrice":123}`},
		{q: NewSearchItems("", "", "").Request(Merchant, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Merchant, "All"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Merchant, "Amazon"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, 1), str: `{"minPrice":1}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, 123), str: `{"minPrice":123}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 1), str: `{"minReviewsRating":1}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 4), str: `{"minReviewsRating":4}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 5), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 1), str: `{"minSavingPercent":1}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 99), str: `{"minSavingPercent":99}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 100), str: `{}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, -1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, 0), str: `{}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, 1), str: `{}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, 123), str: `{}`},
		{q: NewSearchItems("", "", "").Request(PartnerTag, "foo"), str: `{"partnerTag":"foo"}`},
		{q: NewSearchItems("", "", "").Request(PartnerType, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(PartnerType, "Associates"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{"properties":{"foo":"bar"}}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "All"), str: `{"searchIndex":"All"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "AmazonVideo"), str: `{"searchIndex":"AmazonVideo"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Apparel"), str: `{"searchIndex":"Apparel"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Books"), str: `{"searchIndex":"Books"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Watches"), str: `{"searchIndex":"Watches"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "foo"), str: `{}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{"sortBy":"AvgCustomerReviews"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Featured"), str: `{"sortBy":"Featured"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "NewestArrivals"), str: `{"sortBy":"NewestArrivals"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Price:HighToLow"), str: `{"sortBy":"Price:HighToLow"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Price:LowToHigh"), str: `{"sortBy":"Price:LowToHigh"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Relevance"), str: `{"sortBy":"Relevance"}`},
		{q: NewSearchItems("", "", "").Request(Title, "foo"), str: `{"title":"foo"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestResourcesInSearchItems(t *testing.T) {
	testCases := []struct {
		q   *SearchItems
		str string
	}{
		{q: NewSearchItems("", "", "").EnableBrowseNodeInfo(), str: `{"resources":["browseNodeInfo.browseNodes","browseNodeInfo.browseNodes.ancestor","browseNodeInfo.websiteSalesRank"]}`},
		{q: NewSearchItems("", "", "").EnableImages(), str: `{"resources":["images.primary.small","images.primary.medium","images.primary.large","images.variants.small","images.variants.medium","images.variants.large"]}`},
		{q: NewSearchItems("", "", "").EnableItemInfo(), str: `{"resources":["itemInfo.byLineInfo","itemInfo.contentInfo","itemInfo.contentRating","itemInfo.classifications","itemInfo.externalIds","itemInfo.features","itemInfo.manufactureInfo","itemInfo.productInfo","itemInfo.technicalInfo","itemInfo.title","itemInfo.tradeInInfo"]}`},
		{q: NewSearchItems("", "", "").EnableOffers(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: NewSearchItems("", "", "").EnableOffersV2(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: NewSearchItems("", "", "").EnableSearchRefinements(), str: `{"resources":["searchRefinements"]}`},
		{q: NewSearchItems("", "", "").EnableParentASIN(), str: `{"resources":["parentASIN"]}`},
		{q: NewSearchItems("", "", "").EnableCustomerReviews(), str: `{"resources":["customerReviews.count","customerReviews.starRating"]}`},
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
