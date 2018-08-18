package logic

import (
	"context"
	"global"
)

//请求处理函数
type HandlerFunc func(ctx context.Context, request IRequest) IResponse

//待定
type IRequest interface {
	Create() IRequest
	Check(ctx context.Context) (errNo int)
}

type BaseRequest struct {
}

func (r *BaseRequest) Create() IRequest {
	return &BaseRequest{}
}

func (g *BaseRequest) Check(ctx context.Context) (errNo int) {
	return global.ERR_OK
}

type IResponse interface {
	GetError() (errNo int, errMsg string)
	SetError(errNo int)
}

//IResponse的base实现。所有API的response继承该struct
type BaseResponse struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

func NewBaseResponse() *BaseResponse {
	resp := &BaseResponse{}
	resp.SetError(global.ERR_OK)
	return resp
}

func (p *BaseResponse) GetError() (int, string) {
	return p.ErrNo, p.ErrMsg
}

func (p *BaseResponse) SetError(errNo int) {
	p.ErrNo = errNo
	p.ErrMsg = global.GetErrMsg(errNo)
	return
}

func (p *BaseResponse) SetErrorPlusMsg(errNo int, errstr string) {
	p.ErrNo = errNo
	p.ErrMsg = global.GetErrMsg(errNo) + " : " + errstr
	return
}

type RequestMeta struct {
	TraceId *string `json:"trace_id"`
}

func NewRequestMeta() *RequestMeta {
	return &RequestMeta{}
}

func NewEmptyRequestMeta() *RequestMeta {
	traceId := ""
	return &RequestMeta{&traceId}
}

func (p *RequestMeta) GetTraceId() string {
	return *p.TraceId
}

//websocket请求接口
type IWsRequest interface {
	Create() IWsRequest
	Check(ctx context.Context) (errNo int)
	Update(content []byte) (err error)
}

//data内容为前端传入的json串。其他websocket request要包含BaseRequest
//Create函数创建完请求对象后，Update函数用Data里的内容更新请求对象
type BaseWsRequest struct {
	Data interface{} `json:"data"`
}

func (r *BaseWsRequest) Create() IWsRequest {
	return &BaseWsRequest{}
}

func (r *BaseWsRequest) Check(ctx context.Context) (errNo int) {
	return global.ERR_OK
}

func (r *BaseWsRequest) Update(content []byte) (err error) {
	return
}
