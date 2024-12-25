package types

import (
	"gin-vben-admin/dao/models"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"time"
)

type RoleItem struct {
	ID          snowflake.ID  `json:"id"`
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	Sort        int64         `json:"sort"`
	Status      models.Status `json:"status"`
	Remark      string        `json:"remark"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Permissions []string      `json:"permissions"`
}

type RolesData struct {
	Total int64       `json:"total"`
	Items []*RoleItem `json:"items"`
}

func RespRoles(roles []*models.Role, count int64) e.Response {
	data := RolesData{
		Total: count,
		Items: []*RoleItem{},
	}
	for _, role := range roles {
		permissions := []string{}
		if role.Permissions != nil {
			for _, permission := range role.Permissions {
				permissions = append(permissions, permission.Code)
			}
		}
		data.Items = append(data.Items, &RoleItem{
			ID:          role.ID,
			Name:        role.Name,
			Code:        role.Code,
			Sort:        role.Sort,
			Status:      role.Status,
			Remark:      role.Remark,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
			Permissions: permissions,
		})
	}
	return e.RespOK.WithData(data)
}
