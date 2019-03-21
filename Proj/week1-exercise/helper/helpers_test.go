package helper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var mySclice []int
	mySclice = append(mySclice, 1)
	mySclice = append(mySclice, 2)
	expected := false
	actual := IsEmpty(mySclice)
	if actual != expected {
		t.Error("Result should be false")
	}
}

func TestIsEmpty_yesEmpty(t *testing.T) {
	var mySclice []int

	fmt.Println(reflect.ValueOf(mySclice))
	expected := true
	actual := IsEmpty(mySclice)
	if actual != expected {
		t.Error("actual should be true")
	}
}
