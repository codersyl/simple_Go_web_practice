package main

import (
	"fmt"
	"log"
	"net/http"
)

type universalHandler struct{}

func (uniHandler *universalHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	entity := new(universalHandler)
	log.Fatal(http.ListenAndServe(":9999", entity))
}
