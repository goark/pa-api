package query

type resource int

const (
	resourceBrowseNodeInfo    resource = 1 + iota //BrowseNodeInfo resource
	resourceImages                                //Images resource
	resourceItemInfo                              //ItemInfo resource
	resourceOffersV2                              //OffersV2 resource
	resourceSearchRefinements                     //SearchRefinements resource
	resourceParentASIN                            //ParentASIN resource
	resourceCustomerReviews                       //CustomerReviews resource
	resourceBrowseNodes                           //BrowseNodes resource
	resourceVariationSummary                      //VariationSummary resource
)

// Resource string values match the Amazon Creators API enum values
// (lowerCamelCase). Sourced from the official Amazon Creators API SDK.
var (
	//BrowseNodeInfo resource
	resourcesBrowseNodeInfo = []string{
		"browseNodeInfo.browseNodes",
		"browseNodeInfo.browseNodes.ancestor",
		"browseNodeInfo.browseNodes.salesRank",
		"browseNodeInfo.websiteSalesRank",
	}
	//Images resource
	resourcesImages = []string{
		"images.primary.small",
		"images.primary.medium",
		"images.primary.large",
		"images.primary.highRes",
		"images.variants.small",
		"images.variants.medium",
		"images.variants.large",
		"images.variants.highRes",
	}
	//ItemInfo resource
	resourcesItemInfo = []string{
		"itemInfo.byLineInfo",
		"itemInfo.contentInfo",
		"itemInfo.contentRating",
		"itemInfo.classifications",
		"itemInfo.externalIds",
		"itemInfo.features",
		"itemInfo.manufactureInfo",
		"itemInfo.productInfo",
		"itemInfo.technicalInfo",
		"itemInfo.title",
		"itemInfo.tradeInInfo",
	}
	//OffersV2 resource
	resourcesOffersV2 = []string{
		"offersV2.listings.availability",
		"offersV2.listings.condition",
		"offersV2.listings.dealDetails",
		"offersV2.listings.isBuyBoxWinner",
		"offersV2.listings.loyaltyPoints",
		"offersV2.listings.merchantInfo",
		"offersV2.listings.price",
		"offersV2.listings.type",
	}
	//SearchRefinements resource
	resourcesSearchRefinements = []string{
		"searchRefinements",
	}
	//ParentASIN resource
	resourcesParentASIN = []string{
		"parentASIN",
	}
	//CustomerReviews resource
	resourcesCustomerReviews = []string{
		"customerReviews.count",
		"customerReviews.starRating",
	}
	//BrowseNodes resource
	resourcesBrowseNodes = []string{
		"browseNodes.ancestor",
		"browseNodes.children",
	}
	//VariationSummary resource
	resourcesVariationSummary = []string{
		"variationSummary.price.highestPrice",
		"variationSummary.price.lowestPrice",
		"variationSummary.variationDimension",
	}

	resourcesMap = map[resource][]string{
		resourceBrowseNodeInfo:    resourcesBrowseNodeInfo,
		resourceImages:            resourcesImages,
		resourceItemInfo:          resourcesItemInfo,
		resourceOffersV2:          resourcesOffersV2,
		resourceSearchRefinements: resourcesSearchRefinements,
		resourceParentASIN:        resourcesParentASIN,
		resourceCustomerReviews:   resourcesCustomerReviews,
		resourceBrowseNodes:       resourcesBrowseNodes,
		resourceVariationSummary:  resourcesVariationSummary,
	}
)

func (r resource) Strings() []string {
	if ss, ok := resourcesMap[r]; ok {
		return ss
	}
	return []string{}
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
