package e

const (
	// CodeSuccess 请求成功
	CodeSuccess = 0
	// CodeCreated 已创建
	CodeCreated = 201

	// CodeRedirect 301 跳转
	CodeRedirect = 301
	// CodeUnauthorized 未登录
	CodeUnauthorized = 401
	// CodeForbidden 禁止访问
	CodeForbidden = 403
	// CodeNotFound 资源未找到
	CodeNotFound = 404

	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr       = 40001
	CodeUsernamePwdErr = 40002

	// CodeSignExpired 签名过期
	CodeSignExpired = 40109
	// CodeExist 数据已存在
	CodeExist               = 40002
	CodeInsufficientBalance = 40003

	// CodeSystemError 系统错误
	CodeSystemError = 50001
)

var (
	RespOK                 = Response{Code: CodeSuccess, Msg: "Success"}
	RespRedirect           = Response{Code: CodeRedirect, Msg: "Success"}
	ErrNotFound            = AppError{Code: CodeNotFound, Msg: "Resource Not Found"}
	ErrParamErr            = AppError{Code: CodeParamErr, Msg: "Param Error"}
	ErrUsernamePwdErr      = AppError{Code: CodeUsernamePwdErr, Msg: "Username or Password Invalid"}
	ErrForbidden           = AppError{Code: CodeForbidden, Msg: "Forbidden"}
	ErrUnauthorized        = AppError{Code: CodeUnauthorized, Msg: "UnAuthorized"}
	ErrExpired             = AppError{Code: CodeSignExpired, Msg: "Sign Expired"}
	ErrExist               = AppError{Code: CodeExist, Msg: "Resource Exist"}
	ErrSystemError         = AppError{Code: CodeSystemError, Msg: "System Error"}
	ErrInsufficientBalance = AppError{Code: CodeInsufficientBalance, Msg: "Insufficient balance"}
)
