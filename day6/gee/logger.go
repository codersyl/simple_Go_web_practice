package gee

import (
	"time"
	"log"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now() // Now returns the current local time.
		// Process request
		c.Next()
		// 计算并打印处理本次请求所花费时间
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}