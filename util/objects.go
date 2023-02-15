package util

import (
	"reflect"
)

// SetBoolByFieldName 通过字段名称设定bool值
func SetBoolByFieldName(obj any, field string, value bool) {
    v := reflect.ValueOf(obj).Elem().FieldByName(field)
    if v.IsValid() {
        v.SetBool(value)
    }
}