package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	//http重定向
	router := gin.Default()

	router.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.sogou.com/")
	})

	router.GET("/a", func (ctx *gin.Context)  {
		//把请求的地址URL修改
		//跳转到/b对应的路由处理函数
		ctx.Request.URL.Path = "/b"	//将请求的url进行修改
		router.HandleContext(ctx)	//继续后续的处理

		//ctx.Redirect(htttp.StatusMovedPermanently, "http://localhost:9090/b")
	})

	router.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "bbb",
		})
	})
	
	router.Run(":9090")
}