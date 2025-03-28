package main

import (
	"fmt"
	"net/http"

	"gee"
)

func main() {
	uni := gee.New()
	uni.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	}) // 把这个handler注册给 "/"

	uni.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	uni.GET("/sayhello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello!\n")
	}) // 我自己额外写的

	uni.Run(":9999")
}
