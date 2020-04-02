package apis

import (
	model "gin_test/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// get user list
func Users(c *gin.Context) {
	var userModel model.User
	result, err := userModel.Users()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未找到信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

// create  a new user
func Store(c *gin.Context) {
	var userModel model.User
	userModel.Username = c.Request.FormValue("username")
	userModel.Password = c.Request.FormValue("password")
	id, err := userModel.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Insert failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    id,
		"message": "insert success",
	})
}

func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Delete failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Delete success",
	})
}

func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	user.Username = c.Request.FormValue("username")
	user.Age, err = strconv.Atoi(c.Request.FormValue("age"))
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}
