package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	cellphoneRegexString = "^1(\\d){10}$"
	cellphoneRegex       = regexp.MustCompile(cellphoneRegexString)

	numberRegexString           = "[0-9]+"
	alphaRegexString            = "[a-zA-Z]+"
	alphaNumericRegexString     = "^[a-zA-Z0-9]+$"
	alphaNumericDashRegexString = "^[a-zA-Z0-9-]+$"
	// 手机号或邮箱
	usernameRegexString = "^(\\+861(\\d){10}$)|(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"

	alphaRegex            = regexp.MustCompile(alphaRegexString)
	numberRegex           = regexp.MustCompile(numberRegexString)
	alphaNumericRegex     = regexp.MustCompile(alphaNumericRegexString)
	alphaNumericDashRegex = regexp.MustCompile(alphaNumericDashRegexString)
	usernameRegex         = regexp.MustCompile(usernameRegexString)
)

func Init() {
	v := binding.Validator.Engine().(*validator.Validate)
	v.RegisterValidation("cellphone", Cellphone)
	v.RegisterValidation("password", Password)
	v.RegisterValidation("username", Username)
	v.RegisterValidation("alphanumericdash", AlphaNumericDash)
	v.RegisterAlias("status", "oneof=1 2")
	v.RegisterAlias("areaUnit", "oneof=sqm sqft")
	v.RegisterAlias("remark", "max=255")
}

// Username 用户名验证
func AlphaNumericDash(fl validator.FieldLevel) bool {
	return alphaNumericDashRegex.MatchString(fl.Field().String())
}

// Username 用户名验证
func Username(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}

// Cellphone 手机号验证
func Cellphone(fl validator.FieldLevel) bool {
	return cellphoneRegex.MatchString(fl.Field().String())
}

// Password 密码验证 只能是数字和字母且必须包含字母和数字
func Password(fl validator.FieldLevel) bool {
	return numberRegex.MatchString(fl.Field().String()) && alphaRegex.MatchString(fl.Field().String()) && alphaNumericRegex.MatchString(fl.Field().String())
}
