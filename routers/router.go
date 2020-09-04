package routers

import (
	"go-web-frame/controller"
	"go-web-frame/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//router := CreateRouter()

}

func CreateRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Consuming())

	// v1 版本的路由配置
	apiV1 := router.Group("/v1")
	{
		// 会员接口
		v1Member := apiV1.Group("member")
		v1Member.GET("/info", controller.NewMember().View)               // 会员详情
		v1Member.POST("/login", controller.NewMember().Login)            // 登录
		v1Member.POST("/register", controller.NewMember().Register)      // 注册
		v1Member.GET("/group", controller.NewMember().GroupView)         // 获取会员的团队信息
		v1Member.PUT("/password", controller.NewMember().UpdatePassword) // 会员详情

	}

	// 系统相关的配置
	system := router.Group("/system")
	{
		system.GET("/health", controller.NewSystem().Health) // 监控检查
	}

	return router
}