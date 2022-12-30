package main

import (
	"fmt"
	"net/http"
)

//hander这个名字通常用来表示在指定事件触发 负责对事件进行处理的回调函数
//第一个参数为ResponseWriter接口, 第二个参数则为指向Request结构的指针
//handler函数会从Request(请求)结构中提取相关信息, 然后创建一个HTTP响应
//最后再通过ResponseWriter接口将响应返回给客户端
func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello GOLANG!, %s", request.URL.Path[1:])//提取到的参数
}

func main() {
	//第二个参数为函数
	http.HandleFunc("/", handler)

	//server := &Server{Addr: addr, Handler: handler}
	//return server.ListenAndServe()
	http.ListenAndServe(":8080", nil)
}