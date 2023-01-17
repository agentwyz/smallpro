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

func xss(w http.ResponseWriter, r *http.Request) {
	//解析模版
	//解析模版之前定义一个自定义的函数safe
	t, err := template.New("o.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML{
			return template.HTML(s)
		},
	}).ParseFiles("./o.tmpl")

	if err != nil {
		fmt.Printf("%v", err)
	}
	//渲染模版
	//这里防止跨站脚本攻击
	str := "<script>alert(\"嘿嘿\")</script>"
	t.Execute(w, str)
}

func main() {
	http.HandleFunc("/index", xss)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println("HTTP server failed, err:%v", err)
		return
	}
}