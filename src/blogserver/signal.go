package main

import (
	"global"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	logger "github.com/shengkehua/xlog4go"
)

func signalServe() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	sig := <-c

	global.ServerIsRunning = 0
	logger.Warn("signalServe_get_signal|| signal=%v||stack=%s", sig, string(debug.Stack()))
	logger.Warn("send_server_quit_signal")
	global.ServerQuit <- 0
}
