package test

import (
	"context"
	"global"
	"logic"
	"rpc/mongodbwrapper"
)

type TestResponse struct {
	logic.BaseResponse
	Data string `json:"data"`
}

func NewTestResponse() *TestResponse {
	p := &TestResponse{}
	p.BaseResponse.SetError(global.ERR_OK)
	return p
}

func TestHandler(ctx context.Context, req logic.IRequest) logic.IResponse {
	mResp := NewTestResponse()
	mResp.Data = "test"
	mongodbwrapper.SearchAll()
	return mResp
}
