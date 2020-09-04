package api

import (
	"go-web-frame/pkg/api/response"

	"github.com/gin-gonic/gin"
)

func SetResponse(c *gin.Context, httpCode, responseCode int, data interface{}) {
	base := response.NewResponse(c)
	base.SetResponse(httpCode, responseCode, data)

	WriterAccessLog()
}

// 记录api访问日志
func WriterAccessLog() {
	return
}