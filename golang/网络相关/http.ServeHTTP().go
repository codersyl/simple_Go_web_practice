// 目前直接出了第一个
// 接口 Handler 所需的唯一方法
// 需要实现的功能室根据 Request 作出对应的响应

package http // import "net/http"

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
// ServeHTTP calls f(w, r).

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
// ServeHTTP dispatches the request to the handler whose pattern most closely
// matches the request URL.
