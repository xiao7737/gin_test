package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// 获取body里面的参数
func main() {
	r := gin.Default()
	r.POST("/param_body", func(context *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		context.String(http.StatusOK, string(bodyBytes))
	})
	r.Run()
	// curl -X POST localhost:8080/param_body -d "{"body":"param_body_fetch"}"
}
