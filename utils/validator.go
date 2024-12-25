package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

// 参数校验器自定义规则提示
func ProcessErr(u interface{}, err error) string {
	// 如果为nil 说明校验通过
	if err == nil {
		return ""
	}
	//如果是输入参数无效，则直接返回输入参数错误
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	//断言是ValidationErrors
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	// 获取是哪个字段不符合格式
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()
		typeOf := reflect.TypeOf(u)
		// 如果是指针，获取其属性
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		// 通过反射获取filed
		field, ok := typeOf.FieldByName(fieldName)
		if ok {
			// 获取field对应的error_info tag值
			errorInfo := field.Tag.Get("error_info")
			// 返回错误
			return fieldName + ":" + errorInfo
		} else {
			return "缺失error_info"
		}
	}
	return err.Error()
}
