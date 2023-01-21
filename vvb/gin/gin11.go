package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//中间件必须是一个gin.HandlerFunc类型
func StatCost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//
		start := time.Now()
		ctx.Set("name", "小王子")
		ctx.Next()
		//ctx.Abort() 
		cost := time.Since(start)
		log.Println(cost)
	}
}
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//定义一个中间件
func m1(c *gin.Context) {
	fmt.Printf("%v\n", "m1 in ...")
}


func main() {
	router := gin.Default()
	
	router.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})

	router.GET("/indexx", m1, indexHandler)
	router.Run(":9090")
}