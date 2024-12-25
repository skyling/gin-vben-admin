package admin

import (
	"gin-vben-admin/api/admin/types"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"github.com/gin-gonic/gin"
)

type UserTypeOptionsReq struct {
}

func (req *UserTypeOptionsReq) Run(c *gin.Context) e.Response {
	u := auth.CurrentUser(c)
	options := types.OptionsData{
		Items: []*types.OptionsItem{
			{Label: "用户", Value: models.UserTypeTenantUser},
		},
	}
	if u.Type == models.UserTypeAdmin {
		options = types.OptionsData{
			Items: []*types.OptionsItem{
				{Label: "管理员", Value: models.UserTypeAdmin},
				{Label: "租户", Value: models.UserTypeTenant},
			},
		}
	}

	return types.RespOptionsData(&options)
}
