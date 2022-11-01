package route

import "github.com/gin-gonic/gin"

func LoadRoute(r *gin.Engine) {

	//路由规则
	loginGroup := r.Group("/login")
	{
		loginGroup.POST("/")
	}
}
