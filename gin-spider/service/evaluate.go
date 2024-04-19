package service

import (
	"encoding/json"
	"gin-spider/model"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type EvaluateService struct {
}

var waitg sync.WaitGroup
var boolChan chan bool = make(chan bool)

func evaluate_spider(c *colly.Collector, url string, data map[string]string) {
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(string(r.Body), "成功") {
			//评价成功不知道返回啥就没写
		}
	})
	c.Post(url, data)
	waitg.Done()
}

func (evaluate EvaluateService) Evaluate(context *gin.Context) {
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
	var courseInfo model.CourseInfo
	c.AllowURLRevisit = true
	getinfourl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xspjgl/xspj_cxXspjIndex.html?doType=query&gnmkdm=N401605"
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &courseInfo)
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"msg":  "获取课表信息失败!",
				"data": nil,
			})
			context.Abort()
			return
		}
	})
	c.Visit(getinfourl)

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
	py := []string{"我很欣赏老师的教学态度，他总是以严谨认真的态度帮助我们掌握知识，对我们的要求也非常高。", "老师待人热情，乐于答疑，态度认真，对我们的学习和成长非常关心。", "老师的教学风格非常活泼，能够调动学生的积极性，让大家在课堂上保持清醒。", "老师认真负责，课堂秩序维护得当，能很好地激发我们的学习积极性。", "老师工作态度认真，能够很好地完成教学任务。", "老师上课认真负责，讲解细致，能够帮助学生掌握知识。", "老师认真负责，对学生要求严格。", "老师讲课生动有趣，很能吸引学生的注意力，让我们对课程内容产生了浓厚的兴趣。"}
	for i := 0; i < len(courseInfo.Items); i++ {
		evaluate_data["kch_id"] = courseInfo.Items[i].KchID
		evaluate_data["jxb_id"] = courseInfo.Items[i].JxbID
		evaluate_data["jgh_id"] = courseInfo.Items[i].JghID
		// 设置随机种子以确保每次调用都能产生不同的随机数
		rand.Seed(time.Now().UnixNano())
		// 获取数组长度
		n := len(py)
		// 生成一个[0, n)范围内的随机索引
		index := rand.Intn(n)
		// 对应索引处的元素
		evaluate_data["modelList[0].xspjList[0].childXspjList[8].zgpj"] = py[index]
		evaluate_data["modelList[0].xspjList[0].childXspjList[9].zgpj"] = py[index]
		waitg.Add(1)
		go evaluate_spider(c, evaluateurl, evaluate_data)
	}
	waitg.Wait()
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "评价成功",
	})
}
