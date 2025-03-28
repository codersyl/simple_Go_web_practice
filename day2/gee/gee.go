package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type UniServer struct {
	router *router // 之前的路由映射表封装到了 router中
}

func New() *UniServer {
	return &UniServer{
		router: newRouter(),
	}
}


func (uni *UniServer) addRoute(method string, pattern string, handler HandlerFunc) {
	uni.router.addRoute(method, pattern, handler)
}

func (uni *UniServer) GET(pattern string, handler HandlerFunc) {
	uni.addRoute("GET", pattern, handler)
}

func (uni *UniServer) POST(pattern string, handler HandlerFunc) {
	uni.addRoute("POST", pattern, handler)
}

func (uni *UniServer) Run(addr string) (err error) {
	return http.ListenAndServe(addr, uni) // 只要有个 ServeHTTP方法，都能充当参数2
}

func (uni *UniServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 利用 request 和 writer 构造 Context
	c := newContext(w, req)

	// 此时 c 只完成了一半，那一半的数据 from 请求，需要根据请求进行响应

	// 利用 uni 自带的路由映射表查找响应方法并响应
	uni.router.handle(c)
}