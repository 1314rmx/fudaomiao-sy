package service

import (
	"gin-spider/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LogoutService struct {
}

func (logout LogoutService) Logout(context *gin.Context) {
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "session为空，请先登录!",
		})
		context.Abort()
	}
	session.Set("username", nil)
	delete(model.UserCollector, session.Get("username").(string))
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
		"data": nil,
	})
}
