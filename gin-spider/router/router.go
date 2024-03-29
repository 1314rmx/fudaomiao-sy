package router

import (
	"gin-spider/middleware"
	"gin-spider/service"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	gin.ForceConsoleColor()
	//gin.SetMode(gin.TestMode)
	app := gin.Default()
	app.Use(middleware.Cors())
	InitRouter(app)
	return app
}

func InitRouter(app *gin.Engine) {
	app.POST("/login", service.LoginService{}.Login)
	app.GET("/scores", service.QueryService{}.GetScoreList)
	app.GET("/courses", service.CurriculumService{}.Curriculum)
	app.GET("/evaluate", service.EvaluateService{}.Evaluate)
	app.GET("/userInfo", service.UserInfoService{}.UserInfo)
}
