package types

import (
	"gin-vben-admin/dao/models"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
)

type PermissionItem struct {
	ID       snowflake.ID      `json:"id"`
	Name     string            `json:"name"`
	Code     string            `json:"code"`
	Sort     int64             `json:"sort"`
	Desc     string            `json:"desc"`
	ParentId snowflake.ID      `json:"parentId"`
	Children []*PermissionItem `json:"children,omitempty"`
}

type PermissionsTreeData struct {
	Items []*PermissionItem `json:"items"`
}

func RespPermissionsTree(permissions []*models.Permission) e.Response {
	data := PermissionsTreeData{
		Items: []*PermissionItem{},
	}
	if permissions == nil {
		return e.RespOK.WithData(data)
	}
	for _, permission := range permissions {
		data.Items = append(data.Items, buildPermission(permission))
	}
	return e.RespOK.WithData(data)
}

func buildPermission(p *models.Permission) *PermissionItem {
	item := &PermissionItem{
		ID:       p.ID,
		Name:     p.Name,
		Code:     p.Code,
		Sort:     p.Sort,
		ParentId: p.ParentID,
		Desc:     p.Description,
		Children: []*PermissionItem{},
	}
	if p.Permissions != nil {
		for _, permission := range p.Permissions {
			item.Children = append(item.Children, buildPermission(permission))
		}
	}
	return item
}
