package repo

import (
	"fmt"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"github.com/duke-git/lancet/v2/random"
	"time"
)

type userSrv struct{}

var UserSrv = new(userSrv)

type ChangePasswordParams struct {
	UserID      snowflake.ID
	NewPassword string
	OldPassword string
}

func (userSrv) ChangePassword(params *ChangePasswordParams) error {
	qu := query.User
	u, err := qu.Where(qu.ID.Eq(params.UserID.Int64())).First()
	if err != nil {
		return e.ErrNotFound.WithMsg("用户不存在")
	}
	if ok, _ := u.CheckPassword(params.OldPassword); !ok {
		return e.ErrParamErr.WithMsg("旧密码错误")
	}
	u.SetPassword(params.NewPassword)
	_, err = qu.Where(qu.ID.Eq(params.UserID.Int64())).UpdateSimple(qu.Password.Value(u.Password))
	return err
}

type LoginLogPageParams struct {
	PageParams
	CreatedAtStart *time.Time
	CreatedAtEnd   *time.Time
	IP             string
}

func (userSrv) LoginLogs(params *LoginLogPageParams) ([]*models.LoginLog, int64, error) {
	qll := query.LoginLog
	qqll := qll.Where(qll.UserID.Eq(params.UserID.Int64()))
	if params.CreatedAtStart != nil {
		qqll = qqll.Where(qll.CreatedAt.Gte(*params.CreatedAtStart))
	}
	if params.CreatedAtEnd != nil {
		qqll = qqll.Where(qll.CreatedAt.Lte(*params.CreatedAtEnd))
	}
	if params.IP != "" {
		qqll = qqll.Where(qll.IP.Eq(params.IP))
	}
	return qqll.Order(qll.ID.Desc()).FindByPage(params.Offset, params.Limit)
}

type UserPageParams struct {
	PageParams
	Username string
	Name     string
	Type     string
	DeptID   snowflake.ID
}

func (userSrv) Lists(params *UserPageParams) ([]*models.User, int64, error) {
	qu := query.User
	qqu := qu.Scopes(ScopeTenant(params.BaseParams))
	if params.UserType != models.UserTypeAdmin {
		qqu = qqu.Where(qu.ID.Neq(params.UserID.Int64()))
	}
	if params.Username != "" {
		qqu = qqu.Where(qu.Username.Eq(params.Username))
	}
	if params.Name != "" {
		qqu = qqu.Where(qu.Name.Like(fmt.Sprintf("%%%s%%", params.Name)))
	}
	if params.Type != "" {
		qqu = qqu.Where(qu.Type.Eq(params.Type))
	}
	if params.DeptID.Int64() > 0 {
		qqu = qqu.Where(qu.DeptID.Eq(params.DeptID.Int64()))
	}
	return qqu.Preload(qu.Dept, qu.Roles).FindByPage(params.Offset, params.Limit)
}

func (userSrv) GetUserByID(uid snowflake.ID) (*models.User, error) {
	return query.User.GetByID(uid.Int64())
}

func (userSrv) GetUserByIDAndTenantID(tid, uid snowflake.ID) (*models.User, error) {
	qu := query.User
	return qu.Where(qu.TenantID.Eq(tid.Int64()), qu.ID.Eq(uid.Int64())).First()
}

func (userSrv) GetUserWithRoles(uid snowflake.ID) (*models.User, error) {
	qu := query.User
	return qu.Preload(qu.Roles).Where(qu.ID.Eq(uid.Int64())).First()
}

func (userSrv) GetUserDetail(uid snowflake.ID) (*models.User, error) {
	qu := query.User
	return qu.Preload(qu.Roles, qu.Dept).Where(qu.ID.Eq(uid.Int64())).First()
}

func (userSrv) GetUserByUsername(username string) (*models.User, error) {
	qu := query.User
	return qu.Where(qu.Username.Eq(username)).First()
}

func (userSrv) CheckUsername(username string) bool {
	qu := query.User
	cnt, _ := qu.Where(qu.Username.Eq(username)).Count()
	return cnt > 0
}

func (s userSrv) CreateUser(u *models.User) error {
	// 检查用户是否存在
	qu := query.User
	ue, _ := qu.Where(qu.Username.Eq(u.Username)).First()
	if ue != nil {
		return e.ErrExist.WithMsg("用户名已存在")
	}
	u.Code = s.GenCode()
	err := qu.Create(u)
	if err != nil {
		return err
	}
	// 如果是租户, 则让租户ID = 用户ID
	if u.Type == models.UserTypeTenant && u.TenantID == 0 {
		u.TenantID = u.ID
		return qu.Save(u)
	}
	return nil
}

func (s userSrv) SaveUser(u *models.User) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		err := tx.User.Roles.Model(u).Replace(u.Roles...)
		if err != nil {
			return err
		}
		return tx.User.Save(u)
	})
}

func (s userSrv) UpdateUserStatus(u *models.User) error {
	qu := query.User
	_, err := qu.Where(qu.ID.Eq(u.ID.Int64())).UpdateSimple(qu.Status.Value(u.Status.Int64()))
	return err
}

func (s userSrv) Delete(u *models.User) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		err := tx.User.Roles.Model(u).Delete()
		if err != nil {
			return err
		}
		_, err = tx.User.Delete(u)
		return err
	})
}

func (s userSrv) GenCode() string {
	code := fmt.Sprintf("%s%05d", random.RandUpper(1), random.RandInt(0, 99999))
	qu := query.User
	cnt, _ := qu.Where(qu.Code.Eq(code)).Count()
	if cnt > 0 {
		return s.GenCode()
	}
	return code
}

func (userSrv) SaveLoginLog(u *models.User, source, ip, ua string) error {
	qll := query.LoginLog
	ll := models.LoginLog{
		Source:    source,
		UserAgent: ua,
		IP:        ip,
		UserID:    u.ID,
	}
	return qll.Create(&ll)
}
