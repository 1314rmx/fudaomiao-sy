package service

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"gin-spider/model"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"io/ioutil"
	"strings"
)

func CheckNeedCaptcha(context *gin.Context) {
	username := context.Query("username")
	fmt.Println(username)
	checkNeedCaptcha_url := "https://webvpn.hjnu.edu.cn/https/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/checkNeedCaptcha.htl?username=" + username
	flag := false
	model.Collector.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
		if strings.Contains(string(r.Body), "true") {
			flag = true
		} else {
			flag = false
		}
	})
	model.Collector.Visit(checkNeedCaptcha_url)
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
	model.Collector.OnResponse(func(r *colly.Response) {
		imgBytes, err := ioutil.ReadAll(bytes.NewReader(r.Body))
		if err != nil {
			context.JSON(200, gin.H{
				"code": 400,
				"msg":  "发送错误!",
				"data": nil,
			})
			context.Abort()
			return
		}
		base64_str = base64.StdEncoding.EncodeToString(imgBytes)
	})
	model.Collector.Visit(captcha_imgurl)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "show",
		"data": base64_str,
	})
}
