package http // import "net/http"


// http.Header中存储了一些键值对。
// 代表了HTTP头部中的client-server想传达的一些额外信息
// 比如告诉server自己想要的数据类型
type Header map[string][]string
// A Header represents the key-value pairs in an HTTP header.
// The keys should be in canonical form, as returned by CanonicalHeaderKey.