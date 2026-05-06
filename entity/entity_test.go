package entity

import (
	"testing"
)

// TestDecodeGetVariationsResponse exercises the lowerCamelCase Creators API
// response shape for GetVariations, including the variationSummary block
// (pageCount, variationCount, price, variationDimensions).
func TestDecodeGetVariationsResponse(t *testing.T) {
	body := []byte(`{
  "variationsResult": {
    "items": [
      {
        "asin": "B07YCM5K55",
        "parentASIN": "B07YCM5JXX",
        "detailPageURL": "https://www.amazon.com/dp/B07YCM5K55",
        "score": 0.42,
        "images": {
          "primary": {
            "small":  {"url": "https://example.test/s.jpg", "height": 75,  "width": 75},
            "medium": {"url": "https://example.test/m.jpg", "height": 160, "width": 160},
            "large":  {"url": "https://example.test/l.jpg", "height": 500, "width": 500}
          }
        },
        "offersV2": {
          "listings": [
            {
              "isBuyBoxWinner": true,
              "violatesMAP": true,
              "type": "New"
            }
          ]
        }
      }
    ],
    "variationSummary": {
      "pageCount": 3,
      "variationCount": 27,
      "price": {
        "highestPrice": {"amount": 49.99, "currency": "USD", "displayAmount": "$49.99"},
        "lowestPrice":  {"amount": 19.99, "currency": "USD", "displayAmount": "$19.99"}
      },
      "variationDimensions": [
        {"displayName": "Size",  "locale": "en_US", "name": "size_name",  "values": ["S", "M", "L"]},
        {"displayName": "Color", "locale": "en_US", "name": "color_name", "values": ["Red", "Blue"]}
      ]
    }
  }
}`)
	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.VariationsResult == nil {
		t.Fatal("VariationsResult is nil")
	}
	vs := resp.VariationsResult.VariationSummary
	if vs == nil {
		t.Fatal("VariationSummary is nil")
	}
	if vs.PageCount != 3 {
		t.Errorf("PageCount = %d, want 3", vs.PageCount)
	}
	if vs.VariationCount != 27 {
		t.Errorf("VariationCount = %d, want 27", vs.VariationCount)
	}
	if vs.Price == nil || vs.Price.HighestPrice == nil || vs.Price.LowestPrice == nil {
		t.Fatal("price/highestPrice/lowestPrice is nil")
	}
	if got, want := vs.Price.HighestPrice.Amount, 49.99; got != want {
		t.Errorf("HighestPrice.Amount = %v, want %v", got, want)
	}
	if got, want := vs.Price.LowestPrice.DisplayAmount, "$19.99"; got != want {
		t.Errorf("LowestPrice.DisplayAmount = %q, want %q", got, want)
	}
	if got, want := len(vs.VariationDimensions), 2; got != want {
		t.Fatalf("len(VariationDimensions) = %d, want %d", got, want)
	}
	if got, want := vs.VariationDimensions[0].Locale, "en_US"; got != want {
		t.Errorf("VariationDimensions[0].Locale = %q, want %q", got, want)
	}
	if got, want := vs.VariationDimensions[0].Name, "size_name"; got != want {
		t.Errorf("VariationDimensions[0].Name = %q, want %q", got, want)
	}
	// Items: primary image and isBuyBoxWinner / violatesMAP on V2 listing.
	if got, want := len(resp.VariationsResult.Items), 1; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
	item := resp.VariationsResult.Items[0]
	if item.Images == nil || item.Images.Primary == nil || item.Images.Primary.Large == nil {
		t.Fatal("Images.Primary.Large is nil")
	}
	if got, want := item.Images.Primary.Large.URL, "https://example.test/l.jpg"; got != want {
		t.Errorf("Images.Primary.Large.URL = %q, want %q", got, want)
	}
	if item.OffersV2 == nil || item.OffersV2.Listings == nil || len(*item.OffersV2.Listings) != 1 {
		t.Fatal("OffersV2.Listings missing/empty")
	}
	listing := (*item.OffersV2.Listings)[0]
	if !listing.IsBuyboxWinner {
		t.Errorf("isBuyBoxWinner did not decode into IsBuyboxWinner")
	}
	if !listing.ViolatesMAP {
		t.Errorf("violatesMAP did not decode into ViolatesMAP")
	}
}

// TestDecodeGetBrowseNodesResponseWithSalesRank verifies the new SalesRank
// field on the top-level BrowseNodesResult.BrowseNodes.
func TestDecodeGetBrowseNodesResponseWithSalesRank(t *testing.T) {
	body := []byte(`{
  "browseNodesResult": {
    "browseNodes": [
      {
        "id": "3040",
        "displayName": "Books",
        "contextFreeName": "Books",
        "isRoot": true,
        "salesRank": 12345
      }
    ]
  }
}`)
	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.BrowseNodesResult == nil || len(resp.BrowseNodesResult.BrowseNodes) != 1 {
		t.Fatal("BrowseNodesResult.BrowseNodes empty")
	}
	bn := resp.BrowseNodesResult.BrowseNodes[0]
	if bn.SalesRank == nil {
		t.Fatal("BrowseNode.SalesRank is nil")
	}
	if got, want := *bn.SalesRank, 12345; got != want {
		t.Errorf("BrowseNode.SalesRank = %d, want %d", got, want)
	}
}

// TestDecodeBrowseNodeInfoWebsiteSalesRankID exercises the websiteSalesRank.id
// field exposed by the Creators API.
func TestDecodeBrowseNodeInfoWebsiteSalesRankID(t *testing.T) {
	body := []byte(`{
  "itemsResult": {
    "items": [
      {
        "asin": "B0",
        "browseNodeInfo": {
          "browseNodes": [
            {
              "id": "1",
              "displayName": "Books",
              "contextFreeName": "Books",
              "websiteSalesRank": {
                "id": "1000",
                "displayName": "Books",
                "contextFreeName": "Books",
                "salesRank": 7
              }
            }
          ]
        }
      }
    ]
  }
}`)
	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.ItemsResult == nil || len(resp.ItemsResult.Items) == 0 {
		t.Fatal("no items decoded")
	}
	bni := resp.ItemsResult.Items[0].BrowseNodeInfo
	if bni == nil || len(bni.BrowseNodes) == 0 {
		t.Fatal("BrowseNodeInfo.BrowseNodes empty")
	}
	wsr := bni.BrowseNodes[0].WebsiteSalesRank
	if wsr == nil {
		t.Fatal("WebsiteSalesRank is nil")
	}
	if got, want := wsr.Id, "1000"; got != want {
		t.Errorf("WebsiteSalesRank.Id = %q, want %q", got, want)
	}
}

// TestDecodeIgnoresUnknownFields confirms that unknown JSON keys do not
// cause DecodeResponse to error (relevant if the Creators API adds new
// fields between SDK releases).
func TestDecodeIgnoresUnknownFields(t *testing.T) {
	body := []byte(`{"itemsResult":{"items":[{"asin":"A","unknownField":42}]},"newTopLevel":"ok"}`)
	if _, err := DecodeResponse(body); err != nil {
		t.Fatalf("DecodeResponse rejected unknown field: %+v", err)
	}
	// Sanity: malformed JSON still errors.
	if _, err := DecodeResponse([]byte("{not json")); err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}

/* Copyright 2026 goark contributors
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
