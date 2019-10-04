# [pa-api] -- APIs for Amazon Product Advertising API v5 by Golang

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/pa-api.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/pa-api)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/pa-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/pa-api.svg)](https://github.com/spiegel-im-spiegel/pa-api/releases/latest)

## Example

```go
package main

import (
    "fmt"

    paapi5 "github.com/spiegel-im-spiegel/pa-api"
    "github.com/spiegel-im-spiegel/pa-api/entity"
    "github.com/spiegel-im-spiegel/pa-api/query"
)

func main() {
    client := paapi5.DefaultClient("mytag-20", "AKIAIOSFODNN7EXAMPLE", "1234567890")

    q := query.NewGetItems(client.Marketplace(), client.PartnerTag(), client.PartnerType())
    q.ASINs([]string{"B07TYKJQFK"}).EnableImages(true).EnableParentASIN(true)
    body, err := client.Request(q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    //io.Copy(os.Stdout, bytes.NewReader(body))
    res, err := entity.DecodeResponse(body)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    fmt.Println(res.String())
    return
}
```

[pa-api]: https://github.com/spiegel-im-spiegel/pa-api "spiegel-im-spiegel/pa-api: APIs for Amazon Product Advertising API v5 by Golang"
