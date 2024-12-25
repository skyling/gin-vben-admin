package migration

import (
	"errors"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/global"
	"gin-vben-admin/pkg/constant/setting"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

func addDefaultSettings() {
	defaultSettings := []models.Setting{
		// 版本
		{Name: NameDbVersion + RequiredDBVersion, Value: `installed`, Type: setting.TypeVersion},
		{Name: NamePermissionVersion + PermissionVersion, Value: `installed`, Type: setting.TypeVersion},
	}
	db := global.DB
	var err error
	for _, value := range defaultSettings {
		var s models.Setting
		err = db.Where("type=? and name=?", value.Type, value.Name).First(&s).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.Setting{Name: value.Name, Type: value.Type, Value: value.Value})
		}
	}
}

func addDefaultRole() {
	role := models.Role{
		Base: models.Base{
			ID: snowflake.ID(1),
		},
		Name:   "超级管理员",
		Code:   "admin",
		Sort:   0,
		Status: models.StatusOn,
	}
	db := global.DB
	roleT := models.Role{}
	err := db.Where("id=?", role.ID).First(&roleT).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&role)
	}

	ur := models.UserRole{
		UserID: 1,
		RoleID: 1,
	}
	urt := models.UserRole{}
	err = db.Where("role_id=? and user_id=?", ur.RoleID, ur.UserID).First(&urt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&ur)
	}
}

func addDefaultUser() {
	admin := models.User{
		Base: models.Base{
			ID: snowflake.ID(1),
		},
		Type:     models.UserTypeAdmin,
		Name:     "管理员",
		Username: "admin",
		Status:   models.UserStatusON,
		Code:     "ADMIN",
	}
	admin.SetPassword("password")

	users := []models.User{
		admin,
	}
	db := global.DB
	for _, user := range users {
		u := models.User{}
		err := db.Where("username=?", user.Username).First(&u).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&user)
		}
	}
}
