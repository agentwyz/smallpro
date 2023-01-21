package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.HEAD()

	//访问/index, GET请求
	router.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	
	router.POST("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	router.DELETE("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	router.PUT("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	router.Any("/user", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case "GET":
			ctx.JSON(http.StatusOK, gin.H{"method" : "GET"})
		case "POST":
			ctx.JSON(http.StatusOK, gin.H{"method" : "POST"})
		}
	})

	//设置404页面
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg" : "www.fnmain.cn",
		})
	})

	//设置视频的首页
	// router.GET("/vedio/index", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/index",
	// 	})
	// })

	router.GET("/shop/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "shop/index",
		})
	})

	//设置路由组
	videoGroup := router.Group("/video")
	{	videoGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"index": "index",
			})
		})

		videoGroup.GET("/xx", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"xxx": "xx",
			})
		})
	}

	//路由组是支持嵌套的



	router.Run(":9090")
}