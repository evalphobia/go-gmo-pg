go-gmo-pg
----

[![GoDoc][1]][2] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/evalphobia/go-gmo-pg?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-gmo-pg
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-gmo-pg.svg
[6]: https://github.com/evalphobia/go-gmo-pg/releases/latest
[7]: https://travis-ci.org/evalphobia/go-gmo-pg.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-gmo-pg
[9]: https://coveralls.io/repos/evalphobia/go-gmo-pg/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-gmo-pg?branch=master
[11]: https://codecov.io/github/evalphobia/go-gmo-pg/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-gmo-pg?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-gmo-pg
[14]: https://goreportcard.com/report/github.com/evalphobia/go-gmo-pg
[15]: https://img.shields.io/github/downloads/evalphobia/go-gmo-pg/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-gmo-pg/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-gmo-pg.svg
[18]: https://github.com/evalphobia/go-gmo-pg/stargazers
[19]: https://codeclimate.com/github/evalphobia/go-gmo-pg/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/go-gmo-pg
[21]: https://bettercodehub.com/edge/badge/evalphobia/go-gmo-pg?branch=master
[22]: https://bettercodehub.com/

Unofficial golang library for GMO Payment Gateway.

## Current Supported Payment API List

- au
    - AuAcceptUserEnd.idPass
    - AuContinuanceCancel.idPass
    - AuContinuanceChargeCancel.idPass
    - EntryTranAuAccept.idPass
    - EntryTranAuContinuance.idPass
    - EntryTranAu.idPass
    - ExecTranAuAccept.idPass
    - ExecTranAuContinuance.idPass
    - ExecTranAu.idPass

## Quick Usage

### EntryTranAuAccept and ExecTranAuAccept

```go
import (
	"fmt"
	"os"

	"github.com/evalphobia/go-gmo-pg/config"
	"github.com/evalphobia/go-gmo-pg/client"
	"github.com/evalphobia/go-gmo-pg/client/au"
)

func getClient() client.Client {
	shopID := os.Getenv("GMO_PG_SHOP_ID")
	shopPass := os.Getenv("GMO_PG_SHOP_PASS")
	isProdMode := os.Getenv("GMO_PG_PRODUCTION_MODE") != ""
	useDebugOption := os.Getenv("GMO_PG_DEBUG_OPTION") != ""

	conf, err := config.New(shopID, shopPass)
	if err != nil {
		panic(err)
	}

	if isProdMode {
		conf.SetAsProduction()
	}

	cli := client.New(conf)
	cli.Option.Retry = true
	cli.Option.Debug = useDebugOption
	return cli
}

func main() {
	cli := getClient()

	entryAPI := au.EntryTranAu{
		OrderID:     "001_abcdefg",
	}
	entryResp, err := entryAPI.Do(cli)
	switch {
	case err != nil:
		panic(err)
	case !entryResp.IsSuccess():
		// error process...
		return
	}

	execAPI := au.ExecTranAuAccept{
		OrderID:        entryAPI.OrderID,
		AccessID:       entryResp.AccessID,
		AccessPass:     entryResp.AccessPass,
		Commodity:      "the product name",
		ServiceName:    "my service name",
		ServiceTel:     "+81-3-0000-0000",
		RetURL:         "https://example.com/return",
		PaymentTermSec: 300, // 5min
	}
	execResp, err := execAPI.DoWithSjis(cli)
	switch {
	case err != nil:
		panic(err)
	case !execResp.IsSuccess():
		// error process...
		return
	}

	respData := map[string]string{
		"redirect_url": execResp.StartURL,
		"order_id":     entryAPI.OrderID,
		"access_id":    execResp.AccessID,
		"token":        execResp.Token,
	}
	fmt.Printf("%+v", respData)
}
```

### EntryTranAu and ExecTranAu

```go
import (
	"fmt"
	"os"

	"github.com/evalphobia/go-gmo-pg/config"
	"github.com/evalphobia/go-gmo-pg/client"
	"github.com/evalphobia/go-gmo-pg/client/au"
)

func getClient() client.Client {
	shopID := os.Getenv("GMO_PG_SHOP_ID")
	shopPass := os.Getenv("GMO_PG_SHOP_PASS")
	isProdMode := os.Getenv("GMO_PG_PRODUCTION_MODE") != ""
	useDebugOption := os.Getenv("GMO_PG_DEBUG_OPTION") != ""

	conf, err := config.New(shopID, shopPass)
	if err != nil {
		panic(err)
	}

	if isProdMode {
		conf.SetAsProduction()
	}

	cli := client.New(conf)
	cli.Option.Retry = true
	cli.Option.Debug = useDebugOption
	return cli
}

func main() {
	cli := getClient()

	entryAPI := au.EntryTranAu{
		OrderID:     "001_abcdefg",
		Amount:      1200,
		JobCd:       client.StatusCapture,
		PaymentType: client.AuPaymentTypeAcceptCode,
	}
	entryResp, err := entryAPI.Do(cli)
	switch {
	case err != nil:
		panic(err)
	case !entryResp.IsSuccess():
		// error process...
		return
	}

	execAPI := au.ExecTranAu{
		OrderID:      entryAPI.OrderID,
		AccessID:     entryResp.AccessID,
		AccessPass:   entryResp.AccessPass,
		Commodity:    "the product name",
		ServiceName:  "my service name",
		ServiceTel:   "+81-3-0000-0000",
		AuAcceptCode: "T1000000000000",
		ClientField1: "debug code 1",
	}
	execResp, err := execAPI.DoWithSjis(cli)
	switch {
	case err != nil:
		panic(err)
	case execResp.HasErrorAuAcceptCodeNotFound(),
		!execResp.IsSuccess():
		// error process...
		return
	}

	fmt.Printf("success payment: %+v", execResp)
}
```
