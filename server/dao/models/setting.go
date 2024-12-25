package models

import "github.com/bwmarrin/snowflake"

// Setting 配置项
type Setting struct {
	Base
	TenantID snowflake.ID `gorm:"not null;default:0;comment:用户ID"` // 0 未系统配置
	Type     string       `gorm:"size:32;not null;comment:类型"`
	Name     string       `gorm:"index:;type:varchar(100);not null;comment:配置key"`
	Value    string       `gorm:"type:text;comment:配置值"`
}
