# [pa-api] -- Go client for the Amazon Creators API

[![check vulns](https://github.com/goark/pa-api/workflows/vulns/badge.svg)](https://github.com/goark/pa-api/actions)
[![lint status](https://github.com/goark/pa-api/workflows/lint/badge.svg)](https://github.com/goark/pa-api/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/pa-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/pa-api.svg)](https://github.com/goark/pa-api/releases/latest)

A Go client for [Amazon's Creators API][creatorsapi-docs], the OAuth2-based replacement for the now-retired Product Advertising API v5 (PA-API). The package keeps the import path `github.com/goark/pa-api` and the package name `paapi5` so existing call sites need only minor changes; under the hood the wire protocol, authentication and host all change.

This package requires Go 1.21 or later.

## Migrating from PA-API v5

Amazon retired PA-API v5 on 2026-05-15 and replaced it with the Creators API. The two big differences:

| | PA-API v5 (old) | Creators API (new) |
|---|---|---|
| Host | `webservices.amazon.<tld>` per locale | Single `creatorsapi.amazon` for every marketplace |
| Marketplace selection | `Marketplace` body field | `x-marketplace` request header |
| Auth | AWS SigV4 (Access Key + Secret Key) | OAuth2 client_credentials (Credential ID + Credential Secret + Version) |
| Body keys | `PascalCase` | `lowerCamelCase` |
| Resource enums | `Images.Primary.Medium` | `images.primary.medium` |
| `Offers` (V1) | Available | Removed; use `OffersV2` |

Code-level changes you should expect when upgrading:

- **`CreateClient` / `DefaultClient` arguments**: the second/third positional arguments are now `credentialID` and `credentialSecret` (Creators API credentials issued via Associates Central) instead of AWS access/secret keys.
- **`q.EnableOffers()`** is now an alias for `q.EnableOffersV2()` with a deprecation comment; the V1 Offers resource is gone.
- **`Server.Region()` is deprecated** and is no longer used by the client; it remains for back-compat callers that record it as metadata.
- **Response JSON keys** returned by the Creators API are lowerCamelCase. The existing Go field names in `entity.Response` decode case-insensitively from these new keys, but if your code re-serialises a `Response` value you'll see PascalCase output for some fields and lowerCamelCase for the explicitly tagged ones (notably `id`).
- **`SearchItems` filters** `Marketplace`, `PartnerType`, `Merchant`, and `OfferCount` are silently ignored — those fields are not accepted by the Creators API. Existing code using those filters compiles but the values are dropped.

### Getting Creators API credentials

Only the primary Amazon Associates account owner can mint credentials. From [Associates Central][creatorsapi-portal] go to **Tools** → **Creators API** → **Create Application**, then **Add New Credential**. Copy the Credential Secret immediately — it is only shown once. Note the **Credential Version** (`2.1` for North America, `2.2` for Europe, `2.3` for Far East) — you'll need it if you call multiple regions from the same process.

## Usage

### Create a server configuration

```go
sv := paapi5.New() // default: US marketplace, NA credential version 2.1
fmt.Println("Marketplace:", sv.Marketplace())
fmt.Println("CredentialVersion:", sv.CredentialVersion())
fmt.Println("URL:", sv.URL(paapi5.GetItems.Path()))
// Output:
// Marketplace: www.amazon.com
// CredentialVersion: 2.1
// URL: https://creatorsapi.amazon/catalog/v1/getItems
```

For another marketplace:

```go
sv := paapi5.New(paapi5.WithMarketplace(paapi5.LocaleJapan)) // Japan -> credential version 2.3
fmt.Println("Marketplace:", sv.Marketplace())
fmt.Println("CredentialVersion:", sv.CredentialVersion())
// Output:
// Marketplace: www.amazon.co.jp
// CredentialVersion: 2.3
```

The credential version is auto-derived from the configured marketplace's region group. Override it explicitly with `paapi5.WithCredentialVersion("2.2")` when needed.

### Create a client

```go
client := paapi5.DefaultClient("mytag-20", "YOUR_CREDENTIAL_ID", "YOUR_CREDENTIAL_SECRET")
fmt.Println("Marketplace:", client.Marketplace())
// Output:
// Marketplace: www.amazon.com
```

For a different marketplace plus a custom `*http.Client`:

```go
client := paapi5.New(
    paapi5.WithMarketplace(paapi5.LocaleJapan),
).CreateClient(
    "mytag-20",
    "YOUR_CREDENTIAL_ID",
    "YOUR_CREDENTIAL_SECRET",
    paapi5.WithHttpClient(http.DefaultClient),
)
```

The client transparently obtains and caches an OAuth2 access token from the appropriate Cognito endpoint (`expires_in` minus a 30-second leeway) and forwards it to the API as `Authorization: Bearer <token>, Version <2.x>`.

## Sample code

### GetItems

```go
package main

import (
    "context"
    "fmt"

    paapi5 "github.com/goark/pa-api"
    "github.com/goark/pa-api/entity"
    "github.com/goark/pa-api/query"
)

func main() {
    client := paapi5.New(
        paapi5.WithMarketplace(paapi5.LocaleJapan),
    ).CreateClient(
        "mytag-20",
        "YOUR_CREDENTIAL_ID",
        "YOUR_CREDENTIAL_SECRET",
    )

    q := query.NewGetItems(
        client.Marketplace(),
        client.PartnerTag(),
        client.PartnerType(),
    ).ASINs([]string{"B07YCM5K55"}).EnableImages().EnableItemInfo().EnableParentASIN()

    body, err := client.RequestContext(context.Background(), q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }

    res, err := entity.DecodeResponse(body)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    fmt.Println(res.String())
}
```

### GetVariations

```go
q := query.NewGetVariations(
    client.Marketplace(),
    client.PartnerTag(),
    client.PartnerType(),
).ASIN("B07YCM5K55").EnableImages().EnableItemInfo().EnableParentASIN()

body, err := client.RequestContext(context.Background(), q)
```

### SearchItems

```go
q := query.NewSearchItems(
    client.Marketplace(),
    client.PartnerTag(),
    client.PartnerType(),
).Search(query.Keywords, "数学ガール").EnableImages().EnableItemInfo().EnableParentASIN()

body, err := client.RequestContext(context.Background(), q)
```

### GetBrowseNodes

```go
q := query.NewGetBrowseNodes(
    client.Marketplace(),
    client.PartnerTag(),
    client.PartnerType(),
).BrowseNodeIds([]string{"3040", "3045"}).EnableBrowseNodes()

body, err := client.RequestContext(context.Background(), q)
```

## Contributors

Many thanks for [contributors](https://github.com/goark/pa-api/graphs/contributors "Contributors to goark/pa-api")

## Links

- [Amazon Creators API documentation][creatorsapi-docs]
- [Associates Central — Creators API portal][creatorsapi-portal]

[pa-api]: https://github.com/goark/pa-api "goark/pa-api: Go client for the Amazon Creators API"
[creatorsapi-docs]: https://affiliate-program.amazon.com/creatorsapi/docs/en-us/introduction
[creatorsapi-portal]: https://affiliate-program.amazon.com/creatorsapi
