package main

import (
	"log"

	"github.com/y0net/go-okx/examples/rest"
	"github.com/y0net/go-okx/rest/api"
	"github.com/y0net/go-okx/rest/api/trade"
)

func main() {
	param := &trade.PostClosePositionParam{
		InstId:  "OKB-USDT",
		MgnMode: api.MgnModeCross,
		Ccy:     "OKB",
	}
	req, resp := trade.NewPostClosePosition(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.PostClosePositionResponse))
}
