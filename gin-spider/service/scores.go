package service

import (
	"encoding/json"
	"fmt"
	"gin-spider/model"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type QueryService struct {
}

var wg sync.WaitGroup

func (query QueryService) GetScoreList(context *gin.Context) {
	if GetGnmkdmKey(context)["usertype"] == "teacher" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "暂时不支持老师账号查看分数!",
			"data": nil,
		})
		context.Abort()
		return
	}
	semestersChan := make(chan []semesterList)
	infoChan := make(chan model.SemesterInfo)
	scoreChan := make(chan model.Stuscore)
	go getSemester(context, semestersChan, infoChan)
	go Query(context, scoreChan)
	wg.Wait()
	semesters := <-semestersChan
	score := <-scoreChan
	info := <-infoChan
	for i := 0; i < len(score.Items); i++ {
		score.Items[i].Xm = info.Xm
		score.Items[i].Xh = info.Xh
	}
	context.JSON(200, gin.H{
		"code":     200,
		"msg":      "获取成绩成功!",
		"semester": semesters,
		"data":     score.Items,
	})
}

type semesterList struct {
	TermName string `json:"termName"`
}

func getSemester(context *gin.Context, semestersChan chan []semesterList, infoChan chan model.SemesterInfo) {
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
	wg.Add(1)
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
		err := json.Unmarshal(r.Body, &semesterInfo)
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
	info_url := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xsxxxggl/xsxxwh_cxCkDgxsxx.html?vpn-12-o1-jwgl.hjnu.edu.cn:82&gnmkdm=" + GetGnmkdmKey(context)["userinfo"]
	c.Visit(info_url)

	xz, _ := strconv.Atoi(semesterInfo.Xz)
	njdmID, _ := strconv.Atoi(semesterInfo.NjdmID)
	now := time.Now()
	year := now.Year()
	pastyearnum := year - njdmID
	if pastyearnum > xz {
		pastyearnum = xz
	}
	if now.Month() < 7 {
		pastyearnum = pastyearnum*2 - 1
	}
	var semesters = make([]semesterList, pastyearnum)
	for i := 0; i < pastyearnum; i++ {
		if i%2 == 0 {
			semesters[i].TermName = strconv.Itoa(njdmID) + "-" + strconv.Itoa(njdmID+1) + "学年第一学期"
		} else {
			semesters[i].TermName = strconv.Itoa(njdmID) + "-" + strconv.Itoa(njdmID+1) + "学年第二学期"
			njdmID += 1
		}
		if year <= njdmID {
			break
		}
	}
	semestersChan <- semesters
	infoChan <- semesterInfo
	wg.Done()
}

func Query(context *gin.Context, scoreChan chan model.Stuscore) {
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
	wg.Add(1)
	now := time.Now()
	year := now.Year()
	month := now.Month()
	var semester string
	var schoolyear string
	if month < 7 {
		semester = "1"
		schoolyear = strconv.Itoa(year)
	} else {
		semester = "2"
		schoolyear = strconv.Itoa(year)
	}
	termName := context.DefaultQuery("termName", "-1")
	if !strings.Contains(termName, "-1") {
		split_name := strings.Split(termName, "-")
		schoolyear = split_name[0]
		if strings.Contains(split_name[1], "一") {
			semester = "1"
		} else {
			semester = "2"
		}
	}
	if semester == "1" {
		semester = "3"
	} else if semester == "2" {
		semester = "12"
	} else {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "参数错误",
		})
		context.Abort()
		return
	}
	//访问成绩查询
	cjurl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/cjcx/cjcx_cxXsgrcj.html?doType=query"
	timestamp := now.UnixNano() / 1e6
	timestampStr := fmt.Sprintf("%d", timestamp)
	cjdata := map[string]string{
		"zd_fzdm":                GetGnmkdmKey(context)["score"] + "-xsxnm=" + schoolyear,
		"xqm":                    semester,
		"kcbj":                   "",
		"_search":                "false",
		"nd":                     timestampStr,
		"queryModel.showCount":   "50",
		"queryModel.currentPage": "1",
		"queryModel.sortName":    "",
		"queryModel.sortOrder":   "asc",
		"time":                   "1",
	}
	var score model.Stuscore
	session := sessions.Default(context)
	c := model.UserCollector[session.Get("username").(string)]
	c.AllowURLRevisit = true

	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &score)
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"data": nil,
				"msg":  "获取成绩失败",
			})
			context.Abort()
			return
		}
	})
	err := c.Post(cjurl, cjdata)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取成绩失败",
			"data": nil,
		})
		context.Abort()
		return
	}
	scoreChan <- score
	wg.Done()
}
