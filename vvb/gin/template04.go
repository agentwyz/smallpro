// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"html/template"
// )

// func index(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}

// 	msg := "这是index页面"
// 	//err = t.Execute(w, msg)
// 	err = t.ExecuteTemplate(w, "index.tmpl", msg)
	
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("./base.tmpl")
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// 	msg := "这是base页面"
	
// 	err = t.ExecuteTemplate(w, "base.tmpl",msg)

// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// }



// func main() {
// 	http.HandleFunc("/index", index)
// 	http.HandleFunc("/home", home)
	
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		fmt.Println("http server start failed, err:%v", err)
// 		return
// 	}
// }