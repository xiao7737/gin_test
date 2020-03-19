package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取get的参数 和 设置默认参数
func main() {
	r := gin.Default()
	r.GET("/get_param", func(c *gin.Context) {
		name := c.Query("name") // POST METHOD : c.PostForm("name")
		age := c.DefaultQuery("age", "default_age")
		c.String(http.StatusOK, "%s, %s", name, age)
	})
	r.Run()
}
