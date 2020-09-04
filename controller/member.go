package controller

import (
	"go-web-frame/pkg/api"
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

func (m *Member) Login(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

func (m *Member) Register(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

func (m *Member) GroupView(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}

func (m *Member) UpdatePassword(c *gin.Context) {
	api.SetResponse(c, http.StatusOK, 1, "")
}
