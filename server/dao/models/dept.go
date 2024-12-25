package models

import "github.com/bwmarrin/snowflake"

type Dept struct {
	Base
	TenantID snowflake.ID `gorm:"not null;default:0;comment:租户ID"`
	Name     string       `gorm:"size:128;not null;default:'';comment:名称"`
	ParentID snowflake.ID `gorm:"not null;default:0;comment:父级ID"`
	Sort     int64        `gorm:"type:int(11);not null;default:0;comment:排序值"`
	Status   Status       `gorm:"type:int(11);not null;default:0;comment:状态"`
	Remark   string       `gorm:"size:128;not null;default:'';comment:备注"`
	Depts    []*Dept      `gorm:"foreignKey:parent_id"`
}
