package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/:name/:age", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"age":  c.Param("age"),
		})
	})
	r.Run(":8081") // server default port:8080
}
