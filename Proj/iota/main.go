package main

import (
	"fmt"
)

//weekday sẽ thống nhất các hằng số enum
//của chúng ta dưới một kiểu dữ liệu chung
type Weekday int

//biểu thức iota được lặp lại bởi các hằng số cho đến khi một khai
//báo kiểu hoặc một phép gán khác xuất hiện
//iota tăng 1 sau mỗi dòng ngoại trừ dòng trống và chú thích.

const (
	Sunday Weekday = iota + 1 // value: 1, type: Weekday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf(
		"Sunday=%[1]d (Type=%[1]T)\n"+
			"Monday=%[2]d (Type=%[2]T)\n"+
			"Tuesday=%[3]d (Type=%[3]T)\n"+
			"Wednesday=%[4]d (Type=%[4]T)\n"+
			"Thursday=%[5]d (Type=%[5]T)\n"+
			"Friday=%[6]d (Type=%[6]T)\n"+
			"Saturday=%[7]d (Type=%[7]T)\n",
		Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday,
	)
}
