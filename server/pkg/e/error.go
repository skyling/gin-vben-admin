package e

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mylukin/easy-i18n/i18n"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Response 基础数据结构
type Response struct {
	// 返回码 正确返回时为0或者200xxx, 其他时为错误返回码
	Code int `json:"code"`
	// 数据
	Data interface{} `json:"data,omitempty"`
	// 提示语
	Msg string `json:"message"`
	// 错误信息,生产环境不返回
	Error string `json:"error,omitempty" swaggerignore:"true"`
}

// WithData 响应数据
func (resp Response) WithData(data interface{}) Response {
	resp.Data = data
	return resp
}

// WithMsg 消息
func (resp Response) WithMsg(msg string) Response {
	resp.Msg = msg
	return resp
}

// AppError 应用错误，实现了error接口
type AppError struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	RawError error       `json:"rawError"`
	Data     interface{} `json:"data,omitempty"`
}

// NewError 返回新的错误对象
func NewError(code int, msg string, err error) AppError {
	return AppError{
		Code:     code,
		Msg:      msg,
		RawError: err,
	}
}

// WithError 将应用error携带标准库中的error
func (err AppError) WithError(raw error) AppError {
	if errors.As(raw, &AppError{}) {
		err.Code = raw.(AppError).Code
		err.Msg = raw.(AppError).Msg
		err.RawError = raw.(AppError).RawError
	} else if errors.Is(raw, gorm.ErrRecordNotFound) {
		err.Code = ErrNotFound.Code
		err.Msg = ErrNotFound.Msg
		err.RawError = raw
	} else {
		err.RawError = raw
	}
	return err
}

func (err AppError) WithCodeMsg(code int, msg string) AppError {
	err.Code = code
	err.Msg = msg
	return err
}

func (err AppError) WithData(data interface{}) AppError {
	err.Data = data
	return err
}

func (err AppError) WithCode(code int) AppError {
	err.Code = code
	return err
}

func (err AppError) WithMsg(msg string) AppError {
	err.Msg = msg
	return err
}

func (err AppError) AppendMsg(msg string) AppError {
	if msg == "" {
		return err
	}
	err.Msg = err.Msg + msg
	return err
}

// Error 返回业务代码确定的可读错误信息
func (err AppError) Error() string {
	return err.Msg
}

// Resp 返回消息
func (err AppError) Resp() Response {
	resp := Response{
		Code: err.Code,
		Msg:  i18n.Sprintf(err.Msg),
		Data: err.Data,
	}
	// 生产环境隐藏底层报错
	if err.RawError != nil {
		logrus.Errorf("%v", err.RawError)
		if gin.Mode() != gin.ReleaseMode {
			resp.Error = err.RawError.Error()
		}
	}
	return resp
}
