package gee

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type H map[string]interface{}

type Context struct {
	Writer 		http.ResponseWriter
	Req			*http.Request

	// request info
	Method 		string
	Path		string
	Params 		map[string]string // 一些参数
	
	// response info 
	StatusCode	int

	// middlerware part
	handlers 	[]HandlerFunc	// 中间件列表
	index 		int

	// 附属的服务器，用于访问其中的HTML模版
	masterServer 		*UniServer
}

// 根据 ServeHTTP 的两个参数来构造 Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: req,
		Method: req.Method,
		Path: req.URL.Path,
		index: -1,
	}
}

// 报错
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers) // 停止 handlers的轮询
	c.JSON(code, H{"message": err})
}

// 查看参数 key 对应的 value
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 用于查询表单数据，比如 key 填 username，即可查询请求提交的用户名的值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 解析查询字符串
// 并从中提取出我们想要的 key 对应的 value
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key) // net/url
}


// 填写 Response 的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 给响应消息写入消息头的键值对
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}



/// 接下来是在 Response Body 写入一些类型的数据的方法

// 写入 string类型
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}


// JSON 类型
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err!= nil {
		// 如果出错，则打印错误，并把状态码设为500，表示服务器错误
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 写入数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// 写入 HTML 类型
func (c *Context) HTML(code int, name string, data interface{}) {
	// 把响应头部设为 html
	c.SetHeader("Content-Type", "text/html")
	// 把响应状态码设为 code
	c.Status(code)
	if err := c.masterServer.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		// 如果渲染失败，报错500
		c.Fail(500, err.Error())
	}
}

// 先执行下一个中间件
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}