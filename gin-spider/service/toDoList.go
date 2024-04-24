package service

import (
	"fmt"
	"gin-spider/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ToDoList struct {
}

func (toDoList ToDoList) AddToDoList(context *gin.Context) {
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "请先登录!",
		})
		context.Abort()
		return
	}
	stuId := session.Get("username")
	id := time.Now().Unix()
	school := context.PostForm("school")
	if school == "" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	var count int64
	db := model.DB.Table("todolist").Where("stuId = ? and school = ?", stuId, school).Count(&count)
	if db.Error != nil || count >= 2 {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
	}
	todolist := &model.Todolist{
		Id:     strconv.FormatInt(id, 10),
		Title:  "",
		StuId:  stuId.(string),
		School: school,
	}
	result := model.DB.Table("todolist").Create(todolist)
	if result.Error == nil {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加成功",
		})
	} else {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "添加失败",
		})
	}
}

func (toDoList ToDoList) GetToDoList(context *gin.Context) {
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "请先登录!",
		})
		context.Abort()
		return
	}
	stuId := session.Get("username")
	school := context.Query("school")
	var todolist []model.Todolist
	fmt.Println(model.DB)
	result := model.DB.Table("todolist").Where("stuId = ? and school = ?", stuId, school).Order("id asc").Find(&todolist)
	if result.Error != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "查询失败",
			"data": nil,
		})
	} else {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": todolist,
		})
	}
}

func (toDoList ToDoList) UpdateTodoList(context *gin.Context) {
	session := sessions.Default(context)
	if session.Get("username") == nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "请先登录!",
		})
		context.Abort()
		return
	}
	stuId := session.Get("username")
	title := context.PostForm("title")
	school := context.PostForm("school")
	result := model.DB.Model(&model.Todolist{}).Where("stuId=? and school = ?", stuId, school).Update("title", title)
	if result.Error != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "更新失败",
			"data": nil,
		})
	} else {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新成功",
			"data": nil,
		})
	}
}

func (toDoList ToDoList) DeleteTodoList(context *gin.Context) {
	school := context.PostForm("school")
	stuId := context.PostForm("stuId")
	result := model.DB.Table("todolist").Where("stuId = ? and school = ?", stuId, school).Delete(&model.Todolist{})
	if result.Error != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "删除失败",
			"data": nil,
		})
	} else {
		context.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
			"data": nil,
		})
	}
}
