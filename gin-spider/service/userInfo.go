package service

import (
	"encoding/json"
	"fmt"
	"gin-spider/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type UserInfoService struct {
}

func (userInfo UserInfoService) UserInfo(context *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"data": nil,
				"msg":  "发生错误，建议注销再登录!",
			})
			context.Abort()
			return
		}
	}()
	if GetGnmkdmKey(context)["usertype"] == "teacher" {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "教师暂不允许查看个人信息!",
		})
		context.Abort()
		return
	}
	var semesterInfo model.SemesterInfo
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "session为空，请先登录!",
		})
		context.Abort()
	}
	c := model.UserCollector[session.Get("username").(string)].Clone()
	c.AllowURLRevisit = true
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
		err := json.Unmarshal(r.Body, &semesterInfo)
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"data": nil,
				"msg":  "获取个人信息失败!",
			})
			context.Abort()
			return
		}
	})
	info_url := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xsxxxggl/xsxxwh_cxCkDgxsxx.html?vpn-12-o1-jwgl.hjnu.edu.cn:82&gnmkdm=" + GetGnmkdmKey(context)["userinfo"]
	c.Visit(info_url)
	context.JSON(200, gin.H{
		"code": 200,
		"data": semesterInfo,
		"msg":  "获取个人信息成功!",
	})
}
