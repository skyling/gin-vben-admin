package admin

import (
	"gin-vben-admin/api"
	"gin-vben-admin/api/admin/types"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/global"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type ChangePasswordReq struct {
	api.BaseReq
	OldPassword string `form:"oldPassword" json:"oldPassword" binding:"required,min=6,max=64" field:"密码" example:"123456"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required,min=6,max=64" field:"密码" example:"123456"`
}

func (req *ChangePasswordReq) Run(c *gin.Context) e.Response {
	u := auth.CurrentUser(c)
	if global.Conf.System.Env == "dev" {
		return e.ErrForbidden.WithMsg("开发环境禁止修改密码").Resp()
	}
	params := &repo.ChangePasswordParams{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		UserID:      u.ID,
	}
	err := repo.UserSrv.ChangePassword(params)
	if err != nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	return e.RespOK
}

type LoginLogReq struct {
	api.PageReq
	CreatedAtStart *time.Time `form:"createdAtStart" json:"createdAtStart" binding:"omitempty" field:"" example:""` // 开始时间
	CreatedAtEnd   *time.Time `form:"createdAtEnd" json:"createdAtEnd" binding:"omitempty" field:"" example:""`     // 结束时间
	IP             string     `form:"ip" json:"ip" binding:"omitempty" field:"" example:""`                         // ip 地址
}

func (req *LoginLogReq) Run(c *gin.Context) e.Response {
	params := &repo.LoginLogPageParams{
		IP:             req.IP,
		CreatedAtStart: req.CreatedAtStart,
		CreatedAtEnd:   req.CreatedAtEnd,
	}
	params.Offset = req.GetOffset()
	params.Limit = req.GetLimit()
	params.BaseParams = api.GetBaseParams(c)
	logrus.Info(params)
	logs, total, _ := repo.UserSrv.LoginLogs(params)
	return types.RespLoginLogs(logs, total)
}
