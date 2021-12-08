package api

import (
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"time"
)

func addPost(c *gin.Context){
	username :=c.PostForm("username")
	txt := c.PostForm("txt")

	//check
	info,booll:=tool.LengthCheck(username)
	if booll!=true{
		c.String(200,info)
		return
	}
	info1,bool1:=tool.LengthCheck(txt)
	if bool1!=true{
		c.String(200,info1)
		return
	}

	//每次发表需要的基本信息
	post :=model.Post{
		Txt:        txt,
		Username:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	//c.String(200,"%s",post.Txt) 测试点
	err :=service.AddPost(post)
	if err!=nil {
		c.String(2000,"append err",err)
		return
	}
	tool.RespSuccessful(c,"发表")
}

func modifypost(c *gin.Context){
	//读id
	id :=c.Param("post_id")
	ntxt :=c.PostForm("ntxt")

	err :=service.Modifypost(id,ntxt)
	if err!=nil{
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c,"修改")
}

func lookpost(c *gin.Context){
	post,err :=service.AllPosts()
	if err!=nil{
		c.String(200,"get pots err",err)
		return
	}
	tool.RespSuccessful(c,"获取留言")
	c.JSON(200,gin.H{
		  "留言":post,
	})
}

func postDetail(c *gin.Context){

	id:=c.Param("post_id")
	//先取post
	post,err :=service.PostOneGet(id)
	if err!=nil{
		c.String(200,"Get post detail err",err)
		return
	}
	c.JSON(200,gin.H{
		"留言":post,
	})

	//评论
	coms,err :=service.GetCommentOfPost(id)
	if err!=nil{
		tool.RespInternalError(c)
		return
	}
	c.String(200,"\n")
	c.JSON(200,gin.H{
		"评论":coms,
	})
}