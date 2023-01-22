package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

//定义一个中间件: 统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Printf("%v\n", "m1 in ...")

	//开始计时
	start := time.Now()

	c.Next()	//执行下一个函数
	//c.Abort()	//阻止后续的处理函数
	//如果不想使用后面的语句使用return语句
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(ctx *gin.Context) {
	fmt.Println("m2 in...")
	ctx.Set("name", "qimi")	//实现跨中间件去取值
	fmt.Println("m2 out ...")
}

//定义登陆中间件
// func register(c *gin.Context) {

// }

//使用闭包来做一个中间件
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其它准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登陆的判断
			//if 是登陆用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}

	}
}


func main() {
	router := gin.Default()
	//r := gin.New() 默认使用logger()和recovery()中间件
	


	//默认使用两个中间件
	router.Use(m1, m2, authMiddleware(true))

	router.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})
	
	//全局注册中间件函数
	router.Use(m1, authMiddleware(true))

	//可以使用多个中间件

	router.GET("/indexx", indexHandler)
	router.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	// router.GET("/user", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"msg": "user",
	// 	})
	// })

	// xxGroup := router.Group("/xx", authMiddleware(true))
	// {
	// 	xxGroup.GET("/index", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"msg": "xxgroup",
	// 		})
	// 	})
	// }
	
	// //路由注册中间件
	// xx2Group := router.Group("/xx2")
	// xx2Group.Use(authMiddleware(true))
	// {
	// 	xx2Group.GET("/index", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"msg": "xx2Group",
	// 		})
	// 	})
	// }
	router.Run(":9090")
}