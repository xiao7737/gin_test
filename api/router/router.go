package router

import "github.com/gin-gonic/gin"
import . "gin_test/api/apis" //使用不加包名

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", Users)
	return router
}
