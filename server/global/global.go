package global

import (
	"gin-vben-admin/pkg/avatar"
	"gin-vben-admin/pkg/conf"
	"github.com/eko/gocache/lib/v4/marshaler"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Conf  *conf.Conf
	Viper *viper.Viper

	BigCache   *marshaler.Marshaler
	RedisCache *marshaler.Marshaler
	TokenCache *marshaler.Marshaler
	Avatar     *avatar.InitialsAvatar
)
