package api

import "github.com/gin-gonic/gin"

func RUNENGINE() {
	 r :=gin.Default()

	 //登录 login放包内 不大写
	 r.POST("/login",login)
	 //注册 同样
	 r.POST("/newuser",newuser)
	 //找回密码
	 r.POST("/findpassword",findpassword)

	 //设置路由组 同url前缀放一组
	 userGroup :=r.Group("/user")
	 {
		 userGroup.Use(auth)//鉴权
		 userGroup.POST("/newpassword",newpassword)//修改密码
	}

	//留言相关路由组
	postGroup :=r.Group("/post")
	{
		userGroup.Use(auth)//鉴权
		postGroup.POST("/add",addPost)//发留言
		postGroup.POST("/mod/:post_id",modifypost)//修改留言 给出相应ID就可以修改

		postGroup.GET("/",lookpost)//看留言
		postGroup.GET("/:post_id",postDetail) //查看一条留言详细信息和其下属评论
	}

	//评论路由组
	commentGroup :=r.Group("/comment")
	{
		userGroup.Use(auth)//得用userGroup
		commentGroup.POST("/",addComment)  //发评论
		commentGroup.POST("/:comment_id",modifycomment)//改评论
		commentGroup.DELETE("/:comment_id",DropComment) //删评论

		commentGroup.GET("/:comment_id",ComDetail)//查看评论和对它的回复
	}
	//启动服务器
	r.Run("127.0.0.1:9090")
}
