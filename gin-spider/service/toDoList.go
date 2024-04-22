package service

import (
	"fmt"
	"gin-spider/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
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
	title := context.PostForm("title")
	status, err := strconv.Atoi(context.DefaultPostForm("status", "0"))
	id := context.PostForm("id")
	school := context.PostForm("school")
	if err != nil || stuId == "" || id == "" || title == "" || school == "" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	todolist := &model.Todolist{
		Id:     id,
		Title:  title,
		Status: int32(status),
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
	id := context.PostForm("id")
	status := context.PostForm("status")
	school := context.PostForm("school")
	result := model.DB.Model(&model.Todolist{}).Where("id=? and stuId=? and school = ?", id, stuId, school).Update("status", status)
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
	id := context.PostForm("id")
	school := context.PostForm("school")
	stuId := context.PostForm("stuId")
	result := model.DB.Table("todolist").Where("id = ? and stuId = ? and school = ?", id, stuId, school).Delete(&model.Todolist{})
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
