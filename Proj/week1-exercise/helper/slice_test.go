package helper

import "testing"

func TestLast(t *testing.T) {
	//Test true case
	list := []int{1, 2, 3, 4, 5, 6}
	expected := 6
	actual := Last(list)
	if actual != expected {
		t.Error("actual should be 6")
	}
}
func TestLast_1(t *testing.T) {
	//Test case false
	list := []int{1, 2, 3, 4, 6, 5}
	expected := 6
	actual := Last(list)

	if expected == actual {
		t.Error("actual should be 5")
	}
}
