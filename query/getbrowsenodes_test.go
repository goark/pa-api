package query

import (
	"testing"
)

func TestGetBrowseNodes(t *testing.T) {
	testCases := []struct {
		q   *GetBrowseNodes
		str string
	}{
		{q: NewGetBrowseNodes("foo.bar", "mytag-20", "Associates"), str: `{"partnerTag":"mytag-20"}`},
		{q: NewGetBrowseNodes("foo.bar", "mytag-20", "Associates").BrowseNodeIds([]string{"123"}), str: `{"browseNodeIds":["123"],"partnerTag":"mytag-20"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("BrowseNodes.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestRequestFiltersInGetBrowseNodes(t *testing.T) {
	testCases := []struct {
		q   *GetBrowseNodes
		str string
	}{
		{q: NewGetBrowseNodes("", "", ""), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Actor, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Artist, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Availability, "Available"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Author, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Brand, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(BrowseNodeID, "123"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Any"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "New"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Used"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Collectible"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Condition, "Refurbished"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(CurrencyOfPreference, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemIds, "4900900028"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemIdType, "ASIN"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemCount, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(ItemPage, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Keywords, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(BrowseNodeIds, "123"), str: `{"browseNodeIds":["123"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(BrowseNodeIds, []string{"123", "456"}), str: `{"browseNodeIds":["123","456"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"languagesOfPreference":["foo"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"languagesOfPreference":["foo","bar"]}`},
		{q: NewGetBrowseNodes("", "", "").Request(Marketplace, "foo.bar"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(MaxPrice, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Merchant, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Merchant, "All"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Merchant, "Amazon"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(MinPrice, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(MinReviewsRating, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(MinSavingPercent, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(OfferCount, 1), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(PartnerTag, "foo"), str: `{"partnerTag":"foo"}`},
		{q: NewGetBrowseNodes("", "", "").Request(PartnerType, "foo"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(PartnerType, "Associates"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(SearchIndex, "All"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{}`},
		{q: NewGetBrowseNodes("", "", "").Request(Title, "foo"), str: `{}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("BrowseNodes.String() is %q, want %q", str, tc.str)
		}
	}
}

func TestResourcesInGetBrowseNodes(t *testing.T) {
	testCases := []struct {
		q   *GetBrowseNodes
		str string
	}{
		{q: NewGetBrowseNodes("", "", "").EnableBrowseNodes(), str: `{"resources":["browseNodes.ancestor","browseNodes.children"]}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is %q, want %q", str, tc.str)
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
