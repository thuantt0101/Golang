package helper

import (
	"reflect"
	"strings"
)

//abc
func Contains(list interface{}, v interface{}) bool {

	//1. Can dung reflect để convert interface truyen vao
	// thanh doi tuong reflect.value
	listValue := reflect.ValueOf(list)
	vValue := reflect.ValueOf(v)

	//2. Can dung reflect để convert interface truyền vào
	// thành reflect.type

	listType := listValue.Type()

	//3. ban đầu kết quả là false
	// mục đích cho việc này là dùng cho viết unit test
	result := false
	//4. với reflect.type chúng ta dùng method kind để biết
	//loại dữ liệu của list là gì

	switch listType.Kind() {
	case reflect.Slice:
		//vì kiểu dữ liệu là slice nên ta dùng hàm len ở đây được
		// Sau do get ra kieu interface{} de compare voi gia tri truyen vao
		for i := 0; i < listValue.Len(); i++ {
			if reflect.DeepEqual(listValue.Index(i).Interface(), v) {
				result = true
				break
			}
		}
		break
	case reflect.String:
		result = strings.Contains(listValue.String(), vValue.String())
		break
	}
	return result
}
