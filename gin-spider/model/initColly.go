package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

var Collector *colly.Collector

func Initcolly(username string, pwd string) *colly.Collector {
	hjurl := "https://webvpn.hjnu.edu.cn/http/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/login?service=http%3A%2F%2Fjwgl.hjnu.edu.cn%3A82%2Fsso%2Fjziotlogin"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 100,
	})
	c.AllowURLRevisit = true
	//获取密钥和固定值
	var pwdEncryptSalt string
	var execution string
	c.OnHTML(".form", func(e *colly.HTMLElement) {
		pwdEncryptSalt, _ = e.DOM.ParentsUntil(".form").Find("#pwdEncryptSalt").Attr("value")
		execution, _ = e.DOM.ParentsUntil(".form").Find("#execution").Attr("value")
	})
	c.Visit(hjurl)

	//加密
	enpwd := encrypt_pwd(pwd, pwdEncryptSalt)
	fmt.Println(enpwd)

	data := map[string]string{
		"username":  username,
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

	var flag bool
	c.OnHTML("#btn_yd", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "已阅读") {
			flag = true
		} else {
			flag = false
		}
	})
	c.Visit("https://webvpn.hjnu.edu.cn/login?cas_login=true")
	fmt.Println(flag)
	if flag {
		return c
	}
	return nil
}

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

	//cmd := exec.Command("node", "js/encryptPWD.js", pwd, key)
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return string(out)

	encryptedPwd := encryptPassword(pwd, key)
	return encryptedPwd

}
