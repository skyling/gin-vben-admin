package admin

import (
	"gin-vben-admin/api"
	"gin-vben-admin/api/admin/types"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

type DeptsRoot struct {
}

func (req *DeptsRoot) Run(c *gin.Context) e.Response {
	depts, _ := repo.DeptSrv.Roots(auth.CurrentTenantID(c).Int64())
	return types.RespDeptItems(depts)
}

type GetDeptsReq struct {
	api.PageReq
	Name   string `form:"name" json:"name" binding:"omitempty" field:"Name" example:""`        // Name
	Status int64  `form:"status" json:"status" binding:"omitempty" field:"Status" example:"0"` // Status
}

func (req *GetDeptsReq) Run(c *gin.Context) e.Response {
	params := repo.DeptListsParams{
		Name:   req.Name,
		Status: req.Status,
	}
	params.Limit = req.GetLimit()
	params.Offset = req.GetOffset()
	params.TenantID = auth.CurrentTenantID(c)
	depts, cnt, _ := repo.DeptSrv.Lists(params)
	return types.RespDeptsTree(depts, cnt)
}

type GetDeptAllReq struct {
	Name string `form:"name" json:"name" binding:"omitempty" field:"Name" example:""`
}

func (req *GetDeptAllReq) Run(c *gin.Context) e.Response {
	params := repo.DeptListsParams{
		Name: req.Name,
	}
	params.TenantID = auth.CurrentTenantID(c)
	depts, _ := repo.DeptSrv.All(params)
	return types.RespDeptItems(depts)
}

type CreateDeptReq struct {
	Name     string       `form:"name" json:"name" binding:"required" field:"Name" example:""`              // Name
	Remark   string       `form:"remark" json:"remark" binding:"omitempty" field:"Remark" example:""`       // Remark
	Sort     int64        `form:"sort" json:"sort" binding:"omitempty" field:"Sort" example:""`             // Sort
	ParentID snowflake.ID `form:"parentId" json:"parentId" binding:"omitempty" field:"ParentID" example:""` //  ParentID
}

func (req *CreateDeptReq) Run(c *gin.Context) e.Response {
	tid := auth.CurrentTenantID(c)
	dept := &models.Dept{
		Name:     req.Name,
		Remark:   req.Remark,
		Sort:     req.Sort,
		ParentID: req.ParentID,
		TenantID: tid,
	}
	dept, err := repo.DeptSrv.Create(dept)
	if err != nil {
		return e.ErrSystemError.WithError(err).Resp()
	}
	return e.RespOK
}

type UpdateDeptReq struct {
	api.IDReq
	Name     string        `form:"name" json:"name" binding:"required" field:"" example:""`                  //
	Remark   string        `form:"remark" json:"remark" binding:"omitempty" field:"Remark" example:""`       // Remark
	Status   models.Status `form:"status" json:"status" binding:"required,status" field:"Status" example:""` // Status
	Sort     int64         `form:"sort" json:"sort" binding:"omitempty" field:"Sort" example:""`             // Sort
	ParentID snowflake.ID  `form:"parentId" json:"parentId" binding:"omitempty" field:"ParentID" example:""` //  ParentID
}

// Run 更新角色
func (req *UpdateDeptReq) Run(c *gin.Context) e.Response {
	id, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	tid := auth.CurrentTenantID(c)
	dept, err := query.Dept.GetByIDAndTID(tid.Int64(), id.Int64())
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	if req.Name != "" {
		dept.Name = req.Name
	}
	if req.Remark != "" {
		dept.Remark = req.Remark
	}
	dept.Status = req.Status
	if req.Sort > 0 {
		dept.Sort = req.Sort
	}
	if req.ParentID > 0 && req.ParentID != dept.ID {
		pdept, _ := query.Dept.GetByIDAndTID(tid.Int64(), req.ParentID.Int64())
		if pdept != nil {
			dept.ParentID = pdept.ID
		}
	} else {
		dept.ParentID = 0
	}
	err = repo.DeptSrv.Save(dept)
	if err != nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	return e.RespOK
}

type UpdateDeptStatusReq struct {
	api.IDReq
	Status models.Status `form:"status" json:"status" binding:"required,status" field:"Status" example:""` // Status
}

func (req *UpdateDeptStatusReq) Run(c *gin.Context) e.Response {
	id, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	tid := auth.CurrentTenantID(c)
	dept, err := query.Dept.GetByIDAndTID(tid.Int64(), id.Int64())
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	dept.Status = req.Status
	err = repo.DeptSrv.Save(dept)
	if err != nil {
		return e.ErrParamErr.WithError(err).Resp()
	}
	return e.RespOK
}

type DeleteDeptReq struct {
	api.IDReq
}

func (req *DeleteDeptReq) Run(c *gin.Context) e.Response {
	id, err := req.GetID(c)
	if err != nil {
		return e.ErrParamErr.Resp()
	}
	tid := auth.CurrentTenantID(c)
	dept, err := query.Dept.GetByIDAndTID(tid.Int64(), id.Int64())
	if err != nil || dept == nil {
		return e.ErrParamErr.Resp()
	}
	if repo.DeptSrv.CheckHasUser(id) {
		return e.ErrExist.WithMsg("部门已关联用户,无法删除").Resp()
	}
	err = repo.DeptSrv.Delete(id)
	return e.RespOK
}
