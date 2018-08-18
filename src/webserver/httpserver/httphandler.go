package httpserver

import (
	"context"
	"encoding/json"
	"global"
	"logic"
	"net/http"

	"webserver/common"
)

type HttpRequestHandler struct {
	Name    string
	Request logic.IRequest
	Handle  logic.HandlerFunc
}

func NewHttpRequestHandler(name string, request logic.IRequest, handle logic.HandlerFunc) *HttpRequestHandler {
	return &HttpRequestHandler{
		Name:    name,
		Request: request,
		Handle:  handle,
	}
}

func (h *HttpRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var (
		ctx         context.Context
		errNo       int
		err         error
		requestJson string
	)

	//初始化context
	ctx, errNo = common.InitContext(r)
	if errNo != global.ERR_OK {
		processError(ctx, w, r, errNo)
		return
	}

	//处理panic
	defer httpRecoverPanic(ctx, w, r)

	//获取输入参数
	if r.Method == "GET" {
		requestJson = r.FormValue("request")
	} else if r.Method == "POST" {
		requestJson = r.PostFormValue("request")
	}

	request := h.Request.Create()
	if requestJson != "" {
		err = json.Unmarshal([]byte(requestJson), request)
	}
	if requestJson == "" || err != nil {
		processError(ctx, w, r, global.ERR_INVALID_PARAM)
		return
	}
	//检查参数
	if errNo = request.Check(ctx); errNo != global.ERR_OK {
		processError(ctx, w, r, errNo)
		return
	}

	//处理请求
	response := h.Handle(ctx, request)

	//返回及打日志
	writeResponse(ctx, w, response)
	common.WriteLog(ctx, r, response)

	return
}
