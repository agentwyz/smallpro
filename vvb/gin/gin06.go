package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	
	//http://localhost:9090/小王子/18
	//这个可以作为一个比较好的博客平台
	router.GET("/user/:name/:age", func (context *gin.Context)  {
		//获取路径参数
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})


	router.GET("/blog/:year/:month", func (context *gin.Context)  {
		year := context.Param("year")
		month := context.Param("month")

		context.JSON(http.StatusOK, gin.H{
			"year": year,
			"month": month,
		})
	})
	router.Run(":9090")
}