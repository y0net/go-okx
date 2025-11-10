package main

import (
	"log"

	"github.com/y0net/go-okx/examples/rest"
	"github.com/y0net/go-okx/rest/api/market"
)

func main() {
	param := &market.GetCandlesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetIndexCandles(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetCandlesResponse))
}
