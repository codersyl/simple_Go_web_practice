package main

import (
	// "fmt"
	"net/http"
	"gee"
)

func main() {
	uni := gee.New()

	uni.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello! day2 almost done</h1>")
	})

	
	// 带名字的问号
	uni.GET("/hello", func(c *gee.Context) {
		// expect /hello?name= 
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})



	// 解析登录账号密码，给它传回去
	uni.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})


	uni.Run(":9999")
}