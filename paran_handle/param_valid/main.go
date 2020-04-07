package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string `form:"name" binding:"required"` // max,min,|,exists,omitempty,len,eq,ne,gte,lte
	Age  int    `form:"age" binding:"required,gt=20"`
}

func main() {
	r := gin.Default()
	r.GET("/param_valid", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, "%v", err)
			c.Abort()
			return // 关闭出现的tag参数
		}
		c.String(http.StatusOK, "%v", person)
	})
	r.Run()
	// localhost:8080/param_valid?name=1&age=21
}
