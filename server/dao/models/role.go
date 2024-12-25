package models

import "github.com/bwmarrin/snowflake"

type Role struct {
	Base
	TenantID    snowflake.ID  `gorm:"not null;default:0;comment:租户ID"`
	Name        string        `gorm:"size:128;not null;default:'';comment:名称"`
	Code        string        `gorm:"size:64;not null;default:'';comment:代码"`
	Sort        int64         `gorm:"type:int(11);not null;default:0;comment:排序"`
	Status      Status        `gorm:"type:int(11);not null;default:0;comment:状态"`
	Remark      string        `gorm:"size:128;not null;default:'';comment:备注"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
	Users       []*User       `gorm:"many2many:user_roles"`
}
