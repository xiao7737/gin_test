package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		context.String(200, "GET method")
	})
	r.POST("/post", func(context *gin.Context) {
		context.String(200, "POST method")
	})

	// curl -X DELETE localhost:8080/delete
	r.Handle("DELETE", "/delete", func(context *gin.Context) {
		context.String(200, "DELETE method")
	})
	//
	r.Any("/any", func(context *gin.Context) {
		context.String(200, "Any method")
	})
	r.Run()
}
