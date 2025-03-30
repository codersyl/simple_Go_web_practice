package gee

import (
	"net/http"
	"log"
)

type HandlerFunc func(*Context)

type UniServer struct {
	*RouterGroup			// 内嵌类型，相当于一个匿名的属性，使得 UniServer 能够继承 *RouterGroup 的方法
	router *router			// 之前的路由映射表封装到了 router中
	groups []*RouterGroup	// 服务器所有的分组结点登记在册
}

func New() *UniServer {
	one := &UniServer{
		router: newRouter(),
	}
	one.RouterGroup = &RouterGroup{uni : one} // 绑定内嵌 RouterGroup
	one.groups = []*RouterGroup{one.RouterGroup} // 把内嵌 RouterGroup 与服务器互相绑定
	return one
}


// func (uni *UniServer) addRoute(method string, pattern string, handler HandlerFunc) {
// 	uni.router.addRoute(method, pattern, handler)
// }

// func (uni *UniServer) GET(pattern string, handler HandlerFunc) {
// 	uni.addRoute("GET", pattern, handler)
// }

// func (uni *UniServer) POST(pattern string, handler HandlerFunc) {
// 	uni.addRoute("POST", pattern, handler)
// }

func (uni *UniServer) Run(addr string) (err error) {
	return http.ListenAndServe(addr, uni) // 只要有个 ServeHTTP方法，都能充当参数2
}

func (uni *UniServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 利用 request 和 writer 构造 Context
	c := newContext(w, req)
	uni.router.handle(c)
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	parent      *RouterGroup  // support nesting
	uni      *UniServer       // all groups share a Uniserver instance
}

// RouterGroup 的一些方法

// 添加分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	server := group.uni
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, // 前缀累加
		parent: group,	// 绑定父结点
		uni : server,	// 绑定 server
	}
	server.groups = append(server.groups, newGroup)
	return newGroup
}

// 路由注册，由于 Server 有一个内嵌的 RouterGroup，也可以直接调用，实际上顶替了server原先的addRouter
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern) // 服务器端打印注册的路由
	group.uni.router.addRoute(method, pattern, handler) // 跑回去用 server的注册路由
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}