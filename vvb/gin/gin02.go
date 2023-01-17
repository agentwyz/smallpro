// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	route := gin.Default()
	
// 	route.Static("/font", "./statics/assets/")

// 	route.LoadHTMLFiles("./index.html")
	
// 	route.GET("/index" ,func (context *gin.Context)  {
// 		context.HTML(http.StatusOK, "index.html", nil)
// 	})

// 	route.Run(":9090")
// }