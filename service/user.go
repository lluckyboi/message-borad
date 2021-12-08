package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	if user.Password != password {
		return false, nil
	}

	return true, nil
}

func IsPasswordQuesOk(username,ans1,ans2 string)(bool ,error,string){
	user,err :=dao.SelectQuesAndPasswordByUsername(username)
	if err!=nil{
		return false,err,""
	}
	if ans1==user.Ans1&&ans2==user.Ans2{
		return true,nil,user.Password
	}
	return false,nil,""
}

func IsUsernameRepeat(username string)(bool,error){
	user,err :=dao.SelectUsernameByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.Username==username{
		return false,nil
	}
	return true, nil
}

func NewUser(user model.User) error {
	err := dao.NewUser(user)
	return err
}

func NewPassword(username,newpassword string)(bool,error){
	err :=dao.Newpassword(username,newpassword)
	if err!=nil{
		return  false,err
	}
	return true,nil
}