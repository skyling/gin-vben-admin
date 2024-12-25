package repo

import (
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"github.com/bwmarrin/snowflake"
)

var PermissionSrv = new(permissionSrv)

type permissionSrv struct {
}

// GetPermissionsByCode 根据code获取权限列表
func (*permissionSrv) GetPermissionsByCode(code []string) ([]*models.Permission, error) {
	qp := query.Permission
	return qp.Where(qp.Code.In(code...)).Find()
}

func (*permissionSrv) UserTree(uid snowflake.ID) ([]*models.Permission, error) {
	codes, err := CasbinSrv.GetUserPrivileges(uid, "")
	if err != nil {
		return nil, err
	}
	qp := query.Permission
	qpd := qp.Where(qp.ParentID.Eq(0)).Where(qp.Code.In(codes...))
	return qpd.Order(qp.Sort).Preload(qp.Permissions.Order(qp.Sort), qp.Permissions.Permissions.Order(qp.Sort)).Find()
}

func (*permissionSrv) Tree(typ string) ([]*models.Permission, error) {
	qp := query.Permission
	qpd := qp.Where(qp.ParentID.Eq(0))
	if typ != "" {
		qpd = qpd.Where(qp.Type.Eq(typ))
	}
	return qpd.Order(qp.Sort).Preload(qp.Permissions.Order(qp.Sort), qp.Permissions.Permissions.Order(qp.Sort)).Find()
}

// UpdatePolicy 更新策略
func (*permissionSrv) UpdatePolicy() error {
	// 角色
	qr := query.Role
	roles, err := qr.Preload(qr.Permissions, qr.Permissions.Permissions).Find()
	if err != nil {
		return err
	}
	adapter := CasbinSrv.Casbin().GetAdapter()
	for _, role := range roles {
		if role.Permissions == nil {
			continue
		}
		for _, permission := range role.Permissions {
			//p, Member, team, team-create
			po := []string{role.Code, permission.Code, "allow"}
			if !CasbinSrv.CheckPolicy("p", po) {
				adapter.AddPolicy("p", "p", []string{role.Code, permission.Code, "allow"})
			}
		}
	}
	return nil
}
