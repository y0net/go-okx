package main

import (
	"log"

	"github.com/y0net/go-okx/examples/rest"
	"github.com/y0net/go-okx/rest/api/convert"
)

func main() {
	req, resp := convert.NewGetCurrencies()
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*convert.GetCurrenciesResponse))
}
