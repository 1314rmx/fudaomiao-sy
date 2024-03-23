package service

import (
	"encoding/json"
	"gin-spider/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"math"
	"regexp"
	"strconv"
	"strings"
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

	for i := 0; i < len(kb.KbList); i++ {
		//获取周次
		zcd := kb.KbList[i].Zcd
		//获取节次
		jc := kb.KbList[i].Jc
		pattern := `\d+`
		re := regexp.MustCompile(pattern)
		matchs := re.FindAllString(zcd, -1)
		//匹配正则得到周次第一个数字
		zcd_start, _ := strconv.Atoi(matchs[0])
		//匹配正则得到周次第二个数字
		zcd_end, _ := strconv.Atoi(matchs[1])
		match := re.FindAllString(jc, -1)
		//匹配获取节次第一个数字
		jc_start := match[0]
		count := 0
		if strings.Contains(zcd, "单") {
			//为周数组开辟空间,向上取整
			zc := int(math.Ceil(float64((float64(zcd_end) - float64(zcd_start) + 1.0) / 2.0)))
			kb.KbList[i].Weeks = make([]string, zc)
			for j := zcd_start; j <= zcd_end; j++ {
				if j%2 != 0 {
					//周次
					kb.KbList[i].Weeks[count] = strconv.Itoa(j)
					//节数
					kb.KbList[i].SectionCount = "2"
					//开始的节
					kb.KbList[i].Section = jc_start
					count++
				}
			}
		} else if strings.Contains(zcd, "双") {
			//为周数组开辟空间,向上取整
			zc := int(math.Ceil(float64((float64(zcd_end) - float64(zcd_start) + 1.0) / 2.0)))
			kb.KbList[i].Weeks = make([]string, zc)
			for j := zcd_start; j <= zcd_end; j++ {
				if j%2 == 0 {
					kb.KbList[i].Weeks[count] = strconv.Itoa(j)
					kb.KbList[i].SectionCount = "2"
					kb.KbList[i].Section = jc_start
					count++
				}
			}
		} else {
			//为周数组开辟空间
			kb.KbList[i].Weeks = make([]string, zcd_end-zcd_start+1)
			for j := zcd_start; j <= zcd_end; j++ {
				kb.KbList[i].Weeks[count] = strconv.Itoa(j)
				kb.KbList[i].SectionCount = "2"
				kb.KbList[i].Section = jc_start
				count++
			}
		}
	}

	context.JSON(200, gin.H{
		"code": 200,
		"data": kb,
		"msg":  "获取学期信息成功!",
	})
}
