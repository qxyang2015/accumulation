package common

import "reflect"

//支持Array、Slice、String、Map、Chan类型求取长度
func Len(v interface{}) int {
	typeVal := reflect.ValueOf(v)

	switch typeVal.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		return typeVal.Len()
	default:
		return -1
	}
	return typeVal.Len()
}
