package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// 传入 err
// 返回格式化的调用栈，用于debug
func trace(message string) string {
	var pcs [32]uintptr

	// 跳过前3个，即 runtime.Callers 本身，trace()，defer func
	n := runtime.Callers(3, pcs[:])
	


	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		// 根据程序计数器获得函数的信息
		fn := runtime.FuncForPC(pc)
		// 获取文件名、行号
		file, line := fn.FileLine(pc)
		// 格式化添加到str中
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// 一个捕获错误的 handler
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}