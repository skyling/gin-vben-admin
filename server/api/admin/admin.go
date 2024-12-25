package admin

import (
	"fmt"
	"gin-vben-admin/api"
	"gin-vben-admin/api/admin/types"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"gin-vben-admin/tasks"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"time"
)

// InfoReq 基础请求
type InfoReq struct {
}

func (req *InfoReq) Run(c *gin.Context) e.Response {
	info, err := tasks.RunTestTask(fmt.Sprintf("%s-%d", "hello", time.Now().Unix()))
	return e.RespOK.WithData(map[string]interface{}{
		"info": info,
		"err":  err,
	})
}

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required" field:"用户名" example:"hello"`              // 用户名
	Password string `form:"password" json:"password" binding:"required,min=6,max=64" field:"密码" example:"123456"` // 密码
}

// Run 用户登录
//
//	@Tags			Auth
//	@Summary		用户登录
//	@Description	用户登录
//	@ID				admin.login
//	@Accept			json
//	@Produce		json
//	@Param			data	body		LoginReq							true	"参数"
//	@Success		200		{object}	e.Response{data=types.LoginData}	"返回码: 200"
//	@Router			/login [POST]
func (req *LoginReq) Run(c *gin.Context) e.Response {
	u, _ := repo.UserSrv.GetUserByUsername(req.Username)
	if u == nil {
		return e.ErrUsernamePwdErr.Resp()
	}
	if ok, err := u.CheckPassword(req.Password); !ok || err != nil {
		return e.ErrUsernamePwdErr.Resp()
	}
	return loginByUser(c, u)
}

// 用户登录
func loginByUser(c *gin.Context, eu *models.User) e.Response {
	if eu.Status == models.UserStatusOFF {
		return e.ErrForbidden.Resp()
	}
	token, err := auth.GenToken(eu.ID, "default")
	if err != nil {
		return e.ErrSystemError.Resp()
	}
	repo.UserSrv.SaveLoginLog(eu, auth.CurrentSource(c), c.ClientIP(), c.Request.UserAgent())
	repo.CasbinSrv.LoadCasbinPolicy()
	return types.RespLogin(token, eu)
}

type RegisterReq struct {
	Name     string `form:"name" json:"name" binding:"required" field:"Nickname" example:"name"`                  // Nickname
	Username string `form:"username" json:"username" binding:"required" field:"Username" example:"hello"`         // Username
	Password string `form:"password" json:"password" binding:"required,min=6,max=64" field:"Password" example:""` // Password
}

func (req *RegisterReq) Run(c *gin.Context) e.Response {
	u := &models.User{
		Type:     models.UserTypeTenantUser,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
	u.SetPassword(u.Password)
	err := repo.UserSrv.CreateUser(u)
	if err != nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	return e.RespOK
}

type UserDetailReq struct {
	api.IDReq
}

// Run 获取用户详情
// @Tags Admin
// @Summary 获取用户详情
// @Description 获取用户详情
// @ID admin.detail
// @Accept json
// @Produce json
// @Security AuthToken
// @Param data body UserDetailReq true "表单"
// @Success 200 {object} e.Response{data=types.UserItem} "返回码: 200"
// @Router /userinfo [GET]
func (req *UserDetailReq) Run(c *gin.Context) e.Response {
	uid, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	u, err := repo.UserSrv.GetUserDetail(uid)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	return types.RespUser(u)
}

type UserInfoReq struct {
}

func (req *UserInfoReq) Run(c *gin.Context) e.Response {
	uid := auth.CurrentUserID(c)
	if uid == 0 {
		return e.ErrUnauthorized.Resp()
	}
	u, _ := repo.UserSrv.GetUserByID(uid)
	if u == nil {
		return e.ErrNotFound.Resp()
	}
	permissions, _ := repo.CasbinSrv.GetUserPrivileges(uid, "")
	return types.RespUserInfo(u, permissions)
}

type LogoutReq struct {
}

func (req *LogoutReq) Run(c *gin.Context) e.Response {
	claims := auth.CurrentClaims(c)
	auth.DeleteToken(claims.UserID, claims.Source, claims.Time)
	return e.RespOK
}

type UserListReq struct {
	api.PageReq
	Type     string       `form:"type" json:"type" binding:"omitempty" field:"Type" example:""`             // Type
	Username string       `form:"username" json:"username" binding:"omitempty" field:"Username" example:""` // Username
	Name     string       `form:"name" json:"name" binding:"omitempty" field:"Name" example:""`             // Name
	DeptID   snowflake.ID `form:"deptId" json:"deptId" binding:"omitempty" field:"" example:""`             // 部门ID
}

func (req *UserListReq) Run(c *gin.Context) e.Response {
	params := &repo.UserPageParams{
		Username: req.Username,
		Name:     req.Name,
		DeptID:   req.DeptID,
		Type:     req.Type,
	}
	params.Offset = req.GetOffset()
	params.Limit = req.GetLimit()
	params.BaseParams = api.GetBaseParams(c)

	users, cnt, _ := repo.UserSrv.Lists(params)
	return types.RespUsers(users, cnt)
}

type UserPermissionsReq struct {
}

// Run 获取用户权限
func (req *UserPermissionsReq) Run(c *gin.Context) e.Response {
	uid := auth.CurrentUserID(c)
	permissions, _ := repo.CasbinSrv.GetUserPrivileges(uid, "")
	return types.RespUserPermissions(permissions)
}

type CreateUserReq struct {
	Type     string         `form:"type" json:"type" binding:"required,oneof=admin tenant tenant-user" field:"Type" example:"1"` // Type 用户类型
	Username string         `form:"username" json:"username" binding:"required" field:"username" example:""`                     // username
	Password string         `form:"password" json:"password" binding:"required,min=6" field:"password" example:""`               // password
	Roles    []snowflake.ID `form:"roles" json:"roles" binding:"required" field:"Roles" example:""`                              // RoleID
	DeptID   snowflake.ID   `form:"deptId" json:"deptId" binding:"required" field:"Dept ID" example:""`                          // Dept ID
	Name     string         `form:"name" json:"name" binding:"required" field:"Name" example:""`                                 // Name
	Remark   string         `form:"remark" json:"remark" binding:"omitempty" field:"remark" example:""`                          // remark
	Status   models.Status  `form:"status" json:"status" binding:"required,status" field:"Status" example:""`                    // Status
}

func (req *CreateUserReq) Run(c *gin.Context) e.Response {
	lu := auth.CurrentUser(c)

	u := &models.User{
		Type:     req.Type,
		Username: req.Username,
		Name:     req.Name,
		Password: req.Password,
		DeptID:   req.DeptID,
		Remark:   req.Remark,
		Status:   req.Status,
		Roles:    []*models.Role{},
	}

	switch lu.Type {
	case models.UserTypeTenant:
		u.Type = models.UserTypeTenantUser
		u.TenantID = lu.TenantID
	}

	u.SetPassword(u.Password)
	if req.DeptID.Int64() > 0 {
		dept, _ := repo.DeptSrv.GetDeptById(req.DeptID)
		if dept == nil {
			return e.ErrParamErr.WithMsg("部门不存在").Resp()
		}
	}
	if len(req.Roles) > 0 {
		for _, RoleID := range req.Roles {
			role, err := repo.RoleSrv.GetRoleById(RoleID)
			if err != nil || role == nil {
				return e.ErrParamErr.WithMsg("角色不存在").Resp()
			}
			u.Roles = append(u.Roles, role)
		}

	}
	err := repo.UserSrv.CreateUser(u)
	if err != nil {
		return e.ErrSystemError.WithError(err).Resp()
	}
	tasks.RunAddRolesForUserTask(u.ID)
	return e.RespOK
}

type UpdateUserReq struct {
	api.IDReq
	Name     string         `form:"name" json:"name" binding:"required" field:"Name" example:""`                    // Name
	Password string         `form:"password" json:"password" binding:"omitempty,min=6" field:"password" example:""` // password
	Roles    []snowflake.ID `form:"roles" json:"roles" binding:"required" field:"Roles" example:""`                 // RoleID
	DeptID   snowflake.ID   `form:"deptId" json:"deptId" binding:"required" field:"Dept ID" example:""`             // Dept ID
	Remark   string         `form:"remark" json:"remark" binding:"omitempty" field:"remark" example:""`             // remark
	Status   models.Status  `form:"status" json:"status" binding:"required,status" field:"Status" example:""`       // Status
}

func (req *UpdateUserReq) Run(c *gin.Context) e.Response {
	id, _ := req.GetID(c)
	u := auth.CurrentUser(c)
	tid := auth.CurrentTenantID(c)
	user, err := query.User.GetByIDAndTID(tid.Int64(), id.Int64())
	if u.Type == models.UserTypeAdmin {
		user, err = query.User.GetByID(id.Int64())
	}

	if err != nil || user == nil {
		return e.ErrNotFound.WithError(err).Resp()
	}
	user.Name = req.Name
	if req.Password != "" {
		user.SetPassword(req.Password)
	}
	user.Remark = req.Remark
	user.Status = req.Status
	user.DeptID = req.DeptID
	if req.DeptID.Int64() > 0 {
		dept, _ := query.Dept.GetByIDAndTID(tid.Int64(), req.DeptID.Int64())
		if dept == nil {
			return e.ErrParamErr.WithMsg("部门不存在").Resp()
		}
	}
	if len(req.Roles) > 0 {
		for _, roleId := range req.Roles {
			role, err := query.Role.GetByIDAndTID(tid.Int64(), roleId.Int64())
			if err != nil || role == nil {
				return e.ErrParamErr.WithMsg("角色不存在").Resp()
			}
			user.Roles = []*models.Role{role}
		}
	}
	err = repo.UserSrv.SaveUser(user)
	if err != nil {
		return e.ErrSystemError.WithError(err).Resp()
	}
	tasks.RunAddRolesForUserTask(user.ID)
	return e.RespOK
}

type UpdateUserStatusReq struct {
	api.IDReq
	Status models.Status `form:"status" json:"status" binding:"required,status" field:"Status" example:""` // Status
}

func (req *UpdateUserStatusReq) Run(c *gin.Context) e.Response {
	id, _ := req.GetID(c)
	u := auth.CurrentUser(c)
	tid := auth.CurrentTenantID(c)
	user, err := query.User.GetByIDAndTID(tid.Int64(), id.Int64())
	if u.Type == models.UserTypeAdmin {
		user, err = query.User.GetByID(id.Int64())
	}
	if err != nil || user == nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	user.Status = req.Status
	err = repo.UserSrv.UpdateUserStatus(user)
	if err != nil {
		return e.ErrSystemError.WithError(err).Resp()
	}
	return e.RespOK
}

type DeleteUserReq struct {
	api.IDReq
}

func (req *DeleteUserReq) Run(c *gin.Context) e.Response {
	id, _ := req.GetID(c)
	u := auth.CurrentUser(c)
	tid := auth.CurrentTenantID(c)
	user, err := query.User.GetByIDAndTID(tid.Int64(), id.Int64())
	if u.Type == models.UserTypeAdmin {
		user, err = query.User.GetByID(id.Int64())
	}
	if err != nil || user == nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	// 删除用户角色
	repo.CasbinSrv.RemoveUserRoles(user)
	err = repo.UserSrv.Delete(user)
	if err != nil {
		return e.ErrSystemError.WithError(err).Resp()
	}
	return e.RespOK
}

type CheckUsernameReq struct {
	Username string `form:"username" json:"username" binding:"required" field:"Username" example:""` // Username
}

func (req *CheckUsernameReq) Run(c *gin.Context) e.Response {
	if repo.UserSrv.CheckUsername(req.Username) {
		return e.ErrExist.WithMsg("用户名已存在").Resp()
	}
	return e.RespOK
}
