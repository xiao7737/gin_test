package main

import (
	"github.com/gin-gonic/gin"
)

// 设置route访问静态文件
func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	//同下
	//r.StaticFS("/static", http.Dir("static"))
	r.Run()
	//进入route_static后生成执行文件并运行：go build -o route_static && ./route_static
	//访问：localhost:8080/static/test.html
}
