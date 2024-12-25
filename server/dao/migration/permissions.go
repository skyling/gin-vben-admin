package migration

import (
	"errors"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/global"
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var typeFunc = "admin"

// 权限最多三层
var permissions = []*models.Permission{
	// 系统管理
	{
		Type: typeFunc, Name: "系统管理", Code: "system-menu", Description: "系统管理菜单", Sort: 1,
		Permissions: []*models.Permission{
			{
				Type: typeFunc, Name: "角色管理", Code: "role", Description: "角色管理", Sort: 1,
				Permissions: []*models.Permission{
					{Type: typeFunc, Name: "创建角色", Code: "role-create"},
					{Type: typeFunc, Name: "修改角色", Code: "role-update"},
					{Type: typeFunc, Name: "删除角色", Code: "role-delete"},
				},
			},
			{
				Type: typeFunc, Name: "部门管理", Code: "dept", Description: "部门管理", Sort: 2,
				Permissions: []*models.Permission{
					{Type: typeFunc, Name: "创建部门", Code: "dept-create"},
					{Type: typeFunc, Name: "修改部门", Code: "dept-update"},
					{Type: typeFunc, Name: "删除部门", Code: "dept-delete"},
				},
			},
			{Type: typeFunc, Name: "账号管理", Code: "user", Description: "账号管理权限", Sort: 3,
				Permissions: []*models.Permission{
					{Type: typeFunc, Name: "创建账号", Code: "user-create"},
					{Type: typeFunc, Name: "修改账号", Code: "user-update"},
					{Type: typeFunc, Name: "删除账号", Code: "user-delete"},
				},
			},
		},
	},
	// 系统管理END
	{
		Type: typeFunc, Name: "产品管理", Code: "product-menu", Description: "产品管理菜单", Sort: 2,
		Permissions: []*models.Permission{
			{Type: typeFunc, Name: "产品分类", Code: "category", Description: "产品分类管理", Sort: 1,
				Permissions: []*models.Permission{
					{Type: typeFunc, Name: "创建分类", Code: "category-create"},
					{Type: typeFunc, Name: "修改分类", Code: "category-update"},
					{Type: typeFunc, Name: "删除分类", Code: "category-delete"},
				},
			},
		},
	},
}

func addDefaultPermissions() {
	db := global.DB
	logrus.Infof("========add default permissions")
	updatePermissions(db, 0, permissions)
	addAllPermissionToAdmin(db)
}

func updatePermissions(db *gorm.DB, parentId int64, pers []*models.Permission) {
	for k, per := range pers {
		p := models.Permission{}
		if per.Sort == 0 {
			per.Sort = int64(k + 1)
		}
		err := db.Where("type=? and code=?", per.Type, per.Code).First(&p).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			per.ParentID = snowflake.ID(parentId)
			db.Omit(clause.Associations).Create(per)
			db.Where("type=? and code=?", per.Type, per.Code).First(&p)
		} else {
			db.Where("id=?", p.ID).UpdateColumns(map[string]interface{}{
				"name":        per.Name,
				"description": per.Description,
			})
		}
		if per.Permissions != nil {
			updatePermissions(db, per.ID.Int64(), per.Permissions)
		}
	}
}

func addAllPermissionToAdmin(db *gorm.DB) error {
	pers := []*models.Permission{}
	err := db.Where("deleted_at is null").Find(&pers).Error
	if err != nil {
		return err
	}
	rps := []*models.RolePermission{}
	for _, per := range pers {
		rps = append(rps, &models.RolePermission{
			PermissionID: per.ID,
			RoleID:       1,
		})
	}
	return db.CreateInBatches(rps, 50).Error
}
