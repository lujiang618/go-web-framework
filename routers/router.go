package routers

import (
	"go-web-frame/pkg/api/middleware"

	_ "go-web-frame/docs/swagger"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	InitRouter()
	CreateApiRouter()
}

func InitRouter() {
	Router = gin.New()

	if gin.Mode() != gin.ReleaseMode {
		prefix := "/pprof"
		pprof.Register(Router, prefix)
	}

	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(middleware.Consuming())

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
