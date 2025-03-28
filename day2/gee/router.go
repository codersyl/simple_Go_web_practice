package gee

import (
	"net/http"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc // 路由映射表
}

// router 的构造函数
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 注册路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}


// 根据请求所形成的 Context 来找是否有对应的 handler
// 有的话，用 handler 写 Response
// 没的话，返回 404
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}