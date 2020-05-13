package apis

import (
	"gin_test/gredis"
	"github.com/gin-gonic/gin"
	"net/http"
)

//redis action  采用redigo方式
func GetUserFromRedis(c *gin.Context) {
	key := c.Query("redis_key") //name or username
	data, err := gredis.Get(key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": data,
		})
	}

	//查询redis中是否存在
	/*exists := gredis.Exists(c.Query("username"))
	  c.JSON(http.StatusOK, gin.H{
	  	"code": 1,
	  	"exists": exists,
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
	  }
	*/
}
