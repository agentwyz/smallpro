package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//自定义一个函数kua
	//要么只有一个返回值, 要么需要两个返回值, 且第二个返回值必须是error类型
	//如果想在模版中使用这个函数, 一定要在解析模版之前进行
	kua := func(name string) (string, error){
		return name + "年轻又帅气", nil
	}

	//链式操作, 创建一个名字是f的模版对象
	//解析的名字需要相同
	t := template.New("f.tmpl")
	
	t.Funcs(template.FuncMap{
		"kua": kua,
	}) 
	
	_, err := t.ParseFiles("./f.tmpl")

	if err != nil {
		fmt.Printf("parse template failed err: %v\n", err)
	}
	name := "小王子"
	err = t.Execute(w, name)

	if err != nil {
		fmt.Printf("failed to excute %v\n", err)
	}
}

func f2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./u.tmpl", "./t.tmpl")

	if err != nil {
		fmt.Printf("parse template failed err %v\n", err)
	}

	name := "小王子"
	err = t.Execute(w, name)
	
	if err != nil {
		fmt.Printf("failed execute err %v\n", err)
	}
}


func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpl", f2)
		
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v", err)
		return
	}
}