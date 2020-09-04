package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 验证登录的请求参数
func Login(c *gin.Context) error {
	params := &LoginParams{}

	if err := c.ShouldBind(params); err != nil {
		logrus.Error(err)

		return err
	}

	// 调用service对应的方法

	return nil
}

// 验证注册信息
func Register(c *gin.Context) error {
	params := &RegisterParams{}

	if err := c.ShouldBind(params); err != nil {
		logrus.Error(err)

		return err
	}

	// 调用service对应的方法

	return nil
}
