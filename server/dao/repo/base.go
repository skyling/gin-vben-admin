package repo

import (
	"gin-vben-admin/dao/models"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type BaseParams struct {
	TenantID snowflake.ID // 租户ID
	UserID   snowflake.ID // 用户ID
	UserType string       // 1:后台用户 2:客户端用户
}

type PageParams struct {
	BaseParams
	Offset int // 分页偏移量
	Limit  int // 每页数量
}

func ScopeTenant(params BaseParams) func(db gen.Dao) gen.Dao {
	return func(db gen.Dao) gen.Dao {
		if params.UserType != models.UserTypeAdmin && params.TenantID != 0 {
			db.Where(field.Attrs(map[string]interface{}{"tenant_id": params.TenantID.String()}))
		} else {
			db.Where(field.Attrs(map[string]interface{}{"type": []string{models.UserTypeTenant, models.UserTypeAdmin}}))
		}
		return db
	}
}
