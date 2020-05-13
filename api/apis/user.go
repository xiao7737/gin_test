package apis

import (
	model "gin_test/api/models"
	"gin_test/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// c.Query("key")  获取url里面的参数   同 c.Param("key")
// c.DefaultQuery("key", default_value)
// c.PostForm("key")   body里面的form表单
// c.DefaultPostForm("key", default_value)

// get user list
func Users(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	var userModel model.User
	result, count, err := userModel.Users()
	if err != nil {
		ginReturn.Response(http.StatusOK, msg.NO_ROWS, nil)
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, gin.H{"data": result, "count": count})
}

// create  a new user
func Store(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	var userModel model.User
	// 接受json格式，直接绑定在User结构体上，并进行json转化，可以使用结构体的tag进行相关验证
	if err := c.ShouldBindJSON(&userModel); err != nil {
		ginReturn.Response(http.StatusBadRequest, msg.VALID_FAILED, err)
		return
	}
	id, err := userModel.Insert()
	if err != nil {
		ginReturn.Response(http.StatusBadRequest, msg.INSERT_FALIED, err)
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, gin.H{"id": id})
}

func Destroy(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	err = user.Destroy(id)
	if err != nil {
		ginReturn.Response(http.StatusOK, msg.DELETE_FALIED, err)
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, nil)
}

func Update(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := c.ShouldBindJSON(&user); err != nil {
		ginReturn.Response(http.StatusOK, msg.VALID_FAILED, err)
		return
	}
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		ginReturn.Response(http.StatusOK, msg.UPDATE_FALIED, err)
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, nil)
}

func GetUserById(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.GetById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// 根据用户名的like
func GetUserByName(c *gin.Context) {
	var user model.User
	result, err := user.GetUserByName(c.Query("username"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
