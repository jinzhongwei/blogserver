package webserver

import (
	"global"
	"net/http"
	"strconv"
	"time"

	logger "github.com/shengkehua/xlog4go"
)

const BLOG_HTTP_SERVER = "BlogHttpServer"

type HttpServer struct {
	port int
}

func NewHttpServer(p int) *HttpServer {
	return &HttpServer{
		port: p,
	}
}

func (s *HttpServer) Run() error {
	defer global.SysPanicRecover(BLOG_HTTP_SERVER)
	addr := ":" + strconv.Itoa(s.port)
	mux := http.NewServeMux()
	th := http.FileServer(http.Dir("./template"))
	mux.Handle("/", th)
	for k, v := range urlHandlerManager {
		mux.Handle(k, v)
	}

	logger.Info("httpServer.Run||addr=%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		//端口占用导致监听失败，sleep 2秒重试一次
		time.Sleep(time.Second * 2)
		if err := http.ListenAndServe(addr, mux); err != nil {
			logger.Error("funcRetErr=http.ListenAndServe||err=%s", err.Error())
			return err
		}
	}

	return nil
}

func RunHttpServer(port int) error {
	// 设置静态目录
	s := NewHttpServer(port)
	return s.Run()
}
