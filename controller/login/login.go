package login

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-cms/pkg/logger"
	"go-cms/pkg/response"
)

type User struct {
	gorm.Model
	UserName   string `json:"userName" binding:"required"` //binding:"required"为空时报错
	PassWord   string `json:"passWord" binding:"required"`
	Permission string `json:"permission" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		logger.Log.WithField("err", err).Error("call ShouldBind fail")
		response.FailAndCode(c, response.PARAMERR, "参数错误")
		return
	}
	logger.Log.WithField("tag", 1222).Info("ok")

}
