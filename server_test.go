package paapi5

import "testing"

func TestServer(t *testing.T) {
	testCases := []struct {
		sv              *Server
		marketplace     string
		hostName        string
		region          string
		accept          string
		acceptLanguage  string
		contentType     string
		hmacAlgorithm   string
		serviceName     string
		aws4Request     string
		contentEncoding string
		url             string
	}{
		{sv: (*Server)(nil), marketplace: "www.amazon.com", hostName: "webservices.amazon.com", region: "us-east-1", accept: defaultAccept, acceptLanguage: "en_US", contentType: defaultContentType, hmacAlgorithm: defaultHMACAlgorithm, serviceName: defaultServiceName, contentEncoding: defaultContentEncoding, aws4Request: defaultAWS4Request, url: "https://webservices.amazon.com/paapi5/getitems"},
		{sv: New(WithMarketplace(LocaleJapan)), marketplace: "www.amazon.co.jp", hostName: "webservices.amazon.co.jp", region: "us-west-2", accept: defaultAccept, acceptLanguage: "ja_JP", contentType: defaultContentType, hmacAlgorithm: defaultHMACAlgorithm, serviceName: defaultServiceName, contentEncoding: defaultContentEncoding, aws4Request: defaultAWS4Request, url: "https://webservices.amazon.co.jp/paapi5/getitems"},
	}
	for _, tc := range testCases {
		if tc.sv.Marketplace() != tc.marketplace {
			t.Errorf("Server.Marketplace() is \"%v\", want \"%v\"", tc.sv.Marketplace(), tc.marketplace)
		}
		if tc.sv.HostName() != tc.hostName {
			t.Errorf("Server.HostName() is \"%v\", want \"%v\"", tc.sv.HostName(), tc.hostName)
		}
		if tc.sv.Region() != tc.region {
			t.Errorf("Server.Region() is \"%v\", want \"%v\"", tc.sv.Region(), tc.region)
		}
		if tc.sv.Accept() != tc.accept {
			t.Errorf("Server.Accept() is \"%v\", want \"%v\"", tc.sv.Accept(), tc.accept)
		}
		if tc.sv.AcceptLanguage() != tc.acceptLanguage {
			t.Errorf("Server.AcceptLanguage() is \"%v\", want \"%v\"", tc.sv.AcceptLanguage(), tc.acceptLanguage)
		}
		if tc.sv.ContentType() != tc.contentType {
			t.Errorf("Server.ContentType() is \"%v\", want \"%v\"", tc.sv.ContentType(), tc.contentType)
		}
		if tc.sv.HMACAlgorithm() != tc.hmacAlgorithm {
			t.Errorf("Server.HMACAlgorithm() is \"%v\", want \"%v\"", tc.sv.HMACAlgorithm(), tc.hmacAlgorithm)
		}
		if tc.sv.ServiceName() != tc.serviceName {
			t.Errorf("Server.ServiceName() is \"%v\", want \"%v\"", tc.sv.ServiceName(), tc.serviceName)
		}
		if tc.sv.AWS4Request() != tc.aws4Request {
			t.Errorf("Server.AWS4Request() is \"%v\", want \"%v\"", tc.sv.AWS4Request(), tc.aws4Request)
		}
		if tc.sv.ContentEncoding() != tc.contentEncoding {
			t.Errorf("Server.ContentEncoding() is \"%v\", want \"%v\"", tc.sv.ContentEncoding(), tc.contentEncoding)
		}
		url := tc.sv.URL(GetItems.Path()).String()
		if url != tc.url {
			t.Errorf("Server.URL() is \"%v\", want \"%v\"", url, tc.url)
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
