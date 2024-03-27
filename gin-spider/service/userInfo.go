package service

import (
	"encoding/json"
	"gin-spider/model"
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
				"msg":  "发生错误!",
			})
			context.Abort()
			return
		}
	}()
	var semesterInfo model.SemesterInfo
	c := model.Collector.Clone()
	c.AllowURLRevisit = true
	c.OnResponse(func(r *colly.Response) {
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
	info_url := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xsxxxggl/xsxxwh_cxCkDgxsxx.html?vpn-12-o1-jwgl.hjnu.edu.cn:82&gnmkdm=N100801"
	c.Visit(info_url)
	context.JSON(200, gin.H{
		"code": 200,
		"data": semesterInfo,
		"msg":  "获取个人信息成功!",
	})
}
