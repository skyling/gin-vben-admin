package middleware

import (
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/constant"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CheckRight 检查权限
func CheckRight(rights ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := auth.CurrentUser(c)
		if !repo.CasbinSrv.CheckUserRight(u.ID, "", rights...) {
			c.JSON(200, e.ErrForbidden.Resp())
			c.Abort()
			return
		}
		c.Next()
	}
}

// AuthRequired 登录认证
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uid, ok := c.Get(constant.UserIDKey); !ok {
			c.JSON(200, e.ErrUnauthorized.Resp())
			c.Abort()
			return
		} else if uid == nil || uid.(snowflake.ID) == 0 {
			c.JSON(200, e.ErrUnauthorized.Resp())
			c.Abort()
			return
		}
		c.Next()
	}
}

// CurrentUser 当前用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ParseUser(c)
		if err != nil {
			logrus.Infof("current user %v", err)
		}
		c.Next()
	}
}
