package helper

import "reflect"

//Last get the last element of
func Last(arr interface{}) interface{} {
	arrValue := reflect.ValueOf(arr)
	arrType := arrValue.Type()
	switch arrType.Kind() {
	case reflect.Slice:
		return arrValue.Index(arrValue.Len() - 1).Interface()
	default:
		return arrValue.Interface()
	}
}
