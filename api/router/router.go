package router

import "github.com/gin-gonic/gin"
import . "gin_test/api/apis" //使用不加包名

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", Users) //获取所有用户： 127.0.0.1:9999/users
	router.POST("/user", Store) //新增一个用户： 127.0.0.1:9999/user
	return router
}
