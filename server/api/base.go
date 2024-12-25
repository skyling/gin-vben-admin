package api

import (
	"encoding/json"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/pkg/auth"
	"gin-vben-admin/pkg/e"
	"gin-vben-admin/pkg/validators"
	"github.com/bwmarrin/snowflake"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type BaseReq interface {
	Run(c *gin.Context) e.Response
}

type IDReq struct {
}

func (IDReq) GetID(c *gin.Context) (snowflake.ID, error) {
	ids := c.Param("id")
	return snowflake.ParseString(ids)
}

func Req(reqb BaseReq) func(c *gin.Context) {
	return func(c *gin.Context) {
		req := convertor.DeepClone(reqb)
		if err := c.ShouldBind(req); err == nil {
			res := req.Run(c)
			c.JSON(200, res)
		} else {
			c.JSON(200, ErrorResponse(err, req))
		}
	}
}

func GetBaseParams(c *gin.Context) repo.BaseParams {
	u := auth.CurrentUser(c)
	return repo.BaseParams{
		UserID:   u.ID,
		TenantID: auth.CurrentTenantID(c),
		UserType: u.Type,
	}
}

type PageReq struct {
	Page int `form:"page" json:"page" binding:"omitempty" field:"页面数" example:"0"`                            // 页面数 默认为1
	Size int `form:"pageSize" json:"pageSize" binding:"omitempty,min=1,max=1000" field:"每页数据条数" example:"10"` // 每页数据条数 默认为20
}

func (s *PageReq) GetLimit() int {
	if s.Size <= 0 {
		s.Size = 20
	}
	if s.Size > 1000 {
		s.Size = 1000
	}
	return s.Size
}

func (s *PageReq) GetOffset() int {
	limit := s.GetLimit()
	if s.Page == 0 {
		s.Page = 1
	}
	return (s.Page - 1) * limit
}

func (s *PageReq) LimitOffset() (int, int) {
	limit := s.GetLimit()
	if s.Page == 0 {
		s.Page = 1
	}
	return limit, (s.Page - 1) * limit
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error, data interface{}) e.Response {
	// 处理 Validator 产生的错误
	if ve, ok := err.(validator.ValidationErrors); ok {
		logrus.Info(data, ve)
		msg := validators.ProcessErr(data, ve)
		return e.ErrParamErr.WithMsg(msg).Resp()

		//_, msg := validators.ValidateMsg(data, ve)
		//return e.ErrParamErr.WithMsg(msg[0]).Resp()
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return e.ErrParamErr.WithMsg("JSON类型不匹配").WithError(err).Resp()
	}

	return e.ErrParamErr.WithError(err).Resp()
}

func ResponseRedirect(c *gin.Context, resp e.Response) {
	if resp.Code == e.CodeRedirect {
		c.Redirect(301, resp.Data.(string))
		return
	}
	if resp.Code != e.CodeSuccess {
		c.JSON(200, resp)
	}
}
