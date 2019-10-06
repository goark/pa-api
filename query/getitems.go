package query

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
	paapi5 "github.com/spiegel-im-spiegel/pa-api"
)

//GetItems is query data class for PA-API v5
type GetItems struct {
	OpeCode              paapi5.Operation `json:"Operation"`
	Marketplace          string
	PartnerTag           string
	PartnerType          string
	ItemIds              []string `json:",omitempty"`
	ItemIdType           string   `json:",omitempty"`
	Resources            []string `json:",omitempty"`
	enableBrowseNodeInfo bool
	enableImages         bool
	enableItemInfo       bool
	enableOffers         bool
	enableParentASIN     bool
}

var _ paapi5.Query = (*GetItems)(nil) //GetItems is compatible with paapi5.Query interface

//New creates new GetItems instance
func NewGetItems(marketplace, partnerTag, partnerType string) *GetItems {
	q := &GetItems{
		OpeCode:        paapi5.GetItems,
		Marketplace:    marketplace,
		PartnerTag:     partnerTag,
		PartnerType:    partnerType,
		enableItemInfo: true,
	}
	return q
}
func newNilGetItems() *GetItems { return NewGetItems("", "", "") }

//ASINs sets ItemIds in GetItems instance
func (q *GetItems) ASINs(itms []string) *GetItems {
	if q.Operation() == paapi5.GetItems {
		if q == nil {
			q = newNilGetItems()
		}
		q.ItemIds = itms
		q.ItemIdType = "ASIN"
	}
	return q
}

//EnableBrowseNodeInfo sets enableBrowseNodeInfo flag in GetItems instance
func (q *GetItems) EnableBrowseNodeInfo(flag bool) *GetItems {
	if q == nil {
		q = newNilGetItems()
	}
	q.enableBrowseNodeInfo = flag
	return q
}

//EnableImages sets enableImages flag in GetItems instance
func (q *GetItems) EnableImages(flag bool) *GetItems {
	if q == nil {
		q = newNilGetItems()
	}
	q.enableImages = flag
	return q
}

//EnableItemInfo sets enableItemInfo flag in GetItems instance
func (q *GetItems) EnableItemInfo(flag bool) *GetItems {
	if q == nil {
		q = newNilGetItems()
	}
	q.enableItemInfo = flag
	return q
}

//EnableOffers sets enableOffers flag in GetItems instance
func (q *GetItems) EnableOffers(flag bool) *GetItems {
	if q == nil {
		q = newNilGetItems()
	}
	q.enableOffers = flag
	return q
}

//EnableParentASIN sets enableParentASIN flag in GetItems instance
func (q *GetItems) EnableParentASIN(flag bool) *GetItems {
	if q == nil {
		q = newNilGetItems()
	}
	q.enableParentASIN = flag
	return q
}

func (q *GetItems) Operation() paapi5.Operation {
	if q == nil {
		return paapi5.NullOperation
	}
	return q.OpeCode
}

func (q *GetItems) Payload() ([]byte, error) {
	if q == nil {
		return nil, errs.Wrap(paapi5.ErrNullPointer, "")
	}
	q.Resources = []string{}
	if q.enableBrowseNodeInfo {
		q.Resources = append(
			q.Resources,
			"BrowseNodeInfo.BrowseNodes",
			"BrowseNodeInfo.BrowseNodes.Ancestor",
			"BrowseNodeInfo.BrowseNodes.SalesRank",
			"BrowseNodeInfo.WebsiteSalesRank",
		)
	}
	if q.enableImages {
		q.Resources = append(
			q.Resources,
			"Images.Primary.Small",
			"Images.Primary.Medium",
			"Images.Primary.Large",
			"Images.Variants.Small",
			"Images.Variants.Medium",
			"Images.Variants.Large",
		)
	}
	if q.enableItemInfo {
		q.Resources = append(
			q.Resources,
			"ItemInfo.ByLineInfo",
			"ItemInfo.ContentInfo",
			"ItemInfo.ContentRating",
			"ItemInfo.Classifications",
			"ItemInfo.ExternalIds",
			"ItemInfo.Features",
			"ItemInfo.ManufactureInfo",
			"ItemInfo.ProductInfo",
			"ItemInfo.TechnicalInfo",
			"ItemInfo.Title",
			"ItemInfo.TradeInInfo",
		)
	}
	// if q.enableOffers {
	// 	q.Resources = append(
	// 		q.Resources,
	// 		"Offers.Listings.Availability.MaxOrderQuantity",
	// 		"Offers.Listings.Availability.Message",
	// 		"Offers.Listings.Availability.MinOrderQuantity",
	// 		"Offers.Listings.Availability.Type",
	// 		"Offers.Listings.Condition",
	// 		"Offers.Listings.Condition.SubCondition",
	// 		"Offers.Listings.DeliveryInfo.IsAmazonFulfilled",
	// 		"Offers.Listings.DeliveryInfo.IsFreeShippingEligible",
	// 		"Offers.Listings.DeliveryInfo.IsPrimeEligible",
	// 		"Offers.Listings.DeliveryInfo.ShippingCharges",
	// 		"Offers.Listings.IsBuyBoxWinner",
	// 		"Offers.Listings.LoyaltyPoints.Points",
	// 		"Offers.Listings.MerchantInfo",
	// 		"Offers.Listings.Price",
	// 		"Offers.Listings.ProgramEligibility.IsPrimeExclusive",
	// 		"Offers.Listings.ProgramEligibility.IsPrimePantry",
	// 		"Offers.Listings.Promotions",
	// 		"Offers.Listings.SavingBasis",
	// 		"Offers.Summaries.HighestPrice",
	// 		"Offers.Summaries.LowestPrice",
	// 		"Offers.Summaries.OfferCount",
	// 	)
	// }
	if q.enableParentASIN {
		q.Resources = append(
			q.Resources,
			"ParentASIN",
		)
	}
	b, err := json.Marshal(q)
	return b, errs.Wrap(err, "")
}

func (q *GetItems) String() string {
	b, err := q.Payload()
	if err != nil {
		return ""
	}
	return string(b)
}

/* Copyright 2019 Spiegel
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
