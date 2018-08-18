package common

import (
	"context"
	"encoding/json"
	"fmt"
	"global"
	"logic"
	"net"
	"net/http"
	"runtime/debug"
	"time"

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
	WriteLog(ctx, r, response)
	return
}

//向客户端发送处理结果
func writeResponse(ctx context.Context, w http.ResponseWriter, response logic.IResponse) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := encoder.Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		eResponse := &logic.BaseResponse{}
		eResponse.SetError(global.ERR_UNKNOW)
		writeResponse(ctx, w, eResponse)
	}
	return
}

//记录日志
func WriteLog(ctx context.Context, r *http.Request, response logic.IResponse) {
	//请求uri
	path := ctx.Value(global.FIELD_PATH).(string)

	//处理结果
	resultTag := "Fail"
	errNo, _ := response.GetError()
	if errNo == global.ERR_OK {
		resultTag = "Success"
	}

	//计算耗时
	startTime := ctx.Value(global.FIELD_START_TIME).(time.Time)

	//meta信息
	requestMeta := ctx.Value("meta").(*logic.RequestMeta)
	requestMetaByte, _ := json.Marshal(requestMeta)

	request := fmt.Sprintf("%s", r.URL)

	//response信息
	responseInfoJsonByte, _ := json.Marshal(response)

	//计算耗时
	timeSpend := time.Since(startTime)
	//write access log
	logger.Info("WriteLog||status=%s||uri=%s||proc_time=%d||meta=%s||request=%s||response=%s||remoteip=%s||"+
		"xforwardfor=%s||realip=%s",
		resultTag, path, timeSpend.Nanoseconds()/int64(time.Millisecond),
		string(requestMetaByte),
		request[0:retLen(len(request), global.MaxReqLen)], string(responseInfoJsonByte)[0:retLen(len(string(responseInfoJsonByte)), global.MaxRespLen)],
		ctx.Value(global.FIELD_REMOTE_IP).(string),
		ctx.Value(global.FIELD_FORWARD_IP).(string),
		ctx.Value(global.FIELD_REAL_IP).(string))

	return
}

//限制日志长度
func retLen(slen, maxLen int) int {
	if slen > maxLen {
		return maxLen
	}
	return slen
}

//初始化context
func InitContext(r *http.Request) (ctx context.Context, errNo int) {
	now := time.Now()
	ctx = context.Background()
	ctx = context.WithValue(ctx, global.FIELD_START_TIME, now)
	ctx = context.WithValue(ctx, global.FIELD_LOGID, now.UnixNano())
	realIp := r.Header.Get("X-REAL-IP")
	ctx = context.WithValue(ctx, global.FIELD_REAL_IP, realIp)
	forwardIp := r.Header.Get("X-FORWARDED-FOR")
	ctx = context.WithValue(ctx, global.FIELD_FORWARD_IP, forwardIp)
	remoteIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	ctx = context.WithValue(ctx, global.FIELD_REMOTE_IP, remoteIp)
	ctx = context.WithValue(ctx, global.FIELD_PATH, r.URL.Path)

	requestMeta := logic.NewRequestMeta()
	if requestMeta, errNo = getRequestMeta(r); errNo != global.ERR_OK {
		requestMeta = logic.NewEmptyRequestMeta() //默认设为空
		ctx = context.WithValue(ctx, "meta", requestMeta)
		return
	}

	ctx = context.WithValue(ctx, "meta", requestMeta)
	return
}

func getRequestMeta(r *http.Request) (requestMeta *logic.RequestMeta, errNo int) {
	requestMeta = &logic.RequestMeta{}

	var (
		err      error
		metaJson string
	)
	if r.Method == "GET" {
		metaJson = r.FormValue("meta")
	} else if r.Method == "POST" {
		metaJson = r.PostFormValue("meta")
	}
	if metaJson != "" {
		err = json.Unmarshal([]byte(metaJson), requestMeta)
	}
	if metaJson == "" || err != nil || requestMeta.TraceId == nil {
		return nil, global.ILLEGAL_REQUEST_META
	}

	errNo = global.ERR_OK
	return
}
