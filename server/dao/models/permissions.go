package models

import "github.com/bwmarrin/snowflake"

type Permission struct {
	Base
	Type        string        `gorm:"size:32;not null;default:'';comment:类型"`
	ParentID    snowflake.ID  `gorm:"default:0;comment:父级ID"`
	Sort        int64         `gorm:"type:int(11);not null;default:0;comment:排序"`
	Name        string        `gorm:"size:64;not null;default:'';comment:名称"`
	Code        string        `gorm:"size:64;not null;default:'';comment:代码"`
	Description string        `gorm:"size:128;not null;default:'';comment:描述"`
	Roles       []*Role       `gorm:"many2many:role_permissions"`
	Permissions []*Permission `gorm:"foreignKey:parent_id"`
}

type RolePermission struct {
	RoleID       snowflake.ID `gorm:"primaryKey;not null;default:0;comment:角色ID"`
	PermissionID snowflake.ID `gorm:"primaryKey;not null;default:0;comment:权限ID"`
}
