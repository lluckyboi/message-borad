package service

import (
	"message-board/dao"
	"message-board/model"
)

func AddComCom(com model.Comment)(error){
	err :=dao.InsertCommentCom(com)
	if err!=nil{
		return err
	}
	return nil
}

func AddComment(com model.Comment)(error){
	err :=dao.InsertComment(com)
	if err!=nil{
		return err
	}
	return nil
}

func GetCommentOfPost(postid string)([]model.OtComment,error){
	com,err :=dao.SelectComsByPostid(postid)
	if err!=nil{
		return com,err
	}
	return com,nil
}

func GetComComment(cid string)([]model.OtComment,error){
	coms,err :=dao.SelectComComsByCId(cid)
	if err!=nil{
		return coms,err
	}
	return coms,nil
}

func GetComment(id string)(model.OtComment,error){
	com,err :=dao.SelectCommentById(id)
	if err!=nil{
		return com,err
	}
	return com,nil
}
func DelectComment(id string)(error){
	err :=dao.DeleteCommentById(id)
	if err!=nil{
		return err
	}
	return nil
}

func ModifyComment(id,ntxt string)(error){
	err :=dao.UpdateCommentById(id,ntxt)
	if err!=nil{
		return err
	}
	return nil
}