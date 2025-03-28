// 把 / 路径与indexHandler这个函数绑在一起了，用于静态路由
http.HandleFunc("/", indexHandler)

// 参数2是一个handler
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}