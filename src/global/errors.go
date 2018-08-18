package global

//错误码

const (
	ERR_OK            = 0
	ERR_FAIL          = 2
	ERR_INVALID_PARAM = 100
	ERR_INTERNAL      = 101

	ILLEGAL_REQUEST_META = 200

	ERR_METHOD_NOT_ALLOWED = 404
	ERR_UNKNOW             = 500
)

//错误信息映射
var ErrNo2ErrMsg = map[int]string{
	ERR_OK:                 "OK",
	ERR_FAIL:               "Fail",
	ERR_INVALID_PARAM:      "参数错误",
	ERR_INTERNAL:           "内部错误，请稍候重试",
	ERR_METHOD_NOT_ALLOWED: "不支持该请求方式",
	ERR_UNKNOW:             "unknown error",
	ILLEGAL_REQUEST_META:   "illegal request meta",
}

func GetErrMsg(errNo int) string {
	if msg, ok := ErrNo2ErrMsg[errNo]; ok {
		return msg
	}
	return ""
}
