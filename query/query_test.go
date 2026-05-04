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
		{q: (*Query)(nil).With(), op: paapi5.NullOperation, err: nil, jsn: `{}`},
		{q: New(paapi5.GetItems), op: paapi5.GetItems, err: nil, jsn: `{}`},
	}

	for _, tc := range testCases {
		if op := tc.q.Operation(); op != tc.op {
			t.Errorf("Query.Operation() is %v, want %v", op, tc.op)
		}
		if b, err := tc.q.Payload(); !errors.Is(err, tc.err) {
			t.Errorf("Query.Payload() is %v, want %v", err, tc.err)
		} else if string(b) != tc.jsn {
			t.Errorf("Query.Payload() is %q, want %q", string(b), tc.jsn)
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
		{q: empty.With(), str: `{}`},
		{q: empty.With().RequestFilters(), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{RequestFilter(0): "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{Actor: "foo"}), str: `{"actor":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{Artist: "foo"}), str: `{"artist":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{ASIN: "4900900028"}), str: `{"asin":"4900900028"}`},
		{q: empty.With().RequestFilters(RequestMap{Availability: "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{Availability: "Available"}), str: `{"availability":"Available"}`},
		{q: empty.With().RequestFilters(RequestMap{Availability: "IncludeOutOfStock"}), str: `{"availability":"IncludeOutOfStock"}`},
		{q: empty.With().RequestFilters(RequestMap{Author: "foo"}), str: `{"author":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{Brand: "foo"}), str: `{"brand":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeID: "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeID: "123"}), str: `{"browseNodeId":"123"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Any"}), str: `{"condition":"Any"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "New"}), str: `{"condition":"New"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Used"}), str: `{"condition":"Used"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Collectible"}), str: `{"condition":"Collectible"}`},
		{q: empty.With().RequestFilters(RequestMap{Condition: "Refurbished"}), str: `{"condition":"Refurbished"}`},
		{q: empty.With().RequestFilters(RequestMap{CurrencyOfPreference: "foo"}), str: `{"currencyOfPreference":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{DeliveryFlags: "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{DeliveryFlags: "AmazonGlobal"}), str: `{"deliveryFlags":["AmazonGlobal"]}`},
		{q: empty.With().RequestFilters(RequestMap{DeliveryFlags: []string{"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"}}), str: `{"deliveryFlags":["AmazonGlobal","FreeShipping","FulfilledByAmazon","Prime"]}`},
		{q: empty.With().RequestFilters(RequestMap{ItemIds: "4900900028", ItemIdType: "ASIN"}), str: `{"itemIds":["4900900028"],"itemIdType":"ASIN"}`},
		{q: empty.With().RequestFilters(RequestMap{ItemIds: "4900900028"}, RequestMap{ItemIdType: "ASIN"}), str: `{"itemIds":["4900900028"],"itemIdType":"ASIN"}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: -1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 0}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 1}), str: `{"itemCount":1}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 10}), str: `{"itemCount":10}`},
		{q: empty.With().RequestFilters(RequestMap{ItemCount: 11}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: -1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 0}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 1}), str: `{"itemPage":1}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 10}), str: `{"itemPage":10}`},
		{q: empty.With().RequestFilters(RequestMap{ItemPage: 11}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{Keywords: "foo"}), str: `{"keywords":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeIds: "123"}), str: `{"browseNodeIds":["123"]}`},
		{q: empty.With().RequestFilters(RequestMap{BrowseNodeIds: []string{"123", "456"}}), str: `{"browseNodeIds":["123","456"]}`},
		{q: empty.With().RequestFilters(RequestMap{LanguagesOfPreference: "foo"}), str: `{"languagesOfPreference":["foo"]}`},
		{q: empty.With().RequestFilters(RequestMap{LanguagesOfPreference: []string{"foo", "bar"}}), str: `{"languagesOfPreference":["foo","bar"]}`},
		// Marketplace, Merchant, OfferCount, PartnerType are removed in the
		// Creators API and should not appear in the body.
		{q: empty.With().RequestFilters(RequestMap{Marketplace: "foo.bar"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{Merchant: "All"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{OfferCount: 1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{PartnerType: "Associates"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: -1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: 0}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: 1}), str: `{"maxPrice":1}`},
		{q: empty.With().RequestFilters(RequestMap{MaxPrice: 123}), str: `{"maxPrice":123}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: -1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: 0}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: 1}), str: `{"minPrice":1}`},
		{q: empty.With().RequestFilters(RequestMap{MinPrice: 123}), str: `{"minPrice":123}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: -1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 0}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 1}), str: `{"minReviewsRating":1}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 4}), str: `{"minReviewsRating":4}`},
		{q: empty.With().RequestFilters(RequestMap{MinReviewsRating: 5}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: -1}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 0}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 1}), str: `{"minSavingPercent":1}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 99}), str: `{"minSavingPercent":99}`},
		{q: empty.With().RequestFilters(RequestMap{MinSavingPercent: 100}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{PartnerTag: "foo"}), str: `{"partnerTag":"foo"}`},
		{q: empty.With().RequestFilters(RequestMap{Properties: map[string]string{"foo": "bar"}}), str: `{"properties":{"foo":"bar"}}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{SearchIndex: "All"}), str: `{"searchIndex":"All"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "foo"}), str: `{}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "AvgCustomerReviews"}), str: `{"sortBy":"AvgCustomerReviews"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Featured"}), str: `{"sortBy":"Featured"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "NewestArrivals"}), str: `{"sortBy":"NewestArrivals"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Price:HighToLow"}), str: `{"sortBy":"Price:HighToLow"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Price:LowToHigh"}), str: `{"sortBy":"Price:LowToHigh"}`},
		{q: empty.With().RequestFilters(RequestMap{SortBy: "Relevance"}), str: `{"sortBy":"Relevance"}`},
		{q: empty.With().RequestFilters(RequestMap{Title: "foo"}), str: `{"title":"foo"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestResources(t *testing.T) {
	empty := (*Query)(nil)
	testCases := []struct {
		q   *Query
		str string
	}{
		{q: empty.With().BrowseNodeInfo(), str: `{"resources":["browseNodeInfo.browseNodes","browseNodeInfo.browseNodes.ancestor","browseNodeInfo.browseNodes.salesRank","browseNodeInfo.websiteSalesRank"]}`},
		{q: empty.With().Images(), str: `{"resources":["images.primary.small","images.primary.medium","images.primary.large","images.primary.highRes","images.variants.small","images.variants.medium","images.variants.large","images.variants.highRes"]}`},
		{q: empty.With().ItemInfo(), str: `{"resources":["itemInfo.byLineInfo","itemInfo.contentInfo","itemInfo.contentRating","itemInfo.classifications","itemInfo.externalIds","itemInfo.features","itemInfo.manufactureInfo","itemInfo.productInfo","itemInfo.technicalInfo","itemInfo.title","itemInfo.tradeInInfo"]}`},
		// Offers (V1) is now an alias for OffersV2 in the Creators API.
		{q: empty.With().Offers(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: empty.With().OffersV2(), str: `{"resources":["offersV2.listings.availability","offersV2.listings.condition","offersV2.listings.dealDetails","offersV2.listings.isBuyBoxWinner","offersV2.listings.loyaltyPoints","offersV2.listings.merchantInfo","offersV2.listings.price","offersV2.listings.type"]}`},
		{q: empty.With().SearchRefinements(), str: `{"resources":["searchRefinements"]}`},
		{q: empty.With().ParentASIN(), str: `{"resources":["parentASIN"]}`},
		{q: empty.With().CustomerReviews(), str: `{"resources":["customerReviews.count","customerReviews.starRating"]}`},
		{q: empty.With().BrowseNodes(), str: `{"resources":["browseNodes.ancestor","browseNodes.children"]}`},
		// VariationSummary is no longer exposed by the Creators API.
		{q: empty.With().VariationSummary(), str: `{}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is %q, want %q", str, tc.str)
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
