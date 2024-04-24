package service

import (
	"encoding/json"
	"fmt"
	"gin-spider/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"math"
	"strconv"
	"strings"
	"time"
)

type ClassRoomService struct {
}

var campus map[string]string = make(map[string]string)
var lhmap map[string]string = make(map[string]string)

func (classRoom ClassRoomService) GetClassRoom(context *gin.Context) {
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "请先登录!",
		})
		context.Abort()
	}
	campusname := context.Query("campus")                    //校区，例如主校区
	week, _ := strconv.Atoi(context.Query("week"))           //传第几周，例如
	xqj := context.Query("xingqi")                           //传星期几
	lh := context.Query("lh")                                //楼栋
	start_jie, _ := strconv.Atoi(context.Query("start_jie")) //开始节
	end_jie, _ := strconv.Atoi(context.Query("end_jie"))     //结束节
	var total float64 = 0
	for ; start_jie <= end_jie; start_jie++ {
		total += math.Pow(2, float64(start_jie-1))
	}
	zcd := int(math.Pow(2, float64(week-1)))
	c := model.UserCollector[session.Get("username").(string)].Clone()
	c.AllowURLRevisit = true
	var parameter map[string]string = make(map[string]string)
	c.OnResponse(func(r *colly.Response) {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			return
		}
		if len(campus) == 0 {
			dom.Find("#xqh_id").Find("option").Each(func(i int, selection *goquery.Selection) {
				campusName := selection.Text()
				campusValue, _ := selection.Attr("value")
				campus[campusName] = campusValue
			})
		}
		if len(lhmap) == 0 {
			dom.Find("#lh").Find("option").Each(func(i int, selection *goquery.Selection) {
				lhName := selection.Text()
				lhValue, _ := selection.Attr("value")
				lhmap[lhName] = lhValue
			})
		}
		parameter["fwzt"], _ = dom.Find("#fwzt").Attr("value")
		parameter["xnm"], _ = dom.Find("#xnm").Attr("value")
		parameter["xqm"], _ = dom.Find("#xqm").Attr("value")
	})
	url := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/cdjy/cdjy_cxKxcdlb.html?gnmkdm=N2155&layout=default"
	c.Visit(url)

	data := map[string]string{
		"fwzt":                   parameter["fwzt"],
		"xqh_id":                 campus[campusname],
		"xnm":                    parameter["xnm"],
		"xqm":                    parameter["xqm"],
		"cdlb_id":                "",
		"cdejlb_id":              "",
		"qszws":                  "",
		"jszws":                  "",
		"cdmc":                   "",
		"lh":                     lhmap[lh],
		"jyfs":                   "0",
		"cdjylx":                 "",
		"zcd":                    strconv.Itoa(zcd),
		"xqj":                    xqj,
		"jcd":                    fmt.Sprintf("%.0f", total), //strconv.FormatFloat(total, 10, 2, 64),
		"_search":                "false",
		"nd":                     strconv.FormatInt(time.Now().Unix(), 10),
		"queryModel.showCount":   "100",
		"queryModel.currentPage": "1",
		"queryModel.sortName":    "cdbh",
		"queryModel.sortOrder":   "asc",
		"time":                   "1",
	}
	var classroom model.ClassRoom
	var flag bool = false
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &classroom)
		if err != nil {
			flag = false
			return
		} else {
			flag = true
		}
	})
	classRoomUrl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/cdjy/cdjy_cxKxcdlb.html?vpn-12-o1-jwgl.hjnu.edu.cn:82&doType=query&gnmkdm=N2155"
	c.Post(classRoomUrl, data)
	if !flag {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "获取失败!",
		})
		return
	}
	context.JSON(200, gin.H{
		"code": 200,
		"data": classroom,
		"msg":  "获取成功!",
	})
}
