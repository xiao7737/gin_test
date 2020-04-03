package router

import (
	. "gin_test/api/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/users", Users)        //列表
		v1.POST("/user", Store)        //新增
		v1.DELETE("user/:id", Destroy) //删除
		v1.PUT("/user/:id", Update)    //编辑
	}
	return router
}
