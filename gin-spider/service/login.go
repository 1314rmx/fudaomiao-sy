package service

import (
	"fmt"
	"gin-spider/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (login LoginService) Login(context *gin.Context) {
	username := context.PostForm("stuId")
	password := context.PostForm("password")
	captcha := context.DefaultPostForm("captcha", "")
	if username == "" || password == "" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "学号或密码为空",
		})
		return
	}
	keys := make([]string, 0, len(model.UserCollector))
	for key := range model.UserCollector {
		keys = append(keys, key)
	}
	flag := model.Initcolly(username, password, captcha, context)
	if !flag {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "账号或密码错误!",
		})
	} else {
		session := sessions.Default(context)
		session.Options(sessions.Options{
			MaxAge: 60 * 60 * 24 * 30,
		})
		session.Set("username", username)
		err := session.Save()
		if err != nil {
			fmt.Println("session保存失败", err)
			return
		}
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功!",
		})
	}
}
