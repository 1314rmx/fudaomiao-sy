package router

import (
	"gin-spider/middleware"
	"gin-spider/model"
	"gin-spider/service"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func App() *gin.Engine {
	gin.ForceConsoleColor()
	//gin.SetMode(gin.TestMode)
	app := gin.Default()
	app.Use(middleware.Cors())
	InitRouter(app)
	InitColly()
	return app
}

func InitRouter(app *gin.Engine) {
	app.POST("/login", service.LoginService{}.Login)
	app.GET("/scores", service.QueryService{}.GetScoreList)
	app.GET("/courses", service.CurriculumService{}.Curriculum)
	app.GET("/evaluate", service.EvaluateService{}.Evaluate)
	app.GET("/userInfo", service.UserInfoService{}.UserInfo)
	app.GET("/check", service.CheckNeedCaptcha)
}

func InitColly() {
	model.Collector = colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
	)
	model.Collector.Limit(&colly.LimitRule{
		Parallelism: 100,
	})
	model.Collector.AllowURLRevisit = true
	hjurl := "https://webvpn.hjnu.edu.cn/http/736e6d702d6167656e74636f6d6d756efeb964a4bb9598689c84a24f3fe5e0/authserver/login?service=http%3A%2F%2Fjwgl.hjnu.edu.cn%3A82%2Fsso%2Fjziotlogin"
	model.Collector.Visit(hjurl)
}
