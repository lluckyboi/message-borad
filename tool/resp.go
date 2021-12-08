package tool

import (
	"github.com/gin-gonic/gin"
)

// RespErrorWithDate 数据错误
func RespErrorWithDate(c *gin.Context, date interface{}) {
	c.JSON(200, gin.H{
		"info": date,
	})
}

// RespSuccessful 成功反馈
func RespSuccessful(c *gin.Context,info string){
	c.String(200,info+"成功!")
}

// RespInternalError 服务器错误
func RespInternalError(c *gin.Context){
	c.String(500,"服务器错误！")
}
