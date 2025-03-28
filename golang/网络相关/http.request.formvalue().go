// 是net/http库提供的查找 HTTP 请求中表单数据的便捷方法
// 通过提供字段名，可获得对应数据
// 		例如 参数填入 username，可获得请求中的 username填的是什么


package http // import "net/http"

func (r *Request) FormValue(key string) string
// FormValue returns the first value for the named component of the query.
// The precedence order:
//  1. application/x-www-form-urlencoded form body (POST, PUT, PATCH only)
//  2. query parameters (always)
//  3. multipart/form-data form body (always)

// FormValue calls Request.ParseMultipartForm and Request.ParseForm if
// necessary and ignores any errors returned by these functions. If key is not
// present, FormValue returns the empty string. To access multiple values of
// the same key, call ParseForm and then inspect [Request.Form] directly.