package main

import (
	"gee"
	"net/http"
	"time"
	"fmt"
	"html/template"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func FormatAsTime(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsTime": FormatAsTime,
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static") // 把 网络请求路径前缀 /assets 映射到 服务器根目录的/static文件夹下

	stu1 := &student{Name: "Jolyne", Age: 19}
	stu2 := &student{Name: "Ashitaka", Age: 16}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/time", func(c *gee.Context) {
		c.HTML(http.StatusOK, "time_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "date_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})

	r.Run(":9999")
}