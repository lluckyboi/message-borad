package api

import (
	"github.com/gin-gonic/gin"
	"message-board/tool"
)

//鉴权中间件
func auth(jq *gin.Context){
	value,err :=jq.Cookie("gin_cookie")
	//错误处理
	if err!=nil{
		tool.RespErrorWithDate(jq,"请登陆后进行操作")
		jq.Abort()
	}else{
		//获取的cookie写入上下文
		jq.Set("gin_cookie",value)
		//挂起来执行剩下进程
		jq.Next()
	}
}

