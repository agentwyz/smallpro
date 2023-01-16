package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//具有一点lisp的语法
	t, err := template.New("o.tmpl").
		Delims("((", "))").
		ParseFiles("./o.tmpl")
	
	if err != nil {
		fmt.Printf("parse filed %v", err)
		return
	}
	//渲染模版
	name := "小王子"
	err = t.Execute(w, name)

	if err != nil {
		fmt.Printf("execute template failed, err:%v\n", err)
		return
	}
	
}


func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println("HTTP server failed, err:%v", err)
		return
	}

}