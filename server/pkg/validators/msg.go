package validators

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var fieldsMsg = map[string]string{
	"required": "required",
	"min":      "too short",
	"max":      "too long",
	"email":    "格式不正确",
}

func ProcessErr(u interface{}, err error) string {
	if err == nil {
		return ""
	}
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return "参数错误: " + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()
		typeOf := reflect.TypeOf(u)
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}

		field, ok := typeOf.FieldByName(fieldName)
		errorInfo := ""
		if ok {
			errorInfo = field.Tag.Get("emsg")
		}
		if errorInfo == "" {
			fileds := strings.Split(validationErr.Namespace(), ".")[1:]
			for i, filed := range fileds {
				fileds[i] = strutil.SnakeCase(filed)
			}
			nameSpace := strings.Join(fileds, ".")
			errorInfo = fmt.Sprintf("%s %s", nameSpace, validationErr.ActualTag())
		}
		return errorInfo
	}
	return "参数错误"
}

// ValidateMsg 结构体validate后自定义的方法 在tag中定义 msg
func ValidateMsg(data interface{}, err error) (ret map[string]string, rets []string) {
	ret = make(map[string]string)
	//var errorsMsg = make(map[string]string)
	//getMsgTag(data, errorsMsg, "")
	rets = append(rets, "")
	if _, ok := err.(*validator.InvalidValidationError); ok {
		ret["err"] = err.Error()
		rets = append(rets, "")
	} else if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			//key := fmt.Sprintf("%s.%s", err.StructNamespace(), err.ActualTag())
			//s, ok := errorsMsg[key]
			//if ok {
			//	ret[err.StructField()] = s
			//	rets = append(rets, s)
			//	continue
			//}
			//key = fmt.Sprintf("%s._field", err.StructNamespace())
			//sf, okf := errorsMsg[key]
			//if okf {
			//	tagMsg, te := fieldsMsg[err.ActualTag()]
			//	if !te {
			//		tagMsg = "error"
			//	}
			//	s = fmt.Sprintf("%s%s", i18n.Sprintf(sf), i18n.Sprintf(tagMsg))
			//	ret[err.StructField()] = s
			//	rets = append(rets, s)
			//	continue
			//}
			fileds := strings.Split(err.Namespace(), ".")[1:]
			for i, filed := range fileds {
				fileds[i] = strutil.SnakeCase(filed)
			}
			nameSpace := strings.Join(fileds, ".")
			s := fmt.Sprintf("%s %s", nameSpace, err.ActualTag())
			ret[err.StructField()] = s
			rets = append(rets, s)
		}
	} else {
		ret["err"] = err.Error()
	}
	if len(rets) > 1 {
		rets[0] = strings.Join(rets[1:], ";")
	}
	return
}

func getMsgTag(data interface{}, errorsMsg map[string]string, keyPre ...string) {
	st := reflect.TypeOf(data)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	dataName := st.Name()
	if len(keyPre) > 0 {
		dataName = keyPre[0] + dataName
	}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		msg := field.Tag.Get("msg")
		if msg != "" {
			msgs := strings.Split(msg, ",")
			for _, s := range msgs {
				ss := strings.SplitN(s, "=", 2)
				if len(ss) == 2 {
					key := fmt.Sprintf("%s.%s.%s", dataName, field.Name, ss[0])
					errorsMsg[key] = ss[1]
				}
			}
		}

		fieldTitle := field.Tag.Get("field")
		if fieldTitle != "" {
			key := fmt.Sprintf("%s.%s._field", dataName, field.Name)
			errorsMsg[key] = fieldTitle
		}
		switch field.Type.Kind() {
		case reflect.Struct:
			getMsgTag(field, errorsMsg, fieldTitle)
		}
	}
}
