package main

import (
	"go-web-frame/pkg/config"
	"go-web-frame/pkg/db"
)

func init() {
	// 加载配置文件
	config.InitConfig("./config/local.yaml")

	// 初始化DB(MySQL、Redis、Mongo)连接
	db.InitDb()

}

// 入口文件，main函数
func main() {
	// gin.SetMode(setting.ServerSetting.RunMode)
	//gin.Default()

	// 初始化数据库配置

	// 加载路由

	//

}
