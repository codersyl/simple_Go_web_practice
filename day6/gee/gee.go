package gee

import (
	"net/http"
	"log"
	"strings"
	"path"
	"html/template"
)

type HandlerFunc func(*Context)

type UniServer struct {
	*RouterGroup			// 内嵌类型，相当于一个匿名的属性，使得 UniServer 能够继承 *RouterGroup 的方法
	router *router			// 之前的路由映射表封装到了 router中
	groups []*RouterGroup	// 服务器所有的分组结点登记在册

	// HTML渲染 import "html/template"
	htmlTemplates *template.Template // 
	funcMap       template.FuncMap   // 把 string 映射到函数上
}

func New() *UniServer {
	one := &UniServer{
		router: newRouter(),
	}
	one.RouterGroup = &RouterGroup{uni : one} // 绑定内嵌 RouterGroup
	one.groups = []*RouterGroup{one.RouterGroup} // 把内嵌 RouterGroup 与服务器互相绑定
	return one
}

// 设置 FuncMap
func (engine *UniServer) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}


func (uni *UniServer) Run(addr string) (err error) {
	return http.ListenAndServe(addr, uni) // 只要有个 ServeHTTP方法，都能充当参数2
}

func (uni *UniServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range uni.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	
	// 利用 request 和 writer 构造 Context
	c := newContext(w, req)
	c.handlers = middlewares
	c.masterServer = uni
	uni.router.handle(c)
}

func (engine *UniServer) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
	// func template.Must(t *Template, err error) *Template
	// 如果 err 不为空，则调用 panic(err)
	// 如果 err 为空，返回 t

	// func template.New(name string) *Template
	// 创建一个名字为name的模版，并返回它的指针

	// func (t *Template) Funcs(funcMap FuncMap) *Template
	// 必须在 「模版被解析」之前调用，否则 panic，把funcMap中的函数加到模版的函数映射里
	// 一旦模板被解析（使用 Parse、ParseFiles 或 ParseGlob），就不能再添加或修改函数映射

	// func ParseGlob(pattern string) (*Template, error)
	// 于创建一个新的 Template 实例，并从与给定 pattern 匹配的文件中解析模板定义
}





// RouterGroup part

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


// 将中间件or多个中间件 应用到某个路由分组（结点）
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// create static handler
// http.FileSystem: 提供 Open(name string) (File, error) 的方法的接口就ok
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath) // import "path"
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	// func FileServer(root FileSystem) Handler
	// 把某个系统路径作为处理网络请求访问的静态文件服务器根目录

	// func StripPrefix(prefix string, h Handler) Handler
	// 返回了一个新的具有过滤前缀的逻辑的 handler：
	// 如果路径符合前缀，则去掉前缀，然后运行老handler
	// 不符合，返回404

	return func(c *Context) {
		file := c.Param("filepaht")
		if _, err := fs.Open(file); err != nil { // 找不到文件 or 权限不够
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
// 给用户使用，把 参数1打头的网路请求路径映射到 以参数2为根路径的文件夹中
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filerpath")

	// 为这个 handler 注册路由
	group.GET(urlPattern, handler)
}