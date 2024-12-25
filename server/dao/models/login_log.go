package models

import "github.com/bwmarrin/snowflake"

type LoginLog struct {
	Base
	Source    string       `gorm:"size:32;not null;default:'';comment:登录来源"`
	IP        string       `gorm:"size:64;not null;default:'';comment:登录的IP"`
	UserAgent string       `gorm:"size:512;not null;default:'';comment:请求接口头部信息"`
	UserID    snowflake.ID `gorm:"not null;default:0;comment:用户ID"`
}
