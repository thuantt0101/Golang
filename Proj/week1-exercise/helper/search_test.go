package helper

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	//khai báo slice
	var mySlice []int
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)

	value := 2
	expected := true
	actual := Contains(mySlice, value)
	if actual != expected {
		t.Error("actual should be true")
	}
}

func isFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

//Find can return first item matched in array
func Find(arr interface{}, predicate interface{}) interface{} {
	var res interface{}
	arrvalue := reflect.ValueOf(arr)
	value := reflect.ValueOf(predicate)

	if !isFunc(predicate) {
		for index := 0; index < arrvalue.Len(); index++ {
			if value.Interface() == arrvalue.Index(index).Interface() {
				res = arrvalue.Index(index).Interface()
				break
			}
		}
	} else {
		for index := 0; index < arrvalue.Len(); index++ {
			elem := arrvalue.Index(index)
			in := []reflect.Value{elem}
			result := value.Call(in)[0]
			if result.Bool() == true {
				res = arrvalue.Index(index).Interface()
				break
			}

		}
	}
	return res
}
func TestFind(t *testing.T) {
	arr := []int{1}
	actual := Find(arr, 1)
	expected := 1
	if actual != expected {
		t.Error("Expected 1")
	}
}

func TestFind_WithLen(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	v := Find(arr, 4)
	expected := 4
	if v != expected {
		t.Error("Expected 4")
	}
}

func TestFind_WithNil(t *testing.T) {
	arr := []int{1}
	v := Find(arr, 4)
	if v != nil {
		t.Error("Expected nil")
	}
}
func TestFind_WithString(t *testing.T) {
	arr := []string{"a"}
	v := Find(arr, "a")
	expected := "a"
	if v != expected {
		t.Error("Expected a")
	}
}
func TestFind_WithFunc(t *testing.T) {
	arr := []int{1, 2, 3}
	actual := Find(arr, func(elem int) bool {
		return elem%2 == 0
	})
	expected := 2
	if actual != expected {
		t.Error("Expected 2")
	}
}
