package dao

import (
	"message-board/model"
)

func InsertPost(post model.Post) error {
	//中间几个都是longtext 应该不会爆长度？
	sqlstr :="insert into post(username,txt,post_time,update_time,comment_num)values(?,?,?,?,?)"
	_,err:=Db.Exec(sqlstr,post.Username,post.Txt,post.PostTime,post.UpdateTime,post.CommentNum)
	if err!=nil{
		return err
	}
	return nil
}

func SelectIdById(id string) error {
	sqlstr :="select txt from post where id=?"
	_,err :=Db.Exec(sqlstr,id)
	if err!=nil{
		return err
	}
	return nil
}

func UpdatePost(id,txt string) error {
	sqlstr :="update post set txt=? where id=?"
	_,err :=Db.Exec(sqlstr,txt,id)
	if err!=nil{
		return err
	}
	return nil
}

func GetAllPost()([]model.OtPost,error){
	sqlstr:="select id,comment_num,username,txt,post_time,update_time from post"
	rows,err:=Db.Query(sqlstr)
	//post类型切片 存多个post
	var posts []model.OtPost
	if err != nil {
		return posts,err
	}
	//延迟关闭连接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		post:=model.OtPost{}
		err := rows.Scan(&post.Id,&post.CommentNum,&post.Username,&post.Txt,&post.PostTime,&post.UpdateTime)
		if err != nil {
			return posts,err
		}
		posts = append(posts, post)
	}
	return posts,nil
}

func OnePostGet(id string)(model.OtPost,error){
	sqlstr:="select id,comment_num,username,txt,post_time,update_time from post where id=?"
	post :=model.OtPost{}
	err:=Db.QueryRow(sqlstr,id).Scan(&post.Id,&post.CommentNum,&post.Username,&post.Txt,&post.PostTime,&post.UpdateTime)
	if err!=nil{
		return post,err
	}
	return post,nil
}

func GetCommentNumById(id string) (error,int){
	var num int
	sqlstr:="select comment_num from post where id=?"
	err :=Db.QueryRow(sqlstr,id).Scan(&num)
	if err!=nil{
		return err,0
	}
	return nil,num
}

func UpdateCommentNum(num int)(error){
	sqlstr:="update post set comment_num=?"
	_,err :=Db.Exec(sqlstr,num)
	if err!=nil{
		return err
	}
	return  nil
}