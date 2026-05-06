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
        {"displayName": "Size",  "name": "size_name",  "values": ["S", "M", "L"]},
        {"displayName": "Color", "name": "color_name", "values": ["Red", "Blue"]}
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
	if got, want := vs.VariationDimensions[0].Name, "size_name"; got != want {
		t.Errorf("VariationDimensions[0].Name = %q, want %q", got, want)
	}
	// Items: primary image and isBuyBoxWinner on V2 listing.
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
}

// TestDecodeRealGetBrowseNodesSampleResponse pins the decode contract for
// GetBrowseNodes against a sample captured from the live Creators API.
// Exercises both the recursive `ancestor` chain (Mexico → Explore the World
// → Geography & Cultures → Children's Books → Subjects → Books) and the
// `children[]` list returned for root nodes. The top-level container is
// `browseNodesResult` (browseNodes plural, result singular) — matches the
// shape used by GetVariations rather than the unique `itemResults` plural
// used only by GetItems.
func TestDecodeRealGetBrowseNodesSampleResponse(t *testing.T) {
	body := []byte(`{
  "browseNodesResult": {
    "browseNodes": [
      {
        "ancestor": {
          "ancestor": {
            "ancestor": {
              "ancestor": {
                "ancestor": {
                  "contextFreeName": "Books",
                  "displayName": "Books",
                  "id": "283155"
                },
                "contextFreeName": "Subjects",
                "displayName": "Subjects",
                "id": "1000"
              },
              "contextFreeName": "Children's Books",
              "displayName": "Children's Books",
              "id": "4"
            },
            "contextFreeName": "Children's Geography & Cultures Books",
            "displayName": "Geography & Cultures",
            "id": "3344091011"
          },
          "contextFreeName": "Children's Explore the World Books",
          "displayName": "Explore the World",
          "id": "3023"
        },
        "contextFreeName": "Children's Mexico Books",
        "displayName": "Mexico",
        "id": "3040",
        "isRoot": false
      },
      {
        "children": [
          {"contextFreeName": "Subjects",                  "displayName": "Subjects",                  "id": "1000"},
          {"contextFreeName": "Books Featured Categories", "displayName": "Books Featured Categories", "id": "51546011"},
          {"contextFreeName": "Specialty Boutique",        "displayName": "Specialty Boutique",        "id": "2349030011"}
        ],
        "contextFreeName": "Books",
        "displayName": "Books",
        "id": "283155",
        "isRoot": true
      }
    ]
  }
}`)
	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.BrowseNodesResult == nil {
		t.Fatal("BrowseNodesResult is nil")
	}
	if got, want := len(resp.BrowseNodesResult.BrowseNodes), 2; got != want {
		t.Fatalf("len(BrowseNodes) = %d, want %d", got, want)
	}

	// First node: Mexico, with a 5-level ancestor chain ending at the
	// Books root.
	mexico := resp.BrowseNodesResult.BrowseNodes[0]
	if got, want := mexico.Id, "3040"; got != want {
		t.Errorf("Mexico.Id = %q, want %q", got, want)
	}
	if got, want := mexico.DisplayName, "Mexico"; got != want {
		t.Errorf("Mexico.DisplayName = %q, want %q", got, want)
	}
	if mexico.IsRoot {
		t.Errorf("Mexico.IsRoot = true, want false")
	}
	// Walk the ancestor chain.
	expected := []struct{ id, name string }{
		{"3023", "Explore the World"},
		{"3344091011", "Geography & Cultures"},
		{"4", "Children's Books"},
		{"1000", "Subjects"},
		{"283155", "Books"},
	}
	cur := mexico.Ancestor
	for i, want := range expected {
		if cur == nil {
			t.Fatalf("ancestor chain truncated at depth %d (expected %+v)", i, want)
		}
		if cur.Id != want.id || cur.DisplayName != want.name {
			t.Errorf("ancestor[%d] = (%q, %q), want (%q, %q)", i, cur.Id, cur.DisplayName, want.id, want.name)
		}
		cur = cur.Ancestor
	}
	if cur != nil {
		t.Errorf("ancestor chain has extra node beyond root: %+v", cur)
	}

	// Second node: Books root, with three children.
	books := resp.BrowseNodesResult.BrowseNodes[1]
	if got, want := books.Id, "283155"; got != want {
		t.Errorf("Books.Id = %q, want %q", got, want)
	}
	if !books.IsRoot {
		t.Errorf("Books.IsRoot = false, want true")
	}
	if got, want := len(books.Children), 3; got != want {
		t.Fatalf("len(Books.Children) = %d, want %d", got, want)
	}
	if got, want := books.Children[2].DisplayName, "Specialty Boutique"; got != want {
		t.Errorf("Books.Children[2].DisplayName = %q, want %q", got, want)
	}
	if got, want := books.Children[1].Id, "51546011"; got != want {
		t.Errorf("Books.Children[1].Id = %q, want %q", got, want)
	}
}

// TestDecodeBrowseNodeInfoWebsiteSalesRank exercises the websiteSalesRank
// nested block carried inside an item's browseNodeInfo.
func TestDecodeBrowseNodeInfoWebsiteSalesRank(t *testing.T) {
	body := []byte(`{
  "itemResults": {
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
	if got, want := wsr.SalesRank, 7; got != want {
		t.Errorf("WebsiteSalesRank.SalesRank = %d, want %d", got, want)
	}
}

// TestDecodeIgnoresUnknownFields confirms that unknown JSON keys do not
// cause DecodeResponse to error (relevant if the Creators API adds new
// fields between SDK releases).
func TestDecodeIgnoresUnknownFields(t *testing.T) {
	body := []byte(`{"itemResults":{"items":[{"asin":"A","unknownField":42}]},"newTopLevel":"ok"}`)
	if _, err := DecodeResponse(body); err != nil {
		t.Fatalf("DecodeResponse rejected unknown field: %+v", err)
	}
	// Sanity: malformed JSON still errors.
	if _, err := DecodeResponse([]byte("{not json")); err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}

// TestDecodeRealGetVariationsSampleResponse pins the decode contract for
// GetVariations against a sample captured from the live Creators API.
// Notably:
//   - the top-level container is `variationsResult` (variations plural,
//     result singular) — this differs from GetItems' `itemResults`
//     (item singular, results plural).
//   - the `Price` block in `variationSummary` is sometimes returned with a
//     capitalised key; case-insensitive decode covers it.
//   - `variationDimensions[]` does NOT include a `locale` field; ensures
//     the entity stays minimal.
func TestDecodeRealGetVariationsSampleResponse(t *testing.T) {
	body := []byte(`{
  "variationsResult": {
    "items": [
      {
        "asin": "B019MNBMS4",
        "detailPageURL": "https://www.amazon.co.uk/dp/B019MNBMS4?tag=xyz-20&linkCode=ogv&th=1&psc=1",
        "itemInfo": {
          "title": {
            "displayValue": "Tommy Hilfiger Men's Ranger Leather Passcase Wallet",
            "label": "title",
            "locale": "en_GB"
          }
        },
        "variationAttributes": [
          {"name": "size_name",  "value": "One Size"},
          {"name": "color_name", "value": "Navy"}
        ]
      },
      {
        "asin": "B073211XCB",
        "detailPageURL": "https://www.amazon.co.uk/dp/B073211XCB",
        "itemInfo": {
          "title": {
            "displayValue": "Tommy Hilfiger mens RFID Blocking Wallet",
            "label": "title",
            "locale": "en_GB"
          }
        },
        "variationAttributes": [
          {"name": "size_name",  "value": "One Size"},
          {"name": "color_name", "value": "Logan - Tan"}
        ]
      }
    ],
    "variationSummary": {
      "pageCount": 2,
      "Price": {
        "highestPrice": {"amount": 30.87, "currency": "GBP", "displayAmount": "£30.87"},
        "lowestPrice":  {"amount": 17.03, "currency": "GBP", "displayAmount": "£17.03"}
      },
      "variationCount": 13,
      "variationDimensions": [
        {"displayName": "Size",   "name": "size_name",  "values": ["One Size"]},
        {"displayName": "Colour", "name": "color_name", "values": ["Brown", "Navy", "Black", "Burgundy", "Cognac", "Gray", "Green", "Logan - Tan", "Navy/Black", "Red", "Rfid-black", "Rfid-navy", "Tan"]}
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
	// Items decode end-to-end with title + variationAttributes.
	if got, want := len(resp.VariationsResult.Items), 2; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
	first := resp.VariationsResult.Items[0]
	if got, want := first.ASIN, "B019MNBMS4"; got != want {
		t.Errorf("Items[0].ASIN = %q, want %q", got, want)
	}
	if first.ItemInfo == nil || first.ItemInfo.Title == nil {
		t.Fatal("Items[0].ItemInfo.Title is nil")
	}
	if got, want := first.ItemInfo.Title.Locale, "en_GB"; got != want {
		t.Errorf("Items[0].ItemInfo.Title.Locale = %q, want %q", got, want)
	}
	if got, want := len(first.VariationAttributes), 2; got != want {
		t.Fatalf("Items[0].VariationAttributes len = %d, want %d", got, want)
	}
	if got, want := first.VariationAttributes[1].Name, "color_name"; got != want {
		t.Errorf("Items[0].VariationAttributes[1].Name = %q, want %q", got, want)
	}
	if got, want := first.VariationAttributes[1].Value, "Navy"; got != want {
		t.Errorf("Items[0].VariationAttributes[1].Value = %q, want %q", got, want)
	}
	// VariationSummary: pageCount/variationCount/Price/variationDimensions.
	vs := resp.VariationsResult.VariationSummary
	if vs == nil {
		t.Fatal("VariationSummary is nil")
	}
	if vs.PageCount != 2 {
		t.Errorf("VariationSummary.PageCount = %d, want 2", vs.PageCount)
	}
	if vs.VariationCount != 13 {
		t.Errorf("VariationSummary.VariationCount = %d, want 13", vs.VariationCount)
	}
	if vs.Price == nil || vs.Price.HighestPrice == nil || vs.Price.LowestPrice == nil {
		t.Fatal("VariationSummary.Price.{HighestPrice,LowestPrice} missing (capitalized 'Price' key did not decode)")
	}
	if got, want := vs.Price.HighestPrice.Currency, "GBP"; got != want {
		t.Errorf("Highest currency = %q, want %q", got, want)
	}
	if got, want := vs.Price.LowestPrice.DisplayAmount, "£17.03"; got != want {
		t.Errorf("Lowest displayAmount = %q, want %q", got, want)
	}
	if got, want := len(vs.VariationDimensions), 2; got != want {
		t.Fatalf("len(VariationDimensions) = %d, want %d", got, want)
	}
	if got, want := vs.VariationDimensions[1].DisplayName, "Colour"; got != want {
		t.Errorf("VariationDimensions[1].DisplayName = %q, want %q", got, want)
	}
	if got, want := len(vs.VariationDimensions[1].Values), 13; got != want {
		t.Errorf("VariationDimensions[1].Values count = %d, want %d", got, want)
	}
}

// TestDecodeRealGetItemsSampleResponse pins the decode contract against an
// actual GetItems response captured from the Creators API. Notably the
// top-level container is `itemResults` (item singular, results plural) — not
// the `itemsResult` shape used by PA-API v5 / the third-party Python SDK.
func TestDecodeRealGetItemsSampleResponse(t *testing.T) {
	body := []byte(`{
  "errors": [
    {
      "code": "ItemNotAccessible",
      "message": "The ItemId B01180YUXS is not accessible through the Creators API."
    }
  ],
  "itemResults": {
    "items": [
      {
        "asin": "B0199980K4",
        "detailPageURL": "https://www.amazon.com/dp/B0199980K4?tag=xyz-20&linkCode=ogi&language=en_US&th=1&psc=1",
        "images": {
          "primary": {
            "small": {
              "height": 75,
              "url": "https://m.media-amazon.com/images/I/61s4tTAizUL._SL75_.jpg",
              "width": 56
            }
          }
        },
        "itemInfo": {
          "title": {
            "displayValue": "Genghis: The Legend of the Ten",
            "label": "Title",
            "locale": "en_US"
          }
        },
        "parentASIN": "B07QGKM68X"
      },
      {
        "asin": "B00BKQTA4A",
        "detailPageURL": "https://www.amazon.com/dp/B00BKQTA4A?tag=xyz-20&linkCode=ogi&language=en_US&th=1&psc=1",
        "images": {
          "primary": {
            "small": {
              "height": 75,
              "url": "https://m.media-amazon.com/images/I/41OiLOcQVJL._SL75_.jpg",
              "width": 46
            }
          }
        },
        "itemInfo": {
          "features": {
            "displayValues": [
              "Round watch featuring logoed white dial with stick indices",
              "36 mm stainless steel case with mineral dial window",
              "Quartz movement with analog display",
              "Leather calfskin band with buckle closure",
              "Water resistant to 30 m (99 ft): In general, withstands splashes or brief immersion in water, but not suitable for swimming"
            ],
            "label": "Features",
            "locale": "en_US"
          },
          "title": {
            "displayValue": "Daniel Wellington Women's 0608DW Sheffield Stainless Steel Watch",
            "label": "Title",
            "locale": "en_US"
          }
        },
        "parentASIN": "B07L5N7P32"
      }
    ]
  }
}`)
	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	// Errors block populates and preserves order/content.
	if got, want := len(resp.Errors), 1; got != want {
		t.Fatalf("len(Errors) = %d, want %d", got, want)
	}
	if got, want := resp.Errors[0].Code, "ItemNotAccessible"; got != want {
		t.Errorf("Errors[0].Code = %q, want %q", got, want)
	}
	// itemResults must decode (regression: previously absent).
	if resp.ItemsResult == nil {
		t.Fatal("ItemsResult is nil; itemResults JSON tag failed to decode")
	}
	if got, want := len(resp.ItemsResult.Items), 2; got != want {
		t.Fatalf("len(Items) = %d, want %d", got, want)
	}
	first := resp.ItemsResult.Items[0]
	if got, want := first.ASIN, "B0199980K4"; got != want {
		t.Errorf("Items[0].ASIN = %q, want %q", got, want)
	}
	if got, want := first.ParentASIN, "B07QGKM68X"; got != want {
		t.Errorf("Items[0].ParentASIN = %q, want %q", got, want)
	}
	if got, want := first.DetailPageURL, "https://www.amazon.com/dp/B0199980K4?tag=xyz-20&linkCode=ogi&language=en_US&th=1&psc=1"; got != want {
		t.Errorf("Items[0].DetailPageURL = %q, want %q", got, want)
	}
	if first.Images == nil || first.Images.Primary == nil || first.Images.Primary.Small == nil {
		t.Fatal("Items[0].Images.Primary.Small is nil")
	}
	if got, want := first.Images.Primary.Small.URL, "https://m.media-amazon.com/images/I/61s4tTAizUL._SL75_.jpg"; got != want {
		t.Errorf("Items[0] small image URL = %q, want %q", got, want)
	}
	if first.ItemInfo == nil || first.ItemInfo.Title == nil {
		t.Fatal("Items[0].ItemInfo.Title is nil")
	}
	if got, want := first.ItemInfo.Title.DisplayValue, "Genghis: The Legend of the Ten"; got != want {
		t.Errorf("Items[0].ItemInfo.Title.DisplayValue = %q, want %q", got, want)
	}
	// Second item exercises the features block (IdInfo).
	second := resp.ItemsResult.Items[1]
	if second.ItemInfo == nil || second.ItemInfo.Features == nil {
		t.Fatal("Items[1].ItemInfo.Features is nil")
	}
	if got, want := len(second.ItemInfo.Features.DisplayValues), 5; got != want {
		t.Errorf("Items[1].ItemInfo.Features.DisplayValues count = %d, want %d", got, want)
	}
	if got, want := second.ItemInfo.Features.Label, "Features"; got != want {
		t.Errorf("Items[1].ItemInfo.Features.Label = %q, want %q", got, want)
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
