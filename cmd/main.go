package main

import (
	"message-board/api"
	"message-board/dao"
)

func main() {
		//先连接数据库
		dao.RUNDB()
		//启动引擎
		api.RUNENGINE()
}
