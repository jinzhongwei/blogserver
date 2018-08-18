package webserver

import (
	"logic"
	"logic/test"
	"webserver/httpserver"
)

var urlHandlerManager = make(map[string]*httpserver.HttpRequestHandler)

func Register(path string, req logic.IRequest, handle logic.HandlerFunc) {
	urlHandlerManager[path] = httpserver.NewHttpRequestHandler(path, req, handle)
}

func RegisterHttpRequestHandler() {
	Register("/test", new(logic.BaseRequest), test.TestHandler)
}

func init() {
	RegisterHttpRequestHandler()
}
