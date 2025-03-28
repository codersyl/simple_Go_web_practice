// Values是很多键值对
// Get可提取出它某个键对应的值
// 本项目目前把它用于提取 「查询字符串解析后得到的 Values，即 HTTP 请求中提交的一些表单数据

package url // import "net/url"

func (v Values) Get(key string) string
    // Get gets the first value associated with the given key. If there are no
    // values associated with the key, Get returns the empty string. To access
    // multiple values, use the map directly.