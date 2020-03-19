package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(IPMiddleware())
	r.GET("/test_middleware", func(c *gin.Context) {
		c.String(http.StatusOK, "check OK")
	})
	r.Run()
	//127.0.0.1:8080/test_middleware
}

// Default() 默认会加载 logger 和 recovery 两个中间件
// 自定义中间件就可以参照logger进行设置
func IPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// white ip list
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP() //当前访问ip
		for _, ip := range ipList {
			if ip == clientIP {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s 没有访问权限", clientIP)
			c.Abort() //请求错误了使用abort，提前结束后续的handler
		}
	}
}
