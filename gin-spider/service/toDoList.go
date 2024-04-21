package service

import (
	"fmt"
	"gin-spider/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ToDoList struct {
}

func (toDoList ToDoList) AddToDoList(context *gin.Context) {
	title := context.PostForm("title")
	state, err := strconv.Atoi(context.DefaultPostForm("status", "0"))
	id := context.PostForm("id")
	stuId := context.PostForm("stuId")
	school := context.PostForm("school")
	var todolist *model.Todolist
	if err != nil || stuId == "" || id == "" || title == "" || school == "" {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	todolist = &model.Todolist{
		Id:     id,
		Title:  title,
		Status: int32(state),
		StuId:  stuId,
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
	stuId := context.Query("stuId")
	var todolist []model.Todolist
	fmt.Println(model.DB)
	result := model.DB.Table("todolist").Find(&todolist).Where("StuId = ?", stuId).Order("id desc")
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
