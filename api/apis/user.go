package apis

import (
	model "gin_test/api/models"
	"gin_test/gmongo"
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

	//redis action
	//redis获取key
	/*data, _ := gredis.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
	})*/

	//redis设置key，并设置过期时间
	/*if err := gredis.Set("username_redis", "set_name", 100); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": "OK",
		})
	}*/

	// redis删除key
	/*if deleteRes, err := gredis.Delete("username3"); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": deleteRes, //删除一个不存在的key，返回0，经过Bool转换后为false，
		})
	}*/
}

// mongo action
// add test data in your mongodb:
// db.user.insert({"username":"xiaoxc","age":99})
func GetUserByNameFromMongo(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	var userModel model.User
	username := c.Query("username")
	data := gmongo.FindOne("user", "user", "username", username) //查询user数据库的user集合，username=xiaoxc
	if err := data.Decode(&userModel); err != nil {              //mongo返回为bson，需要解析
		ginReturn.Response(http.StatusOK, msg.NO_ROWS, "")
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, gin.H{"data": userModel})
}

// test data  {"username":"xiao444","age":"11"}
func InsertUserIntoMongo(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	var userModel model.User
	if err := c.ShouldBindJSON(&userModel); err != nil {
		ginReturn.Response(http.StatusBadRequest, msg.VALID_FAILED, err)
		return
	}
	data, err := gmongo.InsertOne("user", "user", &userModel)
	if err != nil {
		ginReturn.Response(http.StatusOK, msg.INSERT_FALIED, "")
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, gin.H{"id": data.InsertedID})
}

func DeleteUserFromMongo(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	username := c.Query("username")
	data, err := gmongo.DeleteOne("user", "user", "username", username) //username=xiao444
	if data == 0 || err != nil {
		ginReturn.Response(http.StatusOK, msg.DELETE_FALIED, "can not find the data which you want to delete")
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, "delete success")
}
func updateUserFromMongo(c *gin.Context) {

}
