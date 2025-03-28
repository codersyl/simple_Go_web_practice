// 用于解析「查询字符串」的，即解析 URL 中 ? 后面的内容
// 例：https://www.example.com/page?name=John&age=30
// 传入 Query()的参数就是 name=John&age=30
// 解析后得到 name 是 John， age 是 30
// 存在 url.Values 中

package url // import "net/url"

func (u *URL) Query() Values
    Query parses RawQuery and returns the corresponding values. It silently
    discards malformed value pairs. To check errors use ParseQuery.