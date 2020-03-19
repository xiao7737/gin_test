package main

// 参数绑定结构体
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Age      int       `form:"age"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/test_struct", testing)
	r.Run()
	// localhost:8080/test_struct?name=tony&age=1000&birthday=2020-03-19
}

func testing(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusOK, "error:%v", err)
	}
}
