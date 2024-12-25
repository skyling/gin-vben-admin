package repo

import (
	"fmt"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"gin-vben-admin/global"
	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/duke-git/lancet/v2/slice"
	"strings"
	"sync"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

type casbinSrv struct{}

var CasbinSrv = new(casbinSrv)

// ParamsMatchFunc 自定义规则
func (s *casbinSrv) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return s.ParamsMatch(name1, name2), nil
}

func (s *casbinSrv) AddRolesForUser(u *models.User) error {
	if u.Roles == nil {
		u, _ = UserSrv.GetUserWithRoles(u.ID)
	}

	ukey := s.RoleKey(u.ID)
	s.Casbin().DeleteRolesForUser(ukey)
	if u.Roles == nil {
		return nil
	}
	roleCode := []string{}
	for _, role := range u.Roles {
		roleCode = append(roleCode, role.Code)
	}
	_, err := s.Casbin().AddRolesForUser(ukey, roleCode)
	return err
}

func (s *casbinSrv) RemoveUserRoles(u *models.User) error {
	return s.AddRolesForUser(u)
}

func (s *casbinSrv) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.DB)
		if err != nil {
			panic(err)
		}
		mt, err := model.NewModelFromString(models.CasbinModelText)
		if err != nil {
			panic(err)
		}
		syncedEnforcer, err = casbin.NewSyncedEnforcer(mt, a)
		if err != nil {
			panic(err)
		}
		syncedEnforcer.AddFunction("ParamsMatch", s.ParamsMatchFunc)
		_ = syncedEnforcer.LoadPolicy()
	})
	return syncedEnforcer
}

func (s *casbinSrv) UpdatePolicy(roleId snowflake.ID) error {
	// 角色
	adapter := s.Casbin().GetAdapter()
	qr := query.Role
	qqr := qr.Order(qr.ID)
	if roleId.Int64() > 0 {
		qqr = qqr.Where(qr.ID.Eq(roleId.Int64()))
	}
	qcb := query.CasbinRule
	roles, err := qqr.Preload(qr.Permissions).Find()
	if err != nil {
		return err
	}
	for _, role := range roles {
		if role.Permissions == nil {
			continue
		}
		exists := []string{}
		for _, permission := range role.Permissions {
			//p, Member, team, team-create
			po := []string{role.Code, permission.Code, "allow"}
			exists = append(exists, permission.Code)
			if !CasbinSrv.CheckPolicy("p", po) {
				adapter.AddPolicy("p", "p", []string{role.Code, permission.Code, "allow"})
			}
		}
		// 删除已经移除的
		qcb.Where(qcb.Ptype.Eq("p"), qcb.V0.Eq(role.Code), qcb.V1.NotIn(exists...)).Delete()
	}
	return nil
}

// LoadCasbinPolicy 获取规则
func (*casbinSrv) LoadCasbinPolicy() error {
	return syncedEnforcer.LoadPolicy()
}

// ParamsMatch 自定义规则
func (*casbinSrv) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

// IsAllowed 是否允许
func (*casbinSrv) IsAllowed(sub, obj, act string) bool {
	ret, _ := syncedEnforcer.Enforce(sub, obj, act)
	return ret
}

// RoleKey 角色key
func (*casbinSrv) RoleKey(uid snowflake.ID) string {
	return fmt.Sprintf("role::%d", uid)
}

// UpdateRole 用户角色处理
func (s *casbinSrv) UpdateRole(uid snowflake.ID, roleName string) (bool, error) {
	defer func() {
		_ = syncedEnforcer.LoadPolicy()
	}()
	key := s.RoleKey(uid)
	return s.Casbin().AddRoleForUser(key, roleName)
}

// GetUserPrivileges 获取权限点ID
func (s *casbinSrv) GetUserPrivileges(uid snowflake.ID, domain string, rights ...string) ([]string, error) {
	userKey := s.RoleKey(uid)
	privileges, err := s.Casbin().GetImplicitPermissionsForUser(userKey)
	if err != nil {
		return nil, err
	}
	var permissions []string
	for _, pp := range privileges {
		if domain != "" {
			if pp[1] == domain {
				permissions = append(permissions, pp[2])
			}
		} else {
			permissions = append(permissions, pp[1])
		}
	}
	if len(rights) > 0 {
		permissions = slice.Intersection(rights, permissions)
	}
	return permissions, nil
}

func (s *casbinSrv) CheckUserRight(uid snowflake.ID, domain string, rights ...string) bool {
	if len(rights) == 0 {
		return false
	}
	permissions, err := s.GetUserPrivileges(uid, domain, rights...)
	if err != nil {
		return false
	}
	if len(permissions) == 0 {
		return false
	}
	return true
}

func (s *casbinSrv) CheckPolicy(ptype string, rights []string) bool {
	if len(rights) == 0 {
		return false
	}
	csq := query.CasbinRule
	cq := csq.Where(csq.Ptype.Eq(ptype))
	for i, right := range rights {
		switch i {
		case 0:
			cq.Where(csq.V0.Eq(right))
		case 1:
			cq.Where(csq.V1.Eq(right))
		case 2:
			cq.Where(csq.V2.Eq(right))
		case 3:
			cq.Where(csq.V3.Eq(right))
		case 4:
			cq.Where(csq.V4.Eq(right))
		case 5:
			cq.Where(csq.V5.Eq(right))
		}
	}
	cnt, _ := cq.Count()
	return cnt > 0
}
