package controller

import (
	"go-web-frame/filter"
	"go-web-frame/pkg/api"
	"go-web-frame/pkg/api/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Member struct {
}

func NewMember() *Member {
	return &Member{}
}

func (m *Member) View(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

// 用户登录API
func (m *Member) Login(c *gin.Context) {
	if err := filter.Login(c); err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// 用户注册API
func (m *Member) Register(c *gin.Context) {
	if err := filter.Register(c); err != nil {
		api.SetResponse(c, http.StatusOK, code.ERROR_REQUEST_ERROR, "")
		return
	}

	api.SetResponse(c, http.StatusOK, code.Success, "")
}

// 用户团队API
func (m *Member) GroupView(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

func (m *Member) UpdatePassword(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
