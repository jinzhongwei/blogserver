package main

import (
	"fmt"
	"global"
	"runtime"

	logger "github.com/shengkehua/xlog4go"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//初始化服务对象
	if err := initService(); err != nil {
		fmt.Println(fmt.Sprintf("msg=[initService fail] err=[%s]", err.Error()))
		logger.Fatal("msg=[initService_fail]||err=[%s]", err.Error())
		return
	}
	defer logger.Close()

	//初始化httpserver
	initHttpServer()
	defer global.SysPanicRecover("blogserver entry")

	//系统异常信号捕获
	go signalServe()

	q := <-global.ServerQuit
	if global.ServerIsRunning == 0 {
		logger.Error("server_start_failed||code:%d", q)
		return
	}
}
