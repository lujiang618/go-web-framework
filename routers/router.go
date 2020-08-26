package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//router := CreateRouter()

}

func CreateRouter() *gin.Engine {
	router := gin.New()

	// v1 版本的路由配置
	apiV1 := router.Group("/v1")
	apiV1.Use(gin.Logger())
	apiV1.Use(gin.Recovery())
	{

	}

	// 系统相关的配置
	system := router.Group("/system")

	system.Use(gin.Logger())
	{
		system.GET("/health") // 监控检查
	}

	return router
}
