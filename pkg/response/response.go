package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const PARAMERR = 10001
const CALLLIBFAIL = 10002

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
	})
}

func SuccessNoData(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    nil,
	})
}

func FailWithSys(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    -1,
		"message": "服务器错误",
		"data":    nil,
	})
}

func FailAndCode(c *gin.Context, code int64, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": msg,
		"data":    nil,
	})
}
