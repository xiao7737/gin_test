package router

import (
	. "gin_test/api/apis"
	"gin_test/validator/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func InitRouter() *gin.Engine { // Engine结构体包含了 RouterGroup
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

// RouterGroup是对路由树的包装
// 路由树的每个节点会挂载若干函数和中间件构成一个函数处理链：这也是中间件的加载过程，将中间件加载到对应的节点上面
// Engine 结构体继承了 RouterGroup ，所以 Engine 直接具备了 RouterGroup 所有的路由管理功能
// 这是为什么在 Hello World例子中，可以直接使用 Engine 对象来定义路由规则。同时 RouteGroup 对象里面还会包含一个 Engine 的指针
// 通过调用 Engine.addRoute 方法将请求处理器挂接到路由树中，路由规则被分成9颗前缀树，对应9中http方法
// context可以看做gin对http.Request对象的包装
// Gin不支持https，官方建议用Nginx转发https请求到Gin
// 路由节点挂载函数链：Gin提供了中间件，只有函数链的尾部才是业务处理
// 请求流程：接到请求时，在路由树找到相应的节点，组成请求处理链，构造一个Context对象，依次调用Next()方法进行请求处理
// Abort()中断请求链，只会中断当前操作，后续的中间件或操作会继续执行，原理：将context.index设置较大,让Next()的调用循环立即结束
// Use() 注册中间件
