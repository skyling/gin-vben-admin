package migration

import (
	"context"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/global"
	"github.com/sirupsen/logrus"
)

// 是否需要迁移
func needMigration() bool {
	var s models.Setting
	return global.DB.Where("name = ?", NameDbVersion+RequiredDBVersion).First(&s).Error != nil
}

func needUpdatePermission() bool {
	var s models.Setting
	return global.DB.Where("name = ?", NamePermissionVersion+PermissionVersion).First(&s).Error != nil
}

// Run 执行数据迁移
func Run() {
	// 确认是否需要执行迁移
	// 清除所有缓存

	// 权限节点更新
	if needUpdatePermission() {
		defer func() {
			addDefaultSettings()
			addDefaultPermissions()
			repo.CasbinSrv.UpdatePolicy(1)
			repo.CasbinSrv.AddRolesForUser(&models.User{Base: models.Base{ID: 1}})
		}()
	}

	if !needMigration() {
		logrus.Infof("数据库版本匹配，跳过数据库迁移")
		return
	}
	global.RedisCache.Clear(context.Background())

	logrus.Infof("开始进行数据库初始化...")

	// 自动迁移模式
	if global.Conf.System.DbType == "mysql" {
		global.DB.Set("gorm:table_options", " ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci").
			AutoMigrate(models.ModelsList...)
	} else {
		global.DB.AutoMigrate(models.ModelsList...)
	}

	// 向设置数据表添加初始设置
	addDefaultSettings()
	addDefaultUser()
	addDefaultRole()
	logrus.Infof("数据库初始化结束")
}
