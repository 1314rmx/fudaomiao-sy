package service

import (
	"encoding/json"
	"fmt"
	"gin-spider/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type EvaluateService struct {
}

var waitg sync.WaitGroup

func evaluate_spider(c *colly.Collector, url string, data map[string]string) {
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(string(r.Body), "成功") {
			return
		}
	})
	c.Post(url, data)
	waitg.Done()
}

func (evaluateservice EvaluateService) Evaluate(context *gin.Context) {
	defer model.Error(context)
	if GetGnmkdmKey(context)["usertype"] == "teacher" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "暂时不支持老师账号评价!",
			"data": nil,
		})
		context.Abort()
		return
	}
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "请先登录!",
		})
		context.Abort()
	}
	c := model.UserCollector[session.Get("username").(string)]
	courseInfo := model.CourseInfo{}
	c.AllowURLRevisit = true
	getinfourl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xspjgl/xspj_cxXspjIndex.html?doType=query&gnmkdm=N401605"
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &courseInfo)
		if err != nil {
			context.Abort()
			return
		}
	})
	info_data := map[string]string{
		"nd":                     strconv.FormatInt(time.Now().Unix(), 10),
		"queryModel.currentPage": "1",
		"queryModel.showCount":   "500",
		"queryModel.sortName":    "kcmc",
		"queryModel.sortOrder":   "asc",
		"time":                   "0",
		"_search":                "false",
	}
	c.Post(getinfourl, info_data)

	evaluateurl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xspjgl/xspj_tjXspj.html?gnmkdm=N401605"
	evaluate_data := map[string]string{
		"ztpjbl":                  "100",
		"jszdpjbl":                "0",
		"xykzpjbl":                "0",
		"xsdm":                    "01",
		"modelList[0].pjmbmcb_id": "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].pjdxdm":     "01",
		"modelList[0].fxzgf":      "",
		"modelList[0].py":         "",
		"modelList[0].xspfb_id":   "",
		"modelList[0].xspjList[0].childXspjList[0].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[0].pjzbxm_id":    "0D51C017C73881C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[0].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[0].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[1].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[1].pjzbxm_id":    "0D51C017C73981C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[1].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[1].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[2].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[2].pjzbxm_id":    "0D51C017C73A81C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[2].pfdjdmb_id":   "5",
		"BF1A17FB541C189E0530101080AD234":                        "",
		"modelList[0].xspjList[0].childXspjList[2].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[3].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[3].pjzbxm_id":    "0D51C017C73B81C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[3].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[3].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[4].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[4].pjzbxm_id":    "0D51C017C73C81C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[4].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[4].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[5].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[5].pjzbxm_id":    "0D51C017C73D81C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[5].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[5].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[6].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[6].pjzbxm_id":    "0D51C017C73E81C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[6].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[6].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[7].pfdjdmxmb_id": "6D76E508187C8FAFE0530101080AB546",
		"modelList[0].xspjList[0].childXspjList[7].pjzbxm_id":    "0D51C017C73F81C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[7].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[7].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[8].zgpj":         "很好",
		"modelList[0].xspjList[0].childXspjList[8].pjzbxm_id":    "0D51C017C74081C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[8].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[8].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].childXspjList[9].zgpj":         "无",
		"modelList[0].xspjList[0].childXspjList[9].pjzbxm_id":    "0D52A8FA301682A2E0630101080AF13F",
		"modelList[0].xspjList[0].childXspjList[9].pfdjdmb_id":   "5BF1A17FB541C189E0530101080AD234",
		"modelList[0].xspjList[0].childXspjList[9].zsmbmcb_id":   "0D51C017C73681C8E0630101080AD51B",
		"modelList[0].xspjList[0].pjzbxm_id":                     "0D51C017C73781C8E0630101080AD51B",
		"modelList[0].pjzt":                                      "1",
		"tjzt":                                                   "1",
		"kch_id":                                                 courseInfo.Items[0].KchID,
		"jxb_id":                                                 courseInfo.Items[0].JxbID,
		"jgh_id":                                                 courseInfo.Items[0].JghID,
	}
	fmt.Println(len(courseInfo.Items))
	for i := 0; i < len(courseInfo.Items); i++ {
		if courseInfo.Items[i].Tjzt == "1" {
			continue
		}
		evaluate_data["kch_id"] = courseInfo.Items[i].KchID
		evaluate_data["jxb_id"] = courseInfo.Items[i].JxbID
		evaluate_data["jgh_id"] = courseInfo.Items[i].JghID
		// 设置随机种子以确保每次调用都能产生不同的随机数
		rand.Seed(time.Now().UnixNano())
		// 获取数组长度
		var evaluate *model.Evaluate = new(model.Evaluate)
		result := model.DB.Table("evaluate").Order("RAND()").Limit(1).First(evaluate)
		if result.Error != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "获取评价信息失败!",
				"data": nil,
			})
		}
		evaluate_data["modelList[0].xspjList[0].childXspjList[8].zgpj"] = evaluate.Content
		evaluate_data["modelList[0].xspjList[0].childXspjList[9].zgpj"] = evaluate.Content
		waitg.Add(1)
		go evaluate_spider(c, evaluateurl, evaluate_data)
	}
	waitg.Wait()
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "评价成功",
	})
}

func (evaluateservice EvaluateService) GetEvaluateInfo(context *gin.Context) {
	defer model.Error(context)
	if GetGnmkdmKey(context)["usertype"] == "teacher" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "暂时不支持老师账号评价!",
			"data": nil,
		})
		context.Abort()
		return
	}
	session := sessions.Default(context)
	c := model.UserCollector[session.Get("username").(string)]
	courseInfo := model.CourseInfo{}
	c.AllowURLRevisit = true
	getinfourl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xspjgl/xspj_cxXspjIndex.html?doType=query&gnmkdm=N401605"
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &courseInfo)
		if err != nil {
			context.Abort()
			return
		}
	})
	info_data := map[string]string{
		"nd":                     strconv.FormatInt(time.Now().Unix(), 10),
		"queryModel.currentPage": "1",
		"queryModel.showCount":   "500",
		"queryModel.sortName":    "kcmc",
		"queryModel.sortOrder":   "asc",
		"time":                   "0",
		"_search":                "false",
	}
	c.Post(getinfourl, info_data)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取课程信息成功!",
		"data": courseInfo,
	})
}
