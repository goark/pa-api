package paapi5

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		partnerTag      string
		accessKey       string
		secretKey       string
		marketplace     string
		partnerType     string
		date            TimeStamp
		contentEncoding string
		hostName        string
		xAmzDate        string
		xAmzTarget      string
		payload         []byte
		sigedText       string
		sig             string
		authorization   string
	}{
		{
			partnerTag:      "mytag-20",
			accessKey:       "AKIAIOSFODNN7EXAMPLE",
			secretKey:       "1234567890",
			marketplace:     defaultMarketplace.String(),
			partnerType:     defaultPartnerType,
			date:            NewTimeStamp(time.Date(2019, time.September, 30, 8, 31, 54, 0, time.UTC)),
			contentEncoding: "amz-1.0",
			hostName:        "webservices.amazon.com",
			xAmzDate:        "20190930T083154Z",
			xAmzTarget:      "com.amazon.paapi5.v1.ProductAdvertisingAPIv1.GetItems",
			payload:         []byte(`{"ItemIds": ["B07YCM5K55"],"Resources": ["Images.Primary.Small","Images.Primary.Medium","Images.Primary.Large","ItemInfo.ByLineInfo","ItemInfo.ContentInfo","ItemInfo.Classifications","ItemInfo.ExternalIds","ItemInfo.ProductInfo","ItemInfo.Title"],"PartnerTag": "mytag-20","PartnerType": "Associates","Marketplace": "www.amazon.com","Operation": "GetItems"}`),
			sigedText:       "AWS4-HMAC-SHA256\n20190930T083154Z\n20190930/us-east-1/ProductAdvertisingAPI/aws4_request\n00edab8e9f221dd80f01241b85f4526f204ebf49818d678f041e21404a44b8cb",
			sig:             "717e266b28f02523fc39894f565532bd53fe80b37a9ed1b631b77c50483f8c08",
			authorization:   "AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20190930/us-east-1/ProductAdvertisingAPI/aws4_request,SignedHeaders=content-encoding;host;x-amz-date;x-amz-target,Signature=717e266b28f02523fc39894f565532bd53fe80b37a9ed1b631b77c50483f8c08",
		},
	}
	for _, tc := range testCases {
		client := New().CreateClient(tc.partnerTag, tc.accessKey, tc.secretKey)
		if client.Marketplace() != tc.marketplace {
			t.Errorf("Client.Marketplace() is \"%v\", want \"%v\"", client.Marketplace(), tc.marketplace)
		}
		if client.PartnerTag() != tc.partnerTag {
			t.Errorf("Client.PartnerTag() is \"%v\", want \"%v\"", client.PartnerTag(), tc.partnerTag)
		}
		if client.PartnerType() != tc.partnerType {
			t.Errorf("Client.PartnerType() is \"%v\", want \"%v\"", client.PartnerType(), tc.partnerType)
		}
		hds := newHeaders(client.server, GetItems, tc.date)
		if hds.get("Content-Encoding") != tc.contentEncoding {
			t.Errorf("headers.get(\"Content-Encoding\") is \"%v\", want \"%v\"", hds.get("Content-Encoding"), tc.contentEncoding)
		}
		if hds.get("Host") != tc.hostName {
			t.Errorf("headers.get(\"Host\") is \"%v\", want \"%v\"", hds.get("Host"), tc.hostName)
		}
		if hds.get("X-Amz-Date") != tc.xAmzDate {
			t.Errorf("headers.get(\"X-Amz-Date\") is \"%v\", want \"%v\"", hds.get("X-Amz-Date"), tc.xAmzDate)
		}
		if hds.get("X-Amz-Target") != tc.xAmzTarget {
			t.Errorf("headers.get(\"X-Amz-Target\") is \"%v\", want \"%v\"", hds.get("X-Amz-Target"), tc.xAmzTarget)
		}
		str := client.signedString(hds, tc.payload)
		if str != tc.sigedText {
			t.Errorf("Client.signedString() is \"%v\", want \"%v\"", str, tc.sigedText)
		}
		sig := client.signiture(str, hds)
		if sig != tc.sig {
			t.Errorf("Client.signiture() is \"%v\", want \"%v\"", sig, tc.sig)
		}
		auth := client.authorization(sig, hds)
		if auth != tc.authorization {
			t.Errorf("Client.authorization() is \"%v\", want \"%v\"", auth, tc.authorization)
		}
	}
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
