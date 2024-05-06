package model

import (
	"github.com/gin-gonic/gin"
)

func Error(context *gin.Context) {
	err := recover()
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "内部发生错误!",
		})
		context.Abort()
		return
	}
}
