package helper

import "reflect"

/*
IsEmpty: return true if object is empty
*/
func IsEmpty(obj interface{}) bool {
	//1. Can dung reflect de convert interface truyen vao
	// thanh doi tuong reflect.value

	objValue := reflect.ValueOf(obj)

	//2. Can dung reflect de convert interface truyen vao
	//thanh doi tuong reflect.type
	objType := objValue.Type()

	//3. Ban dau ket qua duoc gan la false
	//muc dich phan nay la viet unit test

	result := false

	//4.  voi ket qua reflect.type chung ta dung method kind de biet
	// kieu du lieu cua obj la gi

	switch objType.Kind() {
	case reflect.Slice:
		if objValue.Len() == 0 {
			result = true
		}
	case reflect.Array:
		if objValue.Len() == 0 {
			result = true
		}
	default:
		result = false
	}

	if obj == nil || obj == "" {
		return true
	}
	return result
}
