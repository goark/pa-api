package paapi5_test

import (
	"context"
	"fmt"
	"net/http"

	paapi5 "github.com/goark/pa-api"
)

func ExampleServer() {
	sv := paapi5.New() //Create default server
	fmt.Println("Marketplace:", sv.Marketplace())
	fmt.Println("Region:", sv.Region())
	fmt.Println("AcceptLanguage:", sv.AcceptLanguage())
	fmt.Println("URL:", sv.URL(paapi5.GetItems.Path()))
	// Output:
	// Marketplace: www.amazon.com
	// Region: us-east-1
	// AcceptLanguage: en_US
	// URL: https://webservices.amazon.com/paapi5/getitems
}

func ExampleNew() {
	sv := paapi5.New(paapi5.WithMarketplace(paapi5.LocaleJapan)) //Create server in Japan region
	fmt.Println("Marketplace:", sv.Marketplace())
	fmt.Println("Region:", sv.Region())
	fmt.Println("AcceptLanguage:", sv.AcceptLanguage())
	fmt.Println("URL:", sv.URL(paapi5.GetItems.Path()))
	// Output:
	// Marketplace: www.amazon.co.jp
	// Region: us-west-2
	// AcceptLanguage: ja_JP
	// URL: https://webservices.amazon.co.jp/paapi5/getitems
}

func ExampleDefaultClient() {
	client := paapi5.DefaultClient("mytag-20", "AKIAIOSFODNN7EXAMPLE", "1234567890") //Create default client
	fmt.Println("Marketplace:", client.Marketplace())
	// Output:
	// Marketplace: www.amazon.com
}

func ExampleClient() {
	//Create client for Janan region
	client := paapi5.New(
		paapi5.WithMarketplace(paapi5.LocaleJapan),
	).CreateClient(
		"mytag-20",
		"AKIAIOSFODNN7EXAMPLE",
		"1234567890",
		paapi5.WithContext(context.Background()),
		paapi5.WithHttpClient(http.DefaultClient),
	)
	fmt.Println("Marketplace:", client.Marketplace())
	// Output:
	// Marketplace: www.amazon.co.jp
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
