package main

import (
	"html/template"
	"net/http"
	"github.com/gin-gonic/gin"
)

//加载静态文件


func main() {
	//创建默认的路由
	router := gin.Default()
	
	//加载静态文件

	router.Static("/css", "./")

	//设置模版函数
	router.SetFuncMap(template.FuncMap{
		"safe": func (str string) template.HTML {
			return template.HTML(str)
		},
	})

	//模版解析
	router.LoadHTMLFiles("./index.tmpl")
	//表示正则表达式进行匹配: r.LoadHTMLGlob("temp/**/*")


	//HTTP请求
	router.GET("/index", func (c *gin.Context)  {
		//模版的渲染
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "<script>alert(\"女朋友\")</script>",
			"data": "<p>哈哈哈</p>",
		})
	})

	//启动
	router.Run(":9090")
}