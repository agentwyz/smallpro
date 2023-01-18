// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )



// func main() {
// 	router := gin.Default()
	
// 	router.GET("/web", func(context *gin.Context) {
// 		//获取浏览器那边的请求携带的query string参数
// 		//通过query获取请求中携带的querystring参数
// 		name := context.Query("name")
		
// 		//可以设置一个默认的查询结果
// 		age := context.DefaultQuery("age", "sha?")

// 		//返回值是一个bool和string类型的
// 		data, ok := context.GetQuery("data")
// 		if !ok {
// 			//取不到
// 			data = "somebody"
// 		}

// 		context.JSON(http.StatusOK, gin.H{
// 			"name": name,
// 			"age": age,
// 			"data": data,
// 		})
// 	})

// 	router.Run(":9090")
// }