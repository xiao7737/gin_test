package user

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

// 名字验证器，用户名不能等于"admin"且长度要大于6
func NameValid(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if s, ok := field.Interface().(string); ok {
		if s == "admin" || len(s) <= 6 {
			return false
		}
	}
	return true
}
