package router

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strings"
)

type Router struct {
	// 控制器路由
	Routes     map[string]http.Handler
	StaticDir  string
	StaticPath string
	// 404 路由
	notFound http.Handler
}

func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]http.Handler),
	}
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, r.StaticPath) {
		file := r.StaticDir + req.URL.Path[len(r.StaticPath):]
		http.ServeFile(rw, req, file)
		return
	}
	// websocket 究竟是什么类型的方法呢?
	fmt.Println(req.Method + ":" + req.URL.Path)
	h := r.Routes[req.Method+":"+req.URL.Path]
	if h == nil {
		rw.Write([]byte("404 Not Found"))
	} else {
		h.ServeHTTP(rw, req)
	}
}

func (r *Router) RegisterFunc(method string, path string, handlerFunc http.HandlerFunc) {
	handler := http.HandlerFunc(handlerFunc)
	r.Routes[method+":"+path] = handler
}

func (r *Router) RegisterHandler(method string, path string, handler http.Handler) {
	r.Routes[method+":"+path] = handler
}

func (r *Router) RegisterWebsocket(method string, path string, handler func(*websocket.Conn)) {
	r.Routes[method+":"+path] = websocket.Handler(handler)
}
