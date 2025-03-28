package http // import "net/http"

// 监听端口并服务
// 参数1:监听的端口
// 参数2: 使用的handler
func ListenAndServe(addr string, handler Handler) error
//   ListenAndServe listens on the TCP network address addr and then calls
//   Serve with handler to handle requests on incoming connections. Accepted
//   connections are configured to enable TCP keep-alives.
//   The handler is typically nil, in which case DefaultServeMux is used.
//   ListenAndServe always returns a non-nil error.

