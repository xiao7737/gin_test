package msg

import (
	"github.com/gin-gonic/gin"
)

var msgMap = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "fail",
	INVALID_PARAMS:  "请求参数错误",
	VALID_FAILED:    "参数验证失败",
	ERROR_EXIST_TAG: "已存在该标签名称",
	NO_ROWS:         "没有记录",
	INSERT_FALIED:   "新增失败",
	DELETE_FALIED:   "删除失败",
	UPDATE_FALIED:   "编辑失败",
}

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  GetErrorMsg(errCode), //也可以采用http包的statusText(code)
		"data": data,
	})
	return
}

func GetErrorMsg(code int) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}
	return msgMap[ERROR]
}