package main

import (
	"go-web-frame/pkg/config"
	"go-web-frame/pkg/db"
	"go-web-frame/pkg/web"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {

	// 创建日志文件并设置为 gin.DefaultWriter
	f, _ := os.Create("runtime/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	if gin.Mode() != gin.ReleaseMode {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

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
