package entity

import "testing"

func TestDecodeResponseItemResultsOffersV2AndBrowseNodeInfo(t *testing.T) {
	body := []byte(`{
  "itemsResult": {
    "items": [
      {
        "asin": "A1",
        "images": {
          "primary": {
            "small": {"url": "https://example/s.jpg", "height": 75, "width": 75},
            "large": {"url": "https://example/l.jpg", "height": 500, "width": 500}
          },
          "variants": [
            {
              "small": {"url": "https://example/vs.jpg", "height": 75, "width": 75},
              "large": {"url": "https://example/vl.jpg", "height": 500, "width": 500}
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
	if item.Images == nil || item.Images.Primary == nil || item.Images.Primary.Large == nil {
		t.Fatal("Images.Primary.Large is nil")
	}
	if got, want := item.Images.Primary.Large.URL, "https://example/l.jpg"; got != want {
		t.Errorf("Images.Primary.Large.URL = %q, want %q", got, want)
	}
	if item.Images.Variants == nil || len(item.Images.Variants) != 1 || item.Images.Variants[0].Large == nil {
		t.Fatal("Images.Variants[0].Large is nil")
	}
	if got, want := item.Images.Variants[0].Large.URL, "https://example/vl.jpg"; got != want {
		t.Errorf("Images.Variants[0].Large.URL = %q, want %q", got, want)
	}

	if item.BrowseNodeInfo == nil || len(item.BrowseNodeInfo.BrowseNodes) != 1 || item.BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank == nil {
		t.Fatal("BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank is nil")
	}
	if got, want := item.BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank.SalesRank, 7; got != want {
		t.Errorf("WebsiteSalesRank.SalesRank = %d, want %d", got, want)
	}

	if item.OffersV2 == nil || item.OffersV2.Listings == nil || len(*item.OffersV2.Listings) != 1 {
		t.Fatal("OffersV2.Listings is empty")
	}
	if !(*item.OffersV2.Listings)[0].ViolatesMAP {
		t.Errorf("ViolatesMAP = false, want true")
	}
}

// Deprecated fields are kept on exported structs for source compatibility; they
// must still decode when present in arbitrary JSON.
func TestDeprecatedFieldsStillDecode(t *testing.T) {
	body := []byte(`{
  "itemsResult": {
    "items": [{
      "asin": "DEP1",
      "score": 0.91,
      "images": {
        "primary": {"hiRes": {"url": "https://example/hi.png", "height": 10, "width": 10}},
        "variants": [{"hiRes": {"url": "https://example/vhi.png", "height": 11, "width": 11}}]
      },
      "browseNodeInfo": {
        "browseNodes": [{
          "id": "bn1",
          "displayName": "Cat",
          "contextFreeName": "Cat",
          "salesRank": 99,
          "websiteSalesRank": {
            "id": "wsr-id",
            "displayName": "SiteCat",
            "contextFreeName": "SiteCat",
            "salesRank": 5
          }
        }]
      }
    }]
  },
  "browseNodesResult": {
    "browseNodes": [{
      "id": "283155",
      "displayName": "Books",
      "contextFreeName": "Books",
      "isRoot": true,
      "salesRank": 11
    }]
  }
}`)
	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	it := resp.ItemsResult.Items[0]
	if it.Score == nil || *it.Score != 0.91 {
		t.Fatalf("Score = %v, want 0.91", it.Score)
	}
	if it.Images.Primary.HiRes == nil || it.Images.Primary.HiRes.URL != "https://example/hi.png" {
		t.Fatal("deprecated HiRes primary not decoded")
	}
	if len(it.Images.Variants) != 1 || it.Images.Variants[0].HiRes == nil {
		t.Fatal("deprecated HiRes variant not decoded")
	}
	bn := it.BrowseNodeInfo.BrowseNodes[0]
	if bn.SalesRank == nil || *bn.SalesRank != 99 {
		t.Fatalf("browseNodes[].salesRank = %v, want 99", bn.SalesRank)
	}
	if bn.WebsiteSalesRank.Id != "wsr-id" || bn.WebsiteSalesRank.SalesRank != 5 {
		t.Fatalf("websiteSalesRank: %+v", bn.WebsiteSalesRank)
	}
	if len(resp.BrowseNodesResult.BrowseNodes) != 1 || resp.BrowseNodesResult.BrowseNodes[0].SalesRank == nil ||
		*resp.BrowseNodesResult.BrowseNodes[0].SalesRank != 11 {
		t.Fatalf("browseNodesResult browseNodes[].salesRank decode failed")
	}
}

func TestDecodeResponseVariationSummary(t *testing.T) {
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
	if got, want := vs.VariationDimensions[0].Name, "color_name"; got != want {
		t.Errorf("VariationDimensions[0].Name = %q, want %q", got, want)
	}
	if got, want := vs.VariationDimensions[0].Locale, "en_GB"; got != want {
		t.Errorf("VariationDimensions[0].Locale = %q, want %q", got, want)
	}
}
