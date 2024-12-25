package repo

import (
	"fmt"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"github.com/bwmarrin/snowflake"
)

var RoleSrv = new(roleSrv)

type roleSrv struct {
}

type RoleListParams struct {
	PageParams
	Name   string
	Code   string
	Status int64
}

func (*roleSrv) Lists(params RoleListParams) ([]*models.Role, int64, error) {
	qr := query.Role
	qqr := qr.Where(qr.TenantID.Eq(params.TenantID.Int64()))
	if params.Name != "" {
		qqr = qqr.Where(qr.Name.Like(fmt.Sprintf("%%%s%%", params.Name)))
	}
	if params.Code != "" {
		qqr = qqr.Where(qr.Code.Eq(params.Code))
	}
	if params.Status > 0 {
		qqr = qqr.Where(qr.Status.Eq(params.Status))
	}
	return qqr.Preload(qr.Permissions).Order(qr.Sort).FindByPage(params.Offset, params.Limit)
}

func (*roleSrv) GetSortNew() int64 {
	qr := query.Role
	var sort int64
	qr.Select(qr.Sort.Max()).Scan(&sort)
	return sort + 1
}

func (*roleSrv) GetRoleById(id snowflake.ID) (*models.Role, error) {
	qr := query.Role
	return qr.Preload(qr.Permissions).GetByID(id.Int64())
}

func (*roleSrv) CheckRoleHasUser(id snowflake.ID) bool {
	qur := query.UserRole
	cnt, _ := qur.Where(qur.RoleID.Eq(id.Int64())).Count()
	return cnt > 0
}

func (*roleSrv) CheckCodeExist(code string) bool {
	qr := query.Role
	cnt, _ := qr.Where(qr.Code.Eq(code)).Count()
	return cnt > 0
}

func (s *roleSrv) Create(role *models.Role) (*models.Role, error) {
	role.Status = models.StatusOn
	if role.Sort == 0 {
		role.Sort = s.GetSortNew()
	}
	qr := query.Role
	err := qr.Create(role)
	return role, err
}

func (s *roleSrv) Save(role *models.Role) error {
	qr := query.Role
	err := qr.Save(role)
	if err == nil {
		qr.Permissions.Model(role).Delete()
		if role.Permissions != nil {
			qr.Permissions.Model(role).Replace(role.Permissions...)
		}
	}
	return err
}

// Delete 删除角色
func (s *roleSrv) Delete(id snowflake.ID) error {
	qr := query.Role
	_, err := qr.Where(qr.ID.Eq(id.Int64())).Delete()
	return err
}
