package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./index03.html")

	router.GET("/user", func (context *gin.Context)  {
		// username := context.Query("username")
		// passward := context.Query("password")
		
		// u := UserInfo{
		// 	username: username,
		// 	password: passward,
		// }
		
		var u UserInfo  //结构体初始化
		err := context.ShouldBind(&u)	//使用指针
		fmt.Printf("%v", u)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index03.html", nil)
	})

	router.POST("/form", func (context *gin.Context)  {
		var u UserInfo  //结构体初始化
		err := context.ShouldBind(&u)	//使用指针
		fmt.Printf("%v", u)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	router.POST("/json", func (context *gin.Context)  {
		var u UserInfo  //结构体初始化
		//根据请求方法来绑定参数
		err := context.ShouldBind(&u)	//使用指针
		fmt.Printf("%v", u)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})


	router.Run(":9090")
}