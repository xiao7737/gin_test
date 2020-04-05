package user

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

// 名字验证器，用户名不能等于"admin"且长度要大于6
func NameValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	// 利用反射，将反射类型对象转换为借口类型变量
	// 通过断言，恢复底层的具体值
	if s, ok := field.Interface().(string); ok {
		if s == "admin" || len(s) <= 6 {
			return false
		}
	}
	return true
}
