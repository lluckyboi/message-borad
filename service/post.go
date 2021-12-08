package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddPost(post model.Post)(error){
	err :=dao.InsertPost(post)
	if err!=nil{
		return err
	}
	return nil
}

func Modifypost(id,txt string)(error){
	//先check 有没有
	err:=dao.SelectIdById(id)
	if err!=nil{
		return err
	}

	err2 :=dao.UpdatePost(id,txt)
	if err2!=nil{
		return err
	}
	return nil
}

func AllPosts()([]model.OtPost,error){
	post,err:=dao.GetAllPost()
	if err!=nil{
		return post,err
	}
	return post,nil
}

func PostOneGet(id string)(model.OtPost,error){
	post,err :=dao.OnePostGet(id)
	if err!=nil{
		return post,err
	}
	return post,nil
}