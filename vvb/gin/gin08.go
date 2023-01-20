package main

import (
	"net/http"
	"path"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./index04.html")

	router.GET("/index", func (context *gin.Context)  {
		context.HTML(http.StatusOK, "index04.html", nil)
	})

	router.POST("/upload", func (context *gin.Context)  {
		//从请求中读取文件
		//从请求中获取携带的参数一样的
		//context.MultipartForm
		file, err := context.FormFile("f1")
		//将读取的文件保存到本地
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			//设置保存目标
			//1. dst := fmt.Sprintf("./%s", file.Filename)
			dst := path.Join("./", file.Filename)
			context.SaveUploadedFile(file, dst)
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}

	})
	router.Run(":9090")
}