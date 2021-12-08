package tool

func LengthCheck(ss string)(string,bool){
	if len(ss)>100{
		err:=ss+"超过长度限制！"
		return err,false
	}
	return "",true
}
