// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"html/template"
// )

// //定义一个用户user
// type User struct {
// 	Name string
// 	Gender string
// 	Age int
// }

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	//解析
// 	tampl, err := template.ParseFiles("./hello.tmpl")
// 	if err != nil {
// 		fmt.Printf("failed parse tamplate%v", err)
// 		return
// 	}

// 	//创建一个结构体
// 	u1 := User{
// 		Name: "小王子",
// 		Gender: "男",	//如果小写这个会出现无法访问的情况
// 		Age: 0,
// 	}

// 	//创建一个map
// 	m1 := map[string] interface{} {
// 		"Name": "小王子",
// 		"Gender": "男",
// 		"Age": 18,
// 	}
	
// 	hobbylist := []string {
// 		"篮球",
// 		"足球",
// 		"双色球",
// 	}
// 	//一种很骚的操作
// 	err = tampl.Execute(w, map[string] interface{}{
// 		"u1": u1,
// 		"m1": m1,
// 		"hobby": hobbylist,
// 	})

// 	if err != nil {
// 		fmt.Printf("falied to execute %v", err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", sayHello)
// 	err := http.ListenAndServe(":9090", nil)

// 	if err != nil {
// 		fmt.Printf("http server failed %v", err)
// 	}
// }

