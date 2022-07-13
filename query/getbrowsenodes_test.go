package query

import (
	"testing"
)

func TestGetBrowseNodes(t *testing.T) {
	testCases := []struct {
		q   *GetBrowseNodes
		str string
	}{
		{q: NewGetBrowseNodes("foo.bar", "mytag-20", "Associates"), str: `{"Operation":"GetBrowseNodes","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewGetBrowseNodes("foo.bar", "mytag-20", "Associates").BrowseNodeIds([]string{"123"}), str: `{"Operation":"GetBrowseNodes","BrowseNodeIds":["123"],"Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("BrowseNodes.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestRequestFiltersInGetBrowseNodes(t *testing.T) {
	testCases := []struct {
		q   *GetBrowseNodes
		str string
	}{
		{q: NewGetBrowseNodes("", "", ""), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Actor, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Artist, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Availability, "Available"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Author, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Brand, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(BrowseNodeID, "123"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Any"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "New"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Used"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Collectible"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Refurbished"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemIds, "4900900028"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemIdType, "ASIN"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemCount, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemPage, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemPage, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Keywords, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(BrowseNodeIds, "123"), str: `{"Operation":"GetBrowseNodes","BrowseNodeIds":["123"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(BrowseNodeIds, []string{"123", "456"}), str: `{"Operation":"GetBrowseNodes","BrowseNodeIds":["123","456"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"Operation":"GetBrowseNodes","LanguagesOfPreference":["foo"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"Operation":"GetBrowseNodes","LanguagesOfPreference":["foo","bar"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(Marketplace, "foo.bar"), str: `{"Operation":"GetBrowseNodes","Marketplace":"foo.bar"}`},
		{q: NewGetBrowseNodes("", "", "").Request(MaxPrice, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Merchant, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Merchant, "All"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Merchant, "Amazon"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(MinPrice, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(MinReviewsRating, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(MinSavingPercent, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(OfferCount, -1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(OfferCount, 0), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(OfferCount, 1), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(OfferCount, 123), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(PartnerTag, "foo"), str: `{"Operation":"GetBrowseNodes","PartnerTag":"foo"}`},
		{q: NewGetBrowseNodes("", "", "").Request(PartnerType, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(PartnerType, "Associates"), str: `{"Operation":"GetBrowseNodes","PartnerType":"Associates"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(SearchIndex, "All"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{"Operation":"GetBrowseNodes"}`},
		{q: NewGetBrowseNodes("", "", "").Request(Title, "foo"), str: `{"Operation":"GetBrowseNodes"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("BrowseNodes.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestResourcesInGetBrowseNodes(t *testing.T) {
	testCases := []struct {
		q   *GetBrowseNodes
		str string
	}{
		{q: NewGetBrowseNodes("", "", "").EnableBrowseNodes(), str: `{"Operation":"GetBrowseNodes","Resources":["BrowseNodes.Ancestor","BrowseNodes.Children"]}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

/* Copyright 2022 Spiegel and contributors
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
