package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

//weekday sẽ thống nhất các hằng số enum
//của chúng ta dưới một kiểu dữ liệu chung
type Weekday int

//Khai báo các hằng số liên quan đến Weekday
//Chỉ định cho chúng các giá trị số khác nhau
//Như vậy chúng sẽ không bị xung đột
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)

//Tạo các hành vi chung cho Enum
//Chúng ta thêm vào các phương thức cho một kiểu
//để xác định hành vi của nó
//Các phương thức được đính kèm sẽ là các phần không thể tách
//rời của Weekday và được chia sẽ với các hằng số của Weekday
//String(): là tên function
func (day Weekday) String() string {
	//Khai báo một mảng các string
	//toán tử ... để đếm số phần tử
	//số phần tử của mảng là (7)
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	//`day`: là một trong các giá trị của hằng số Weekday.
	//Nếu hằng số là Sunday thì day có giá trị là 0
	//Bắt lỗi trong Th `day` nằm ngoài khoảng của Weekday

	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	//return tên của 1 hằng số Weekday từ mảng names bên trên
	return names[day]
}

func (day Weekday) Weekend() bool {

	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}
func main() {
	pp.Println("Start Program")
	fmt.Println(Sunday)
	fmt.Println(Saturday)

	fmt.Printf("Which day it is? %s\n", Sunday)

	fmt.Printf("Is Staturday a weekend day?%t\n", Saturday.Weekend())
}
