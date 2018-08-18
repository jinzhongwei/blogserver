package global

import (
	"runtime/debug"

	logger "github.com/shengkehua/xlog4go"
)

const CONF_FILE = "./conf/service.conf"

const (
	//最大打印请求的长度
	MaxReqLen = 500
	//最大应答打印长度
	MaxRespLen = 500
	//最大锁定时长
	MaxLockTTL = 60
)

var ServerIsRunning int32 = 1

var ServerQuit chan int = make(chan int, 3)

func SysPanicRecover(owner string) {
	if err := recover(); err != nil {
		logger.Error("owner=%s recover=[%v] stack=[%s]", owner, err, string(debug.Stack()))
	}
}
