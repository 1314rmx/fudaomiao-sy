package service

import (
	"encoding/json"
	"gin-spider/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"strconv"
	"time"
)

type CurriculumService struct {
}

func (curriculum CurriculumService) Curriculum(context *gin.Context) {
	c := model.Collector.Clone()
	var kb model.Curriculum
	c.AllowURLRevisit = true
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &kb)
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"data": nil,
				"msg":  "获取学期信息失败!",
			})
			context.Abort()
			return
		}
	})
	curriculum_url := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/kbcx/xskbcx_cxXsgrkb.html?vpn-12-o1-jwgl.hjnu.edu.cn:82&gnmkdm=N2151"
	now := time.Now()
	year := now.Year()
	month := now.Month()
	xq := "2"
	if month < 9 {
		year = year - 1
		xq = "12" //12代表第二学期
	} else {
		xq = "3" //3代表第一学期
	}
	curriculum_data := map[string]string{
		"xnm":  strconv.Itoa(year),
		"xqm":  xq,
		"kzlx": "ck",
		"xsdm": "",
	}
	c.Post(curriculum_url, curriculum_data)
	context.JSON(200, gin.H{
		"code": 200,
		"data": kb,
		"msg":  "获取学期信息成功!",
	})
}
