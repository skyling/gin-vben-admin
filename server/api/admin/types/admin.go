package types

import (
	"gin-vben-admin/dao/models"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"time"
)

type LoginData struct {
	AccessToken string       `json:"accessToken"` // 登录token
	ID          snowflake.ID `json:"id"`          // 用户ID
	Username    string       `json:"username"`    // 用户名
}

func RespLogin(token string, u *models.User) e.Response {
	return e.RespOK.WithData(&LoginData{
		AccessToken: token,
		ID:          u.ID,
		Username:    u.Username,
	})
}

type UserDept struct {
	ID   snowflake.ID `json:"id" swaggertype:"string"`
	Name string       `json:"name"`
}

type UserRole struct {
	ID   snowflake.ID `json:"id" swaggertype:"string"`
	Name string       `json:"name"`
}

type UserItem struct {
	ID        snowflake.ID  `json:"id" swaggertype:"string"`       // 用户ID
	Type      string        `json:"type"`                          // 用户类型
	TenantID  snowflake.ID  `json:"tenantId" swaggertype:"string"` // 租户ID
	Status    models.Status `json:"status"`                        // 状态
	Code      string        `json:"code"`                          // 编码
	Avatar    string        `json:"avatar"`                        // 头像
	Name      string        `json:"name"`                          // 昵称
	Username  string        `json:"username"`                      // 邮箱
	Remark    string        `json:"remark"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdateAt  time.Time     `json:"updateAt"`
	DeptID    snowflake.ID  `json:"deptId,omitempty" swaggertype:"string"`
	Dept      *UserDept     `json:"dept,omitempty"`
	Roles     []*UserRole   `json:"roles,omitempty"`
}

type UserInfoData struct {
	UserItem
	Permissions []string `json:"permissions"`
}

func buildUserItem(u *models.User) UserItem {
	item := UserItem{
		ID:        u.ID,
		Type:      u.Type,
		TenantID:  u.TenantID,
		Status:    u.Status,
		Avatar:    u.AvatarURL(),
		Code:      u.Code,
		Name:      u.Name,
		Username:  u.Username,
		DeptID:    u.DeptID,
		CreatedAt: u.CreatedAt,
		UpdateAt:  u.UpdatedAt,
		Remark:    u.Remark,
	}
	if u.Dept != nil {
		item.Dept = &UserDept{
			ID:   u.Dept.ID,
			Name: u.Dept.Name,
		}
	}
	if u.Roles != nil {
		for _, role := range u.Roles {
			item.Roles = append(item.Roles, &UserRole{ID: role.ID, Name: role.Name})
		}
	}
	return item
}

func RespUserInfo(u *models.User, permissions []string) e.Response {
	return e.RespOK.WithData(&UserInfoData{
		UserItem:    buildUserItem(u),
		Permissions: permissions,
	})
}

func RespUser(u *models.User) e.Response {
	return e.RespOK.WithData(buildUserItem(u))
}

type UserPermissionsData struct {
	Permissions []string `json:"permissions"`
}

func RespUserPermissions(permissions []string) e.Response {
	return e.RespOK.WithData(&UserPermissionsData{
		Permissions: permissions,
	})
}

type UsersData struct {
	Total int64       `json:"total"`
	Items []*UserItem `json:"items"`
}

func RespUsers(users []*models.User, total int64) e.Response {
	data := UsersData{
		Total: total,
		Items: []*UserItem{},
	}
	for _, u := range users {
		item := buildUserItem(u)
		data.Items = append(data.Items, &item)
	}
	return e.RespOK.WithData(data)
}
