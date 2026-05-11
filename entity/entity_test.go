package entity

import "testing"

func TestDecodeResponseItemResultsAndCreatorsFields(t *testing.T) {
	body := []byte(`{
  "itemsResult": {
    "items": [
      {
        "asin": "A1",
        "images": {
          "primary": {
            "small": {"url": "https://example/s.jpg", "height": 75, "width": 75},
            "hiRes": {"url": "https://example/hi.jpg", "height": 1000, "width": 1000}
          },
          "variants": [
            {
              "small": {"url": "https://example/vs.jpg", "height": 75, "width": 75},
              "hiRes": {"url": "https://example/vh.jpg", "height": 1000, "width": 1000}
            }
          ]
        },
        "browseNodeInfo": {
          "browseNodes": [
            {
              "id": "123",
              "displayName": "Books",
              "contextFreeName": "Books",
              "websiteSalesRank": {
                "id": "999",
                "displayName": "Books",
                "contextFreeName": "Books",
                "salesRank": 7
              }
            }
          ]
        },
        "offersV2": {
          "listings": [
            {
              "type": "New",
              "violatesMAP": true
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
	if resp.ItemsResult == nil {
		t.Fatal("ItemsResult is nil")
	}
	if got, want := len(resp.ItemsResult.Items), 1; got != want {
		t.Fatalf("len(ItemsResult.Items) = %d, want %d", got, want)
	}

	item := resp.ItemsResult.Items[0]
	if item.Images == nil || item.Images.Primary == nil || item.Images.Primary.HiRes == nil {
		t.Fatal("Images.Primary.HiRes is nil")
	}
	if got, want := item.Images.Primary.HiRes.URL, "https://example/hi.jpg"; got != want {
		t.Errorf("Images.Primary.HiRes.URL = %q, want %q", got, want)
	}
	if item.Images.Variants == nil || len(item.Images.Variants) != 1 || item.Images.Variants[0].HiRes == nil {
		t.Fatal("Images.Variants[0].HiRes is nil")
	}
	if got, want := item.Images.Variants[0].HiRes.URL, "https://example/vh.jpg"; got != want {
		t.Errorf("Images.Variants[0].HiRes.URL = %q, want %q", got, want)
	}

	if item.BrowseNodeInfo == nil || len(item.BrowseNodeInfo.BrowseNodes) != 1 || item.BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank == nil {
		t.Fatal("BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank is nil")
	}
	if got, want := item.BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank.Id, "999"; got != want {
		t.Errorf("WebsiteSalesRank.Id = %q, want %q", got, want)
	}

	if item.OffersV2 == nil || item.OffersV2.Listings == nil || len(*item.OffersV2.Listings) != 1 {
		t.Fatal("OffersV2.Listings is empty")
	}
	if !(*item.OffersV2.Listings)[0].ViolatesMAP {
		t.Errorf("ViolatesMAP = false, want true")
	}
}

func TestDecodeResponseVariationSummaryAndBrowseSalesRank(t *testing.T) {
	body := []byte(`{
  "variationsResult": {
    "variationSummary": {
      "pageCount": 2,
      "variationCount": 13,
      "price": {
        "highestPrice": {"amount": 30.87, "currency": "GBP", "displayAmount": "GBP 30.87"},
        "lowestPrice": {"amount": 17.03, "currency": "GBP", "displayAmount": "GBP 17.03"}
      },
      "variationDimensions": [
        {
          "displayName": "Color",
          "locale": "en_GB",
          "name": "color_name",
          "values": ["Red", "Blue"]
        }
      ]
    }
  },
  "browseNodesResult": {
    "browseNodes": [
      {
        "id": "283155",
        "displayName": "Books",
        "contextFreeName": "Books",
        "isRoot": true,
        "salesRank": 11
      }
    ]
  }
}`)

	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.VariationsResult == nil || resp.VariationsResult.VariationSummary == nil {
		t.Fatal("VariationsResult.VariationSummary is nil")
	}
	vs := resp.VariationsResult.VariationSummary
	if got, want := vs.PageCount, 2; got != want {
		t.Errorf("PageCount = %d, want %d", got, want)
	}
	if got, want := vs.VariationCount, 13; got != want {
		t.Errorf("VariationCount = %d, want %d", got, want)
	}
	if vs.Price == nil || vs.Price.HighestPrice == nil || vs.Price.LowestPrice == nil {
		t.Fatal("VariationSummary.Price is nil")
	}
	if got, want := vs.Price.HighestPrice.Currency, "GBP"; got != want {
		t.Errorf("HighestPrice.Currency = %q, want %q", got, want)
	}
	if got, want := len(vs.VariationDimensions), 1; got != want {
		t.Fatalf("len(VariationDimensions) = %d, want %d", got, want)
	}
	if got, want := vs.VariationDimensions[0].Locale, "en_GB"; got != want {
		t.Errorf("VariationDimensions[0].Locale = %q, want %q", got, want)
	}

	if resp.BrowseNodesResult == nil || len(resp.BrowseNodesResult.BrowseNodes) != 1 {
		t.Fatal("BrowseNodesResult.BrowseNodes is empty")
	}
	if resp.BrowseNodesResult.BrowseNodes[0].SalesRank == nil {
		t.Fatal("BrowseNodes[0].SalesRank is nil")
	}
	if got, want := *resp.BrowseNodesResult.BrowseNodes[0].SalesRank, 11; got != want {
		t.Errorf("BrowseNodes[0].SalesRank = %d, want %d", got, want)
	}
}
