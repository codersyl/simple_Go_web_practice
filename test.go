package main

import (
	"fmt"
)
func main() {
	type A = int;
	type B = int;
	arr := []A{}
	c := B(10)
	arr = append(arr, c)
	fmt.Println(arr)
}