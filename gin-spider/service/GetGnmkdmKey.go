package service

import (
	"gin-spider/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

func GetGnmkdmKey(context *gin.Context) map[string]string {
	defer model.Error(context)
	session := sessions.Default(context)
	c := model.UserCollector[session.Get("username").(string)].Clone()
	c.AllowURLRevisit = true
	gnmkdmkey := make(map[string]string, 4)
	c.Visit("https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xtgl/index_initMenu.html?jsdm=")
	c.OnResponse(func(r *colly.Response) {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			return
		}
		myDiv := dom.Find("#myDiv").Text()
		if strings.Contains(myDiv, "学生") {
			gnmkdmkey["usertype"] = "student"
		} else {
			gnmkdmkey["usertype"] = "teacher"
		}
		var userinfo *regexp.Regexp
		if gnmkdmkey["usertype"] == "student" {
			userinfo = regexp.MustCompile(`clickMenu\(&#39;(?P<key>[0-9a-zA-Z]+)&#39;[^<>()]+查询个人信息`)
			gnmkdmkey["userinfo"] = userinfo.FindStringSubmatch(string(r.Body))[1]
		} else {
			userinfo = regexp.MustCompile(`clickMenu\(&#39;(?P<key>[0-9a-zA-Z]+)&#39;[^<>()]+个人信息查询`)
			gnmkdmkey["userinfo"] = userinfo.FindStringSubmatch(string(r.Body))[1]
		}
		kb := regexp.MustCompile(`clickMenu\(&#39;(?P<key>[0-9a-zA-Z]+)&#39;,[^<>()]+,&#39;个人课表查询&#39;`)
		gnmkdmkey["kb"] = kb.FindStringSubmatch(string(r.Body))[1]
		scores := regexp.MustCompile(`clickMenu\(&#39;(?P<key>[0-9a-zA-Z]+)&#39;,[^<>()]+,&#39;学生成绩查询&#39;`)
		gnmkdmkey["score"] = scores.FindStringSubmatch(string(r.Body))[1]
	})
	c.Visit("https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/xtgl/index_initMenu.html?jsdm=")
	return gnmkdmkey
}
