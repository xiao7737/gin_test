package apis

import (
	model "gin_test/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get user list
func Users(c *gin.Context) {
	var userModel model.User
	userModel.Username = c.Request.FormValue("username")
	userModel.Password = c.Request.FormValue("password")
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
