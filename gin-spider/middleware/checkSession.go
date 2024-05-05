package middleware

import (
	"gin-spider/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"strings"
)

func CheckSession() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer model.Error(context)
		if strings.Contains(context.Request.URL.Path, "/login") || strings.Contains(context.Request.URL.Path, "/check") {
			context.Next()
			return
		}
		session := sessions.Default(context)
		if session.Get("username") == nil || model.UserCollector[session.Get("username").(string)] == nil {
			context.JSON(200, gin.H{
				"code": 400,
				"data": nil,
				"msg":  "请先登录!",
			})
			context.Abort()
			return
		}
		c := model.UserCollector[session.Get("username").(string)].Clone()
		c.AllowURLRevisit = true
		flag := false
		c.OnResponse(func(r *colly.Response) {
			if strings.Contains(string(r.Body), "教学管理信息服务平台") {
				flag = true
			}
		})
		c.Visit("https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xtgl/index_initMenu.html?jsdm=")
		if flag {
			context.Next()
			return
		} else {
			context.JSON(200, gin.H{
				"code": 401,
				"data": nil,
				"msg":  "cookie失效!",
			})
			context.Abort()
			return
		}
	}
}
