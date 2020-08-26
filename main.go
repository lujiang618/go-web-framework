package main

import (
	"go-web-frame/pkg/config"
	"go-web-frame/pkg/db"
	"go-web-frame/pkg/web"
)

func init() {

	config.InitConfig("./config/local.yaml") // 加载配置文件
	db.InitDb()                              // 初始化DB(MySQL、Redis、Mongo)连接

}

// 入口文件，main函数
func main() {
	// gin.SetMode(setting.ServerSetting.RunMode)

	web.StartWebServer()
	destroy()
}

func destroy() {
	db.Destroy() // 关闭数据库连接
}
