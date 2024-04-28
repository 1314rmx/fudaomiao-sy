package router

import (
	"gin-spider/middleware"
	"gin-spider/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	gin.ForceConsoleColor()
	//gin.SetMode(gin.TestMode)
	app := gin.Default()
	store := cookie.NewStore([]byte("1314rmx"))
	app.Use(sessions.Sessions("sessionID", store))
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
	app.GET("/check", service.CheckNeedCaptcha)
	app.GET("/logout", service.LogoutService{}.Logout)
	app.POST("/addtodolist", service.ToDoList{}.AddToDoList)
	app.GET("/todolist", service.ToDoList{}.GetToDoList)
	app.POST("/updatetodolist", service.ToDoList{}.UpdateTodoList)
	app.POST("/deletetodolist", service.ToDoList{}.DeleteTodoList)
	app.GET("/classroom", service.ClassRoomService{}.GetClassRoom)
	app.GET("/GetEvaluateStatus", service.EvaluateService{}.GetEvaluateInfo)
}
