package api

import (
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"time"
)

func addComment(c *gin.Context) {
	txt := c.PostForm("txt")
	postid := c.PostForm("postid")
	username := c.PostForm("username")
	iscomcom := c.PostForm("IsComCom")
	comid 	 :=c.PostForm("comcomid")
	//入参check
	info, booll := tool.LengthCheck(txt)
	if booll != true {
		c.String(200, info)
		return
	}
	info1, bool1 := tool.LengthCheck(postid)
	if bool1 != true {
		c.String(200, info1)
		return
	}
	info2, bool2 := tool.LengthCheck(username)
	if bool2 != true {
		c.String(200, info2)
		return
	}

	if iscomcom == "yes" {
		com := model.Comment{
			ComComId:    comid,
			Txt:         txt,
			Username:    username,
			IsComCom: 	 iscomcom,
			CommentTime: time.Now(),
		}
		err := service.AddComCom(com)
		if err != nil {
			c.String(200, "Add CommentCom err", err)
			return
		}
		tool.RespSuccessful(c, "评论套娃")

	} else {
		com := model.Comment{
			PostId:      postid,
			Txt:         txt,
			Username:    username,
			CommentTime: time.Now(),
		}
		err := service.AddComment(com)
		if err != nil {
			c.String(200, "Add comment err", err)
			return
		}
		tool.RespSuccessful(c, "评论")
	}
}

func DropComment(c*gin.Context){
	id := c.Param("comment_id")
	err :=service.DelectComment(id)
	if err!=nil{
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c,"删除")
}

func modifycomment(c *gin.Context){
	id :=c.Param("comment_id")
	ntxt :=c.PostForm("ntxt")

	err :=service.ModifyComment(id,ntxt)
	if err!=nil{
		tool.RespInternalError(c)
		return
	}
	tool.RespSuccessful(c,"修改")
}


func ComDetail(c *gin.Context){
	id :=c.Param("comment_id")
	com,err:=service.GetComment(id)
	if err!=nil{
		tool.RespInternalError(c)
		return
	}
	coms,err:=service.GetComComment(id)
	com.CommentTime=time.Now()

	c.JSON(200,gin.H{
		"评论":com,
	})
	c.JSON(200,gin.H{
		"评论的评论":coms,
	})
}
