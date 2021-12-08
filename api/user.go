package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
)

func newuser(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	ques1	:=c.PostForm("ques1")
	ques2	:=c.PostForm("ques2")
	ans1	:=c.PostForm("ans1")
	ans2	:=c.PostForm("ans2")

	user :=model.User{
		Username :username,
		Password :password,
		Ques1    :ques1,
		Ans1 	 :ans1,
		Ques2	 :ques2,
		Ans2	 :ans2,
	}
	isok,err:=service.IsUsernameRepeat(username)
	if err!=nil{
		fmt.Println("judge username repeat err: ", err)
		tool.RespInternalError(c)
		return
	}
	//重复了
	if isok!=true{
		c.String(200,"用户名重复了！")
	}

	err =service.NewUser(user)
	if err!=nil{
		fmt.Println("register err: ", err)
		tool.RespInternalError(c)
		return
	}
}

func login(c *gin.Context) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	//check账号密码
	isok, err :=service.IsPasswordCorrect(username,password)

	//错误处理
	if err !=nil{
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(c)
		return
	}
	if isok != true {
		c.String(200,"密码错误")
		return
	}


	//登录成功的反馈
	tool.RespSuccessful(c,username+"登录")

	//设置cookie
	c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
}

func findpassword(c *gin.Context){
	//获取相关信息
	username:=c.PostForm("username")
	ans1:=c.PostForm("ans1")
	ans2:=c.PostForm("ans2")

	isok,err,password:=service.IsPasswordQuesOk(username,ans1,ans2)
	if err!=nil{
		fmt.Println("judge answers correct err: ", err)
		tool.RespInternalError(c)
		return
	}
	if isok!=true{
		c.String(200,"密保问题答案错误！")
		return
	}

	tool.RespSuccessful(c,username+"找回密码")
	c.String(200,"你的密码是%s",password)
}

//改密码
func newpassword(c *gin.Context){
	BefPassword := c.PostForm("BefPassword")
	NewPassword := c.PostForm("NewPassword")
	username 	:= c.PostForm("username")


	isok, err := service.IsPasswordCorrect(username,BefPassword)
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		tool.RespInternalError(c)
		return
	}
	if isok!=true{
		c.String(200,"密码错误！")
		return
	}

	isok,err =service.NewPassword(username,NewPassword)
	if err!=nil{
		fmt.Println("updata password  err: ", err)
		tool.RespInternalError(c)
		return
	}
	if isok==false {
		c.String(200, "修改失败！")
		return
	}
	c.String(200,"修改成功！")
}