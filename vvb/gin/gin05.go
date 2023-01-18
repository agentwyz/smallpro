package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./login.html", "./page.html")

	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})

	//登陆使用POST进行提交
	router.POST("/login", func(context *gin.Context) {
		//得到数据
		//设置默认值
		//context.DefaultPostForm("")
		
		//username := context.PostForm("username")
		username, ok := context.GetPostForm("username")
		if !ok {
			username = "sb"
		}


		password := context.PostForm("pssword")

		
		context.HTML(http.StatusOK, "page.html", gin.H{
			"username": username,
			"password": password,
		})

	})

	router.Run(":9090")
}
