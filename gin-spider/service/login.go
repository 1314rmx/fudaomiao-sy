package service

import (
	"gin-spider/model"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (login LoginService) Login(context *gin.Context) {
	username := context.PostForm("stuId")
	password := context.PostForm("password")
	if username == "" || password == "" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "学号或密码为空",
		})
		return
	}
	model.Collector = model.Initcolly(username, password)
	if model.Collector == nil {
		context.JSON(200, gin.H{
			"code": "400",
			"msg":  "账号或密码错误!",
		})
	} else {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功!",
		})
	}
}
