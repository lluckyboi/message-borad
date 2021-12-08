这周压力太大，还有些新功能没做完

后续再填坑

准备校赛去了

# 入参检验

实现是tool包里写了个lengthcheck

```go
func LengthCheck(ss string)(string,bool){
	if len(ss)>100{
		err:=ss+"超过长度限制！"
		return err,false
	}
	return "",true
}
```

简单粗暴

# 套娃评论的实现

在comment表中 加了一个判断字段

```go
type Comment struct {
	Id          int
	PostId      string
	Txt         string
	Username    string
	CommentTime time.Time
	IsComCom    string//是否是评论的评论
	ComComId	string//套娃的目标
}
```

添加评论时判断即可

在API层都是统一的AddComment

只不过在server和dao层通过不同函数区分开来

这样每一个评论都可能是一个评论的评论

由于修改与删除评论我直接用的评论ID作为主键 因此可以**对任意评论有效**



## 如果要查看一个评论的所有评论

直接SELECT出相应ComCom的comm
## 匿名

感觉差不多

输出的时候替换username 只输出正文


## 关于点赞

加一个点赞表即可

放所有的点赞记录（点赞者，被赞者，被点赞的类型，点赞时间）

需要什么直接SELECT



## 关于私密留言

加一个字段isprim，标志一个留言是否私密，然后记录，初始值为空，有值就是私密，不影响正常留言方法的复用

如果私密，在数据输出时检查一下isprim是否为空，然后和cookie中的username(我是gin_cookie)对比一下

不一样该留言就不输出