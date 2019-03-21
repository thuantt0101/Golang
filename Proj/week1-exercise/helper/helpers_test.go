package helper

import "testing"

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
