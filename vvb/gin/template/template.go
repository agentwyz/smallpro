package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模版
	tmpl, err := template.ParseFiles("./hello.tmpl")
	//渲染模版
	if err != nil {
		fmt.Printf("parse template filed, err: %v", err)
		return
	}
	//渲染模版, 第一个参数是你往哪个地方写, 第二个参数是传递的数据
	err = tmpl.Execute(w, "hello, wangyangzheng")
	if err != nil {
		fmt.Printf("Execute tamplate error %v", err)
		return
	}
}

func main() {
	//创建解析路径
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)
	
	if err != nil {
		fmt.Printf("HTTP server start error failed, err %v", err)
	}
}