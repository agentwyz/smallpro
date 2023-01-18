package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//前端框架直接接收json数据
func main() {
	router := gin.Default()
	
	//router.LoadHTMLFiles()

	router.GET("/json", func(context *gin.Context) {
		//方法1: 使用map
		data := map[string]interface{} {
			"name": "小王子",
			"message": "hello world",
			"age": 18,
		}

		//方法2: 使用gin.H{}-->map[string]any
		data1 := gin.H{
			"name:": "小王子",
			"message": "Hello world",
			"age": 18,
		}
 
		_ = data
		context.JSON(http.StatusOK, data1)
	})

	//方法三: 使用一个结构体
	//注意这个Name如果是大写, 则表示这个可以访问
	//但是可以通过打tag来进行表示访问
	type msg struct {
		Name string
		Age int  `json`
		Message string
	}

	router.GET("/anotherJSON", func (context *gin.Context)  {
		data := msg{
			Name: "小王子",
			Message: "Hello world",
			Age: 18,
		}
		context.JSON(http.StatusOK, data)
	})

	router.Run(":9090")
}