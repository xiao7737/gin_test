package apis

import (
	model "gin_test/api/models"
	"gin_test/gmongo"
	"gin_test/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

//mongo  test data  {"username":"xiao444","age":"11"}
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

//mongo delete user by username
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

//mongo update age by username
func UpdateUserAgeByUsername(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	username := c.Query("username")
	age := c.Query("age")
	result, err := gmongo.UpdateOne("user", "user", "username", username, "age", age)
	if result == -1 {
		ginReturn.Response(http.StatusOK, msg.NO_ROWS, "")
		return
	}
	if err != nil {
		ginReturn.Response(http.StatusOK, msg.UPDATE_FALIED, "")
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, "")
}

//mongo获取集合的数量
func DataCountOfCollection(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	collection := c.Query("collection")
	name, size := gmongo.CollectionCount("user", collection) //collection=user
	if size == 0 {
		ginReturn.Response(http.StatusOK, msg.NO_ROWS, "empty collection")
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, gin.H{"collection": name, "count": size})
}

//mongo 获取全部数据
func GetAllData(c *gin.Context) {
	ginReturn := msg.Gin{C: c}
	name, data, err := gmongo.FindAll("user", "user")
	if err != nil {
		ginReturn.Response(http.StatusOK, msg.ERROR, "get data error")
		return
	}
	ginReturn.Response(http.StatusOK, msg.SUCCESS, gin.H{"collection": name, "data": data})
}
