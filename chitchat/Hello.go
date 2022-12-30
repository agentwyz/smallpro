package main

import (
	"fmt"
	"net/http"
)

//这个名字通常用来表示在指定事件触发
func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}