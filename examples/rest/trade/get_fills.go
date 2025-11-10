package main

import (
	"log"

	"github.com/y0net/go-okx/examples/rest"
	"github.com/y0net/go-okx/rest/api"
	"github.com/y0net/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetFillsParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := trade.NewGetFills(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetFillsResponse))
}
