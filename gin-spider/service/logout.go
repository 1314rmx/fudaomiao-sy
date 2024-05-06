package service

import (
	"gin-spider/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LogoutService struct {
}

func (logout LogoutService) Logout(context *gin.Context) {
	defer model.Error(context)
	session := sessions.Default(context)
	delete(model.UserCollector, session.Get("username").(string))
	session.Set("username", nil)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出成功",
		"data": nil,
	})
}
