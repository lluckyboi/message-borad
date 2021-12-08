package dao

import "message-board/model"

// SelectUserByUsername 登录信息查询
func SelectUserByUsername(username string)(model.User,error) {
	user :=model.User{}
	sqlstr := "select id,password from user where username=?"
	//单行查询
	errs :=Db.QueryRow(sqlstr,username)

	//错误处理
	if errs.Err()!=nil{
		return user, errs.Err()
	}
	err :=errs.Scan(&user.Id,&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SelectQuesAndPasswordByUsername 通过密保找回密码模块
func SelectQuesAndPasswordByUsername(username string)(model.User,error){
	user :=model.User{}
	sqlstr :="select ans1,ans2,password from user where username=?"
	errs :=Db.QueryRow(sqlstr,username)
	//Err provides a way for wrapping packages to check for query errors without calling Scan.
	if errs.Err() != nil {
		return user, errs.Err()
	}
	//扫进user
	err := errs.Scan(&user.Ans1,&user.Ans2,&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SelectUsernameByUsername 用户名查重
func SelectUsernameByUsername(username string)(model.User,error){
	user :=model.User{}
	sqlstr :="select username from user where username=?"
	errs :=Db.QueryRow(sqlstr,username)
	if errs.Err() != nil {
		return user, errs.Err()
	}
	err := errs.Scan(&user.Ans1,&user.Ans2,&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

// NewUser 注册账户
func NewUser(user model.User)(error){
	sqlstr:="insert into user(username,password,ques1,ques2,ans1,ans2)values(?,?,?,?,?,?);"
	_, errs :=Db.Exec(sqlstr,user.Username,user.Password,user.Ques1,user.Ques2,user.Ans1,user.Ans2)
	if errs!=nil{
		return errs
	}
	return nil
}

// Newpassword 修改密码
func Newpassword(username,newpassword string)(error){
	user:=model.User{}
	sqlstr :="select password from user where username=?"
	errs :=Db.QueryRow(sqlstr,username)
	if errs.Err() != nil {
		return errs.Err()
	}
	err := errs.Scan(&user.Password)
	if err != nil {
		return err
	}

	sqlstr ="UPDATE user SET password=? WHERE username =?"
	_,err=Db.Exec(sqlstr,newpassword,username)
	if err!=nil{
		return err
	}
	return nil
}
