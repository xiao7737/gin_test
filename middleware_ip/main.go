package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

const requestID = "requestID"

func main() {
	r := gin.Default()
	//两个中间件：①检测路由，②给每个请求添加一个requestId
	r.Use(IPMiddleware(), func(c *gin.Context) {
		c.Set(requestID, rand.Int())
		c.Next() // 每个中间件添加next，继续执行不中断后续
	})

	r.GET("/test_middleware", func(c *gin.Context) {
		Mes := gin.H{
			"message": "ip is allow!",
		}
		if reqID, exists := c.Get(requestID); exists {
			Mes[requestID] = reqID
		}
		c.JSON(http.StatusOK, Mes)
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
