package httpserver

import (
	"context"
	"encoding/json"
	"global"
	"logic"
	"net/http"
	"runtime/debug"
	"webserver/common"

	logger "github.com/shengkehua/xlog4go"
)

func httpRecoverPanic(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		logger.Error("funcRetErr=recover||err=%v||stack=%s", err, string(debug.Stack()))
		w.WriteHeader(http.StatusInternalServerError)
		processError(ctx, w, r, global.ERR_UNKNOW)
	}
}

func processError(ctx context.Context, w http.ResponseWriter, r *http.Request, errNo int) {
	response := &logic.BaseResponse{}
	response.SetError(errNo)
	writeResponse(ctx, w, response)
	common.WriteLog(ctx, r, response)
	return
}

//向客户端发送处理结果
func writeResponse(ctx context.Context, w http.ResponseWriter, response logic.IResponse) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		eResponse := &logic.BaseResponse{}
		eResponse.SetError(global.ERR_UNKNOW)
		writeResponse(ctx, w, eResponse)
	}
	return
}
