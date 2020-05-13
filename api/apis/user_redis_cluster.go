package apis

import (
	"gin_test/go_redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserFromRedisCluster(c *gin.Context) {
	key := c.Query("redis_cluster_key") //name or username
	data, err := go_redis.Get(key)
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
}
