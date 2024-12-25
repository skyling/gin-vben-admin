package routers

import (
	"gin-vben-admin/api"
	"gin-vben-admin/api/admin"
	"gin-vben-admin/api/common"
	"gin-vben-admin/middleware"
	"github.com/gin-gonic/gin"
)

var (
	mAuth = middleware.AuthRequired()
	cr    = middleware.CheckRight
)

func AdminRouter(r *gin.Engine) {
	// 路由
	unAuthV1 := r.Group("/mapi/")
	{
		unAuthV1.GET("avatar/:key", common.GenAvatarPicture)
	}

	adminV1 := r.Group("/mapi/auth")
	{
		adminV1.POST("login", api.Req(&admin.LoginReq{}))
		adminV1.POST("register", api.Req(&admin.RegisterReq{}))
	}

	authV1 := r.Group("/mapi/", middleware.CurrentUser(), mAuth)
	{
		authV1.POST("auth/logout", api.Req(&admin.LogoutReq{}))
		authV1.GET("options/userType", api.Req(&admin.UserTypeOptionsReq{}))
		SystemRouter(authV1) // 系统用户
	}

}

func SystemRouter(r *gin.RouterGroup) {

	r.GET("user/loginLogs", api.Req(&admin.LoginLogReq{}))
	r.GET("user/info", api.Req(&admin.UserInfoReq{}))
	r.GET("auth/codes", api.Req(&admin.UserPermissionsReq{}))
	r.PUT("user/password", api.Req(&admin.ChangePasswordReq{}))
	r.POST("logout", api.Req(&admin.LogoutReq{}))

	// 角色管理
	r.GET("permissions", cr("role"), api.Req(&admin.PermissionsReq{}))
	r.GET("roles", cr("role", "user"), api.Req(&admin.GetRolesReq{}))
	r.POST("role", cr("role-create"), api.Req(&admin.CreateRoleReq{}))
	r.PUT("role/:id", cr("role-update"), api.Req(&admin.UpdateRoleReq{}))
	r.PUT("role/:id/status", cr("role-update"), api.Req(&admin.UpdateRoleStatusReq{}))
	r.DELETE("role/:id", cr("role-delete"), api.Req(&admin.DeleteRoleReq{}))

	// 部门管理
	r.GET("depts", cr("dept"), api.Req(&admin.GetDeptsReq{}))
	r.GET("deptsAll", cr("user"), api.Req(&admin.GetDeptAllReq{}))
	r.GET("deptsRoot", cr("dept"), api.Req(&admin.DeptsRoot{}))
	r.POST("dept", cr("dept-create"), api.Req(&admin.CreateDeptReq{}))
	r.PUT("dept/:id/status", cr("dept-update"), api.Req(&admin.UpdateDeptStatusReq{}))
	r.PUT("dept/:id", cr("dept-update"), api.Req(&admin.UpdateDeptReq{}))
	r.DELETE("dept/:id", cr("dept-delete"), api.Req(&admin.DeleteDeptReq{}))

	// 用户管理
	r.GET("users", cr("user"), api.Req(&admin.UserListReq{}))
	r.POST("checkUsername", cr("user-create", "user-update"), api.Req(&admin.CheckUsernameReq{}))
	r.POST("user", cr("user-create"), api.Req(&admin.CreateUserReq{}))
	r.PUT("user/:id/status", cr("user-update"), api.Req(&admin.UpdateUserStatusReq{}))

	r.PUT("user/:id", cr("user-update"), api.Req(&admin.UpdateUserReq{}))
	r.DELETE("user/:id", cr("user-delete"), api.Req(&admin.DeleteUserReq{}))
}
