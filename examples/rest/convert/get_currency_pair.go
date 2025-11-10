package main

import (
	"log"

	"github.com/y0net/go-okx/examples/rest"
	"github.com/y0net/go-okx/rest/api/convert"
)

func main() {
	param := &convert.GetCurrencyPairParam{
		FromCcy: "BTC",
		ToCcy:   "USDT",
	}
	req, resp := convert.NewGetCurrencyPair(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.GetCurrencyPairResponse))
}
