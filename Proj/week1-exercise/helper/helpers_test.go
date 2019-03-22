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

func TestContainsInt(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	value := 5
	expected := true
	actual := ContainsInt(list, value)
	if expected != actual {
		t.Error("actual should be true")
	}
}

func TestContainsInt2(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	value := 10
	expected := false
	actual := ContainsInt(list, value)
	if expected != actual {
		t.Error("actual should be false")
	}
}

func TestContainsString(t *testing.T) {
	list := []string{"thuan", "tam", "lam"}
	value := "thuan"
	expected := true
	actual := ContainsString(list, value)
	if actual != expected {
		t.Error("actual should be true")
	}
}

func TestContainString(t *testing.T) {
	list := []string{"thuan", "tam", "lam"}
	value := "hung"
	expected := false
	actual := ContainsString(list, value)
	if expected != actual {
		t.Error("actual should be false")
	}
}
