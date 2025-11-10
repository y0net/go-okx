package market

import "github.com/y0net/go-okx/rest/api"

func NewGetCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}

type GetCandlesParam struct {
	api.PagingParam
	InstId string `url:"instId"`
	Bar    string `url:"bar,omitempty"`
	Before string `url:"before,omitempty"`
	Limit  string `url:"limit,omitempty"`
}

type GetCandlesResponse struct {
	api.Response
	Code    string     `json:"code"`
	Message string     `json:"msg"`
	Data    [][]string `json:"data"`
}
