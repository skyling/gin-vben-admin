package models

import (
	"errors"
	"gin-vben-admin/pkg/conf"
	"github.com/bwmarrin/snowflake"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"gorm.io/datatypes"
	"strings"
)

const (
	UserTypeAdmin      = "admin"       // 管理员
	UserTypeTenant     = "tenant"      // 租户
	UserTypeTenantUser = "tenant-user" // 用户

	UserStatusON  = 1 // 启用
	UserStatusOFF = 2 // 禁用
)

type UserExtend struct {
}

type User struct {
	Base
	Type     string                          `gorm:"size:32;not null;default:''"`
	TenantID snowflake.ID                    `gorm:"not null;default:0;comment:租户ID"`
	Avatar   string                          `gorm:"size:512;not null;default:'';comment:'头像'"`
	Code     string                          `gorm:"size:32;not null;default:''"`  // 用户编码
	Name     string                          `gorm:"size:128;not null;default:''"` // 名称
	Username string                          `gorm:"size:128;not null;default:''"` // 登录用户名
	Password string                          `gorm:"size:128;not null;default:''"` // 密码
	Remark   string                          `gorm:"size:512;not null;default:''"` // 用户备注
	DeptID   snowflake.ID                    `gorm:"not null;default:0;comment:部门ID"`
	Extend   datatypes.JSONType[*UserExtend] `gorm:"type:json"`
	Status   Status                          `gorm:"not null;default:0;"` // 状态 1:启用 2:禁用
	Dept     *Dept                           `gorm:"foreignKey:dept_id"`
	Roles    []*Role                         `gorm:"many2many:user_roles"`
}

type UserRole struct {
	UserID snowflake.ID `gorm:"primaryKey;not null;default:0;comment:用户ID"`
	RoleID snowflake.ID `gorm:"primaryKey;not null;default:0;comment:角色ID"`
}

func (u *User) SetExtend(extend *UserExtend) {
	u.Extend.Data = extend
}

func (u *User) GetExtend() *UserExtend {
	return u.Extend.Data
}

// SetPassword 根据给定明文设定 User 的 Password 字段
func (u *User) SetPassword(password string) error {
	//生成16位 Salt
	salt := random.RandString(16)
	//计算 Salt 和密码组合的SHA1摘要
	bs := cryptor.Sha1(password + salt)
	//存储 Salt 值和摘要， ":"分割
	u.Password = salt + ":" + bs
	return nil
}

// CheckPassword 根据明文校验密码
func (u *User) CheckPassword(password string) (bool, error) {
	// 根据存储密码拆分为 Salt 和 Digest
	passwordStore := strings.Split(u.Password, ":")
	if len(passwordStore) != 2 {
		return false, errors.New("Unknown password type")
	}
	//计算 Salt 和密码组合的SHA1摘要
	bs := cryptor.Sha1(password + passwordStore[0])
	return bs == passwordStore[1], nil
}

func (u *User) AvatarURL() string {
	return conf.BuildUrl("mapi/avatar/" + u.Name)
}
