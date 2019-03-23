package main

import (
	"fmt"
	"time"
)

//khai báo như bên dưới ...int là khai báo 1 slice
//hàm bên dưới trả về 1 channel với n được đọc từ mảng
//và channel out sẽ được ghi vào từ n
func gen(nums ...int) <-chan int {

	//make(chan int):tạo ra một Unbuffered channel
	//Kênh đơn là kênh chỉ chứa tối đa một giá trị dữ liệu.
	//Khi thực hiện một hành động gửi dữ liệu vào kênh đơn
	//goroutine này sẽ bị khóa cho đến khi có goroutine
	//khác nhận dữ liệu thì goroutine này mới tiếp tục được thực thi.
	//Ngược lại nếu một goroutine nhận dữ liệu từ một kênh
	//nhưng kênh chưa có dữ liệu sẵn thì nó sẽ bị khóa cho
	//đến khi có dữ liệu được gửi vào kênh
	//Trao đổi dữ liệu ở kênh đơn buộc các goroutine gửi và nhận dữ liệu
	//phải được đồng bộ nên đôi lúc kênh đơn còn gọi là kênh đồng bộ
	// (synchronous channel).

	//Kênh đa (buffered channel)
	//Kênh đa nghĩa là tại một thời điểm kênh có thể chứa nhiều hơn
	//một giá trị dữ liệu. Số giá trị dữ liệu mà nó có thể chứa được
	//khai báo khi tạo kênh. Cách tạo kênh đa cũng tương tự như tạo
	//kênh đơn với hàm make, có thêm tham số xác định số giá trị dữ
	//liệu tối đa có thể chứa. Cụ thể để tạo một kênh đa chứa 5 giá
	//trị kiểu chuỗi, ta khai báo như sau: ch := make(chan string, 5).
	//Hoạt động gửi dữ liệu sẽ thêm dữ liệu vào cuối danh sách dữ liệu
	//lưu trong kênh. Ngược lại hoạt động nhận sẽ lấy dữ liệu từ đầu
	//danh sách dữ liệu trong kênh. Hàm len(ch) trả về số giá trị dữ
	//liệu đang có trong kênh và hàm cap(ch) cho biết số giá trị dữ
	//liệu tối đa mà kênh có thể chứa.
	//Điều kiện khóa goroutine trên kênh đa cũng thay đổi.
	//Goroutine nhận dữ liệu chỉ bị khóa nếu không còn dữ liệu trong
	//kênh để nhận. Ngược lại, goroutine gửi dữ liệu chỉ bị khóa khi
	//khả năng chứa của kênh đã hết. Điều này khiến cho 2 goroutine
	//gửi và nhận không nhất thiết phải đồng bộ như ở kênh đơn.
	out := make(chan int, 1)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//in la 1 channel,chi duoc phep doc du lieu
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {

	//set up the pupeline
	c := gen(1, 2, 3, 4, 5, 6)
	out := sq(c)

	fmt.Println("capacity of channel c", cap(c))
	fmt.Println("capacity of channel out", cap(out))

	//Consume the output
	//duyệt qua channel out để lấy dữ liệu và in ra màn hình
	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()
	time.Sleep(5 * time.Second)
}
