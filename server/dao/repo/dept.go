package repo

import (
	"fmt"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"github.com/bwmarrin/snowflake"
)

var DeptSrv = new(deptSrv)

type deptSrv struct {
}

func (*deptSrv) Roots(tenantID int64) ([]*models.Dept, error) {
	qr := query.Dept

	return qr.Where(qr.TenantID.Eq(tenantID), qr.ParentID.Eq(0)).Order(qr.Sort).Find()
}

type DeptListsParams struct {
	PageParams
	Name   string
	Status int64
}

func (*deptSrv) Lists(params DeptListsParams) ([]*models.Dept, int64, error) {
	qr := query.Dept
	qqr := qr.Where(qr.TenantID.Eq(params.TenantID.Int64()))
	if params.Name != "" {
		qqr = qqr.Where(qr.Name.Like(fmt.Sprintf("%%%s%%", params.Name)))
	} else {
		qqr = qqr.Where(qr.ParentID.Eq(0))
	}
	if params.Status > 0 {
		qqr = qqr.Where(qr.Status.Eq(params.Status))
	}
	depts, cnt, err := qqr.Order(qr.Sort).Preload(qr.Depts.Order(qr.Sort)).FindByPage(params.Offset, params.Limit)
	return depts, cnt, err
}

func (*deptSrv) All(params DeptListsParams) ([]*models.Dept, error) {
	qr := query.Dept
	qqr := qr.Where(qr.TenantID.Eq(params.TenantID.Int64())).Select(qr.ID, qr.Name, qr.ParentID)
	if params.Name != "" {
		qqr = qqr.Where(qr.Name.Like(fmt.Sprintf("%%%s%%", params.Name)))
	} else {
		qqr = qqr.Where(qr.ParentID.Eq(0))
	}
	if params.Status > 0 {
		qqr = qqr.Where(qr.Status.Eq(params.Status))
	}
	depts, err := qqr.Order(qr.Sort).Preload(qr.Depts.Order(qr.Sort)).Find()
	return depts, err
}

func (*deptSrv) GetDeptById(id snowflake.ID) (*models.Dept, error) {
	qr := query.Dept
	return qr.GetByID(id.Int64())
}

func (s *deptSrv) Create(dept *models.Dept) (*models.Dept, error) {
	dept.Status = models.StatusOn
	qr := query.Dept
	err := qr.Create(dept)
	return dept, err
}

func (s *deptSrv) Save(dept *models.Dept) error {
	qr := query.Dept
	return qr.Save(dept)
}

// Delete 删除部分
func (s *deptSrv) Delete(id snowflake.ID) error {
	qr := query.Dept
	_, err := qr.Where(qr.ID.Eq(id.Int64())).Delete()
	if err == nil {
		// 删除子部门
		qr.Where(qr.ParentID.Eq(id.Int64())).Delete()
	}
	return err
}

func (*deptSrv) CheckHasUser(id snowflake.ID) bool {
	qur := query.User
	cnt, _ := qur.Where(qur.DeptID.Eq(id.Int64())).Count()
	return cnt > 0
}
