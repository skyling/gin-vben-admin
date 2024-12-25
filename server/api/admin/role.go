package admin

import (
	"gin-vben-admin/api"
	"gin-vben-admin/api/admin/types"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"gin-vben-admin/tasks"
	"github.com/gin-gonic/gin"
	"time"
)

type GetRolesReq struct {
	api.PageReq
	Name   string `form:"name" json:"name" binding:"omitempty" field:"Name" example:""`        // Name
	Code   string `form:"code" json:"code" binding:"omitempty" field:"Code" example:""`        // Code
	Status int64  `form:"status" json:"status" binding:"omitempty" field:"Status" example:"0"` // Status
}

func (req *GetRolesReq) Run(c *gin.Context) e.Response {
	params := repo.RoleListParams{
		Name: req.Name, Status: req.Status, Code: req.Code,
	}
	params.TenantID = auth.CurrentTenantID(c)
	params.Offset = req.GetOffset()
	params.Limit = req.GetLimit()
	roles, cnt, _ := repo.RoleSrv.Lists(params)
	time.Sleep(250 * time.Millisecond)
	return types.RespRoles(roles, cnt)
}

type CreateRoleReq struct {
	Name        string   `form:"name" json:"name" binding:"required" field:"Name" example:""`                      // Name
	Code        string   `form:"code" json:"code" binding:"required" field:"Code" example:""`                      // Code
	Remark      string   `form:"remark" json:"remark" binding:"omitempty" field:"Remark" example:""`               // Remark
	Sort        int64    `form:"sort" json:"sort" binding:"omitempty" field:"Sort" example:""`                     // Sort
	Permissions []string `form:"permissions" json:"permissions" binding:"required" field:"Permissions" example:""` // Permissions
}

func (req *CreateRoleReq) Run(c *gin.Context) e.Response {
	tid := auth.CurrentTenantID(c)
	role := &models.Role{
		Name:     req.Name,
		Code:     req.Code,
		Remark:   req.Remark,
		Sort:     req.Sort,
		TenantID: tid,
	}
	if repo.RoleSrv.CheckCodeExist(req.Code) {
		return e.ErrExist.WithMsg("角色值已存在").Resp()
	}
	if req.Permissions != nil && len(req.Permissions) > 0 {
		permissions, _ := repo.PermissionSrv.GetPermissionsByCode(req.Permissions)
		if permissions != nil {
			role.Permissions = permissions
		}
	}
	role, err := repo.RoleSrv.Create(role)
	if err != nil {
		return e.ErrSystemError.WithError(err).Resp()
	}
	tasks.RunRoleCasbinPolicyTask(role.ID)
	return e.RespOK
}

type UpdateRoleStatusReq struct {
	api.IDReq
	Status models.Status `form:"status" json:"status" binding:"required" field:"Status" example:""` // Status
}

func (req *UpdateRoleStatusReq) Run(c *gin.Context) e.Response {
	id, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	role, err := repo.RoleSrv.GetRoleById(id)
	if req.Status != 0 && req.Status > 0 && req.Status < 3 {
		role.Status = req.Status
		err = repo.RoleSrv.Save(role)
	}
	return e.RespOK
}

type UpdateRoleReq struct {
	api.IDReq
	Name        string        `form:"name" json:"name" binding:"required" field:"" example:""`                           // 名称
	Remark      string        `form:"remark" json:"remark" binding:"omitempty" field:"Remark" example:""`                // Remark
	Status      models.Status `form:"status" json:"status" binding:"required" field:"Status" example:""`                 // Status
	Sort        int64         `form:"sort" json:"sort" binding:"omitempty" field:"Sort" example:""`                      // Sort
	Permissions []string      `form:"permissions" json:"permissions" binding:"omitempty" field:"Permissions" example:""` // Permissions
}

// Run 更新角色
func (req *UpdateRoleReq) Run(c *gin.Context) e.Response {
	id, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	role, err := repo.RoleSrv.GetRoleById(id)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	role.Name = req.Name
	role.Remark = req.Remark

	if req.Status != 0 && req.Status > 0 && req.Status < 3 {
		role.Status = req.Status
	}
	role.Sort = req.Sort

	if req.Permissions != nil {
		permissions, _ := repo.PermissionSrv.GetPermissionsByCode(req.Permissions)
		if permissions != nil {
			role.Permissions = permissions
		}
	} else {
		req.Permissions = nil
	}
	err = repo.RoleSrv.Save(role)
	if err != nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	tasks.RunRoleCasbinPolicyTask(role.ID)
	return e.RespOK
}

type DeleteRoleReq struct {
	api.IDReq
}

func (req *DeleteRoleReq) Run(c *gin.Context) e.Response {
	id, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	if repo.RoleSrv.CheckRoleHasUser(id) {
		return e.ErrExist.WithMsg("角色已关联用户,无法删除").Resp()
	}
	tasks.RunRoleCasbinPolicyTask(id)
	err = repo.RoleSrv.Delete(id)
	return e.RespOK
}
