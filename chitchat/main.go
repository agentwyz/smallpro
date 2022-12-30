/*
package main

import (
	"net/http"
)

func main() {
	//多路复用器会对url请求进行检查, 并将它
	mux := http.NewServeMux()				//首先创建一个多路复用器
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	files := []string {
		""
	}
}
*/



