package models

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

// 一对一
// foreignKey: 外键(自己表的)
// references: 引用(关联表的) 默认为 字段名称+ID

// 一对多
// foreignKey: 外键(关联表的)
// references: 关联字段(自己表的)

var node *snowflake.Node

func init() {
	snowflake.NodeBits = 3
	node, _ = snowflake.NewNode(1)
}

type Base struct {
	ID        snowflake.ID `gorm:"primaryKey"`
	CreatedAt time.Time    `gorm:"index:idx_created_at;"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at;"`
}

func (base *Base) BeforeCreate(scope *gorm.DB) error {
	if base.ID == 0 {
		base.ID = node.Generate()
	}
	return nil
}

var ModelsList = []interface{}{
	Setting{}, LoginLog{}, CasbinRule{}, Permission{}, Role{}, UserRole{}, RolePermission{}, Dept{}, CodeIndex{}, User{}, // 用户相关
}

var ModelsTenantList = []interface{}{User{}, Role{}, Dept{}}

type Status int64

func (s Status) Int64() int64 {
	return int64(s)
}

const (
	StatusOn  = Status(1)
	StatusOff = Status(2)
)
