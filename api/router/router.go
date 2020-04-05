package router

import (
	. "gin_test/api/apis"
	"gin_test/validator/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
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

	// 注册验证器--用户名处理
	// 验证器验证未通过，返回码400
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("NameValid", user.NameValid)
	}
	return router
}
