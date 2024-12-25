package admin

import (
	"gin-vben-admin/api/admin/types"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"github.com/gin-gonic/gin"
)

type PermissionsReq struct {
}

func (req *PermissionsReq) Run(c *gin.Context) e.Response {
	uid := auth.CurrentUserID(c)
	permissions, _ := repo.PermissionSrv.UserTree(uid)
	return types.RespPermissionsTree(permissions)
}
