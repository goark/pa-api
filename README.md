# [pa-api] -- Go client for the Amazon Creators API

[![check vulns](https://github.com/goark/pa-api/workflows/vulns/badge.svg)](https://github.com/goark/pa-api/actions)
[![lint status](https://github.com/goark/pa-api/workflows/lint/badge.svg)](https://github.com/goark/pa-api/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/pa-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/pa-api.svg)](https://github.com/goark/pa-api/releases/latest)

A Go client for [Amazon's Creators API][creatorsapi-docs], the OAuth2-based replacement for the now-retired Product Advertising API v5 (PA-API). The package keeps the import path `github.com/goark/pa-api` and the package name `paapi5` so existing call sites need only minor changes; under the hood the wire protocol, authentication and host all change.

This package requires Go 1.25 or later.

### Go version policy

- The source of truth for the minimum supported Go version is `go.mod`.
- README and CI are kept aligned with `go.mod`.
- If the `go` directive changes, update this README line in the same change.

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
- **Query constructor `marketplace` arguments are compatibility-only** (`NewGetItems`, `NewSearchItems`, `NewGetVariations`, `NewGetBrowseNodes`). Actual routing always uses the client's configured marketplace via the `x-marketplace` header, so set marketplace on `Server`/`Client` (`creatorsapi.WithMarketplace(...)`) rather than per-query.

### Ignored legacy request fields

The library keeps several legacy knobs for source compatibility, but the Creators API does not accept them in request bodies. They are currently treated as no-ops.

| Legacy field / option | Previous behavior (PA-API v5) | Current behavior (Creators API) |
|---|---|---|
| `Marketplace` request body field | Selected target marketplace in-body | Ignored in-body. Routing is done by `x-marketplace` header |
| `PartnerType` (`Associates`) | Explicit body parameter | Ignored. Partner type is implicit |
| `Merchant` | Offer filtering selector | Ignored |
| `OfferCount` | Offer summary limiter | Ignored |

### Marketplace routing precedence

Marketplace selection is evaluated in this order:

1. `Server`/`Client` marketplace (for example `creatorsapi.WithMarketplace(...)`)
2. Request header `x-marketplace` sent by the client

Query constructor marketplace arguments are retained only for compatibility and are not used to route requests.

### Credential Version map

Creators API credentials are region-group scoped. Use credentials that match the marketplace group you call.

New applications typically receive **v3.x** credentials (Login with Amazon token endpoints). **v2.x** credentials use regional Cognito endpoints and a different OAuth scope and catalog `Authorization` header shape; if Associates Central still shows `2.1`/`2.2`/`2.3`, pass `creatorsapi.WithCredentialVersion("2.1")` (etc.) on `CreateClient`.

| Credential Version | Region group | Token endpoint (default) | Catalog `Authorization` |
|---|---|---|---|
| `3.1` | North America | `https://api.amazon.com/auth/o2/token` | `Bearer <token>` |
| `3.2` | Europe / Middle East / India | `https://api.amazon.co.uk/auth/o2/token` | `Bearer <token>` |
| `3.3` | Far East | `https://api.amazon.co.jp/auth/o2/token` | `Bearer <token>` |
| `2.1` | North America | `https://creatorsapi.auth.us-east-1.amazoncognito.com/oauth2/token` | `Bearer <token>, Version 2.1` |
| `2.2` | Europe / Middle East / India | `https://creatorsapi.auth.eu-south-2.amazoncognito.com/oauth2/token` | `Bearer <token>, Version 2.2` |
| `2.3` | Far East | `https://creatorsapi.auth.us-west-2.amazoncognito.com/oauth2/token` | `Bearer <token>, Version 2.3` |

Example marketplace groups (same as before): NA — `www.amazon.com`, `www.amazon.ca`, `www.amazon.com.mx`, `www.amazon.com.br`; EU — `www.amazon.co.uk`, `www.amazon.de`, `www.amazon.fr`, `www.amazon.in`, `www.amazon.sa`, `www.amazon.ae`; FE — `www.amazon.co.jp`, `www.amazon.sg`, `www.amazon.com.au`.

If token acquisition fails with an auth error, verify that your credential version and marketplace group match.

### Rate limits and retry guidance

Amazon starts new credential sets with low throughput (commonly around 1 TPS; check your Associates Central quota for authoritative limits). Build callers with backpressure and retry control.

Recommended client behavior:

- Treat `429` and transient `5xx` as retryable.
- Use exponential backoff with jitter.
- Cap retry attempts and total request timeout.
- Avoid synchronized retries across workers.
- Consider per-credential rate limiting in your application.

### Getting Creators API credentials

Only the primary Amazon Associates account owner can mint credentials. From [Associates Central][creatorsapi-portal] go to **Tools** → **Creators API** → **Create Application**, then **Add New Credential**. Copy the Credential Secret immediately — it is only shown once. Note the **Credential Version** shown for your credential (`3.1`/`3.2`/`3.3` for Login with Amazon, or legacy `2.1`/`2.2`/`2.3` for Cognito). You'll need it if you call multiple regions from the same process or if your credential region group does not match the configured marketplace.

### Quick migration checklist

Use this path when migrating existing PA-API v5 call sites:

1. Keep the import path as `github.com/goark/pa-api` (alias as `creatorsapi` in examples if preferred).
2. Replace client credential inputs: AWS Access Key / Secret Key -> Creators API Credential ID / Credential Secret.
3. Configure marketplace on `Server`/`Client` (`WithMarketplace`) and do not rely on per-query marketplace arguments.
4. Replace V1 offers usage with OffersV2 (`EnableOffersV2`; `EnableOffers` remains as a compatibility alias).
5. Remove expectations around `Merchant`, `OfferCount`, and `PartnerType` request effects; these are ignored.
6. Confirm Credential Version matches the marketplace group you call (`3.1`/`3.2`/`3.3` by default per marketplace, or `2.1`/`2.2`/`2.3` for legacy Cognito credentials via `WithCredentialVersion`).
7. Add retry and rate-limit control for `429` and transient `5xx` responses.
8. Run local verification with your project's standard test/lint workflow before opening a PR.

## Usage

### Create a server configuration

```go
sv := creatorsapi.New() // default: US marketplace, NA credential version 3.1 (Login with Amazon)
fmt.Println("Marketplace:", sv.Marketplace())
fmt.Println("CredentialVersion:", sv.CredentialVersion())
fmt.Println("URL:", sv.URL(creatorsapi.GetItems.Path()))
// Output:
// Marketplace: www.amazon.com
// CredentialVersion: 3.1
// URL: https://creatorsapi.amazon/catalog/v1/getItems
```

For another marketplace:

```go
sv := creatorsapi.New(creatorsapi.WithMarketplace(creatorsapi.LocaleJapan)) // Japan -> credential version 3.3
fmt.Println("Marketplace:", sv.Marketplace())
fmt.Println("CredentialVersion:", sv.CredentialVersion())
// Output:
// Marketplace: www.amazon.co.jp
// CredentialVersion: 3.3
```

The credential version is auto-derived from the configured marketplace's region group (`3.1`/`3.2`/`3.3`). Override with `creatorsapi.WithCredentialVersion("2.2")` when using legacy Cognito credentials (`2.1`/`2.2`/`2.3`), or when your credential group differs from the marketplace default.

### Create a client

```go
client := creatorsapi.DefaultClient("mytag-20", "YOUR_CREDENTIAL_ID", "YOUR_CREDENTIAL_SECRET")
fmt.Println("Marketplace:", client.Marketplace())
// Output:
// Marketplace: www.amazon.com
```

For a different marketplace plus a custom `*http.Client`:

```go
client := creatorsapi.New(
    creatorsapi.WithMarketplace(creatorsapi.LocaleJapan),
).CreateClient(
    "mytag-20",
    "YOUR_CREDENTIAL_ID",
    "YOUR_CREDENTIAL_SECRET",
    creatorsapi.WithHttpClient(http.DefaultClient),
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

    creatorsapi "github.com/goark/pa-api"
    "github.com/goark/pa-api/entity"
    "github.com/goark/pa-api/query"
)

func main() {
    client := creatorsapi.New(
        creatorsapi.WithMarketplace(creatorsapi.LocaleJapan),
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
