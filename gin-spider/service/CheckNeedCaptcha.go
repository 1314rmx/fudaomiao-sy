package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"gin-spider/model"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func CheckNeedCaptcha(context *gin.Context) {
	username := context.Query("stuId")
	if username == "" {
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  "学号为空，请输入学号",
			"data": nil,
		})
		context.Abort()
	}
	fmt.Println(username)
	model.UserCollector[username] = colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
	)
	model.UserCollector[username].Limit(&colly.LimitRule{
		Parallelism: 100,
	})
	model.UserCollector[username].AllowURLRevisit = true

	login_url := "https://webvpn.hjnu.edu.cn/http/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/login?service=http%3A%2F%2Fjwgl.hjnu.edu.cn%3A82%2Fsso%2Fjziotlogin"
	model.UserCollector[username].OnResponse(func(r *colly.Response) {
		_, exists := context.Get("cookie")
		if exists {
			return
		}
		fmt.Println(r.Headers.Values("Set-Cookie"))
		context.Set("cookie", r.Headers.Values("Set-Cookie"))
	})
	model.UserCollector[username].Visit(login_url)

	checkNeedCaptcha_url := "https://webvpn.hjnu.edu.cn/https/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/checkNeedCaptcha.htl?username=" + username
	flag := false

	model.UserCollector[username].OnResponse(func(r *colly.Response) {
		//fmt.Println(string(r.Body))
		if strings.Contains(string(r.Body), "true") {
			flag = true
		} else {
			flag = false
		}
	})
	model.UserCollector[username].Visit(checkNeedCaptcha_url)
	if !flag {
		context.JSON(200, gin.H{
			"code": 201,
			"msg":  "hide",
			"data": nil,
		})
		context.Abort()
		return
	}
	captcha_imgurl := "https://webvpn.hjnu.edu.cn/https/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/getCaptcha.htl"
	var base64_str string
	model.UserCollector[username].OnResponse(func(r *colly.Response) {
		imgBytes, err := ioutil.ReadAll(bytes.NewReader(r.Body))
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"msg":  "验证码发送错误!",
				"data": nil,
			})
			context.Abort()
			return
		}
		base64_str = base64.StdEncoding.EncodeToString(imgBytes)
	})
	model.UserCollector[username].Visit(captcha_imgurl)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "验证码show",
		"data": base64_str,
	})
}
