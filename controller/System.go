package controller

import (
	"go-web-frame/pkg/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type System struct {
}

func NewSystem() *System {
	return &System{}
}

func (s *System) Health(c *gin.Context) {

	api.SetResponse(c, http.StatusOK, 1, "")
}
