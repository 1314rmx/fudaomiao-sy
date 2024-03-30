package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"strings"
)

// 生成随机字符串
func randomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err) // 更安全的错误处理机制应该替换此处
	}
	const letters = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678"
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

// AES加密字符串
func encryptAesString(data, keyStr, ivStr string) string {
	key := []byte(keyStr)
	iv := []byte(ivStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err) // 更安全的错误处理机制应该替换此处
	}

	padding := aes.BlockSize - len(data)%aes.BlockSize
	dataPad := data + strings.Repeat(string(padding), padding)
	encrypted := make([]byte, len(dataPad))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, []byte(dataPad))

	return base64.StdEncoding.EncodeToString(encrypted)
}

// 加密密码
func encryptPassword(pwd, key string) string {
	iv := randomString(aes.BlockSize) // IV通常为aes.BlockSize长
	// 随机字符串通过aes.BlockSize和随机字符串长度确定
	randomData := randomString(64)
	return encryptAesString(randomData+pwd, key, iv)
}

func encrypt_pwd(pwd string, key string) string {
	/*
		cmd := exec.Command("node", "js/encryptPWD.js", pwd, key)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		return string(out)
	*/
	encryptedPwd := encryptPassword(pwd, key)
	return encryptedPwd

}

type stuscore struct {
	Items []struct {
		Bj     string `json:"bj"`     //班级
		Cj     string `json:"cj"`     //成绩
		Jd     string `json:"jd"`     //绩点
		Jsxm   string `json:"jsxm"`   //老师
		Kcxzmc string `json:"kcxzmc"` //类型
		Sfxwkc string `json:"sfxwkc"` //是否学位课程
		Xf     string `json:"xf"`     //学分
		Xfjd   string `json:"xfjd"`   //学分绩点
		Xnmmc  string `json:"xnmmc"`  //2023-2024学年
		Xqmmc  string `json:"xqmmc"`  //学期数
	} `json:"items"`
}

func main() {
	/*hjurl := "https://webvpn.hjnu.edu.cn/http/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/login?service=http%3A%2F%2Fjwgl.hjnu.edu.cn%3A82%2Fsso%2Fjziotlogin"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
	)
	//获取密钥和固定值
	var pwdEncryptSalt string
	var execution string
	c.OnHTML(".form", func(e *colly.HTMLElement) {
		pwdEncryptSalt, _ = e.DOM.ParentsUntil(".form").Find("#pwdEncryptSalt").Attr("value")
		execution, _ = e.DOM.ParentsUntil(".form").Find("#execution").Attr("value")
	})
	c.Visit(hjurl)

	//加密
	enpwd := encrypt_pwd("163155", pwdEncryptSalt)

	data := map[string]string{
		"username":  "2023303255",
		"password":  enpwd,
		"captcha":   "",
		"_eventId":  "submit",
		"_cllt":     "userNameLogin",
		"dllt":      "generalLogin",
		"lt":        "",
		"execution": execution,
	}
	//模拟登录
	c.Post(hjurl, data)
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.OnHTML("#btn_yd", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "已阅读") {
			fmt.Println("已阅读")
		} else {
			fmt.Println("账号或密码错误!")
		}
		fmt.Println(e.Text)
	})
	c.Visit("https://webvpn.hjnu.edu.cn/login?cas_login=true")


	*/

	//hjurl := "https://webvpn.hjnu.edu.cn/http/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/login?service=http%3A%2F%2Fjwgl.hjnu.edu.cn%3A82%2Fsso%2Fjziotlogin"
	//c := colly.NewCollector(
	//	colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
	//)
	////获取密钥和固定值
	//var pwdEncryptSalt string
	//var execution string
	//c.OnHTML(".form", func(e *colly.HTMLElement) {
	//	pwdEncryptSalt, _ = e.DOM.ParentsUntil(".form").Find("#pwdEncryptSalt").Attr("value")
	//	execution, _ = e.DOM.ParentsUntil(".form").Find("#execution").Attr("value")
	//})
	//c.Visit(hjurl)
	//
	////加密
	//enpwd := encrypt_pwd("163155", pwdEncryptSalt)
	//fmt.Println(enpwd)
	//
	//data := map[string]string{
	//	"username":  "202330325",
	//	"password":  enpwd,
	//	"captcha":   "",
	//	"_eventId":  "submit",
	//	"_cllt":     "userNameLogin",
	//	"dllt":      "generalLogin",
	//	"lt":        "",
	//	"execution": execution,
	//}
	////模拟登录
	//c.Post(hjurl, data)
	//
	//var flag bool
	//c.OnHTML("#btn_yd", func(e *colly.HTMLElement) {
	//	if strings.Contains(e.Text, "已阅读") {
	//		flag = true
	//	} else {
	//		flag = false
	//	}
	//	fmt.Println(e.Text)
	//})
	//c.Visit("https://webvpn.hjnu.edu.cn/login?cas_login=true")
	//fmt.Println(flag)
	//
	////访问成绩查询
	//cjurl := "https://webvpn.hjnu.edu.cn/http-82/736e6d702d6167656e74636f6d6d756ef7af70e6fd979c73c7cfa35e64a8ed2b/jwglxt/cjcx/cjcx_cxXsgrcj.html?doType=query"
	//timestamp := time.Now().UnixNano() / 1e6
	//timestampStr := fmt.Sprintf("%d", timestamp)
	//data33 := map[string]string{
	//	"zd_fzdm":                "N305005-xsxnm=2023",
	//	"xqm":                    "3",
	//	"kcbj":                   "",
	//	"_search":                "false",
	//	"nd":                     timestampStr,
	//	"queryModel.showCount":   "15",
	//	"queryModel.currentPage": "1",
	//	"queryModel.sortName":    "",
	//	"queryModel.sortOrder":   "asc",
	//	"time":                   "2",
	//}
	////var dataMap map[string]interface{}
	//var score stuscore
	//
	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("---------------------成绩")
	//	err := json.Unmarshal(r.Body, &score)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//})
	//c.Post(cjurl, data33)
	//fmt.Println(score)

	//str := "1-15周(单)"
	//pattern := `\d+`
	//re := regexp.MustCompile(pattern)
	//matchs := re.FindAllString(str, -1)
	//fmt.Println(matchs)

	//fmt.Println((14 + 1) / 2.0)
	//fmt.Println(int(math.Ceil(float64((14 + 1) / 2.0))))
	//str := "星期一"
	//fmt.Println(str[6:])

	captcha_imgurl := "https://webvpn.hjnu.edu.cn/https/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/getCaptcha.htl"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 100,
	})
	c.AllowURLRevisit = true
	c.OnResponse(func(r *colly.Response) {
		imgBytes, err := ioutil.ReadAll(bytes.NewReader(r.Body))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(base64.StdEncoding.EncodeToString(imgBytes))
	})
	c.Visit(captcha_imgurl)

}
