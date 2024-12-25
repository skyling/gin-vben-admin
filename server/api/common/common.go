package common

import (
	"gin-vben-admin/global"
	"gin-vben-admin/pkg/avatar"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GenAvatarPicture(c *gin.Context) {
	name := c.Param("key")
	size := c.Query("size")
	if size == "" {
		size = "200"
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		sz = 200
	}

	data, _ := global.Avatar.DrawToBytes(avatar.BGC, name, sz)
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Header().Set("Cache-Control", "max-age=86400")
	c.Writer.Write(data)
}
