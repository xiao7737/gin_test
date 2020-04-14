package IpMiddleware

import "github.com/gin-gonic/gin"

// 自定义中间件就可以参照logger进行设置
func IpMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// white ip list，当然这里不应该这么简单==！
		ipList := []string{
			"127.0.0.1",
			"localhost",
			"::1",
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
			c.Abort() // 结束当前的handler
		}
		c.Next() // 每个中间件添加next，继续执行不中断后续handler
	}
}
