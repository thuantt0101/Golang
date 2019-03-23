package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//wg được sử dụng để đợi các gorountine kết thúc
var wg sync.WaitGroup

//hàm này chạy trước khi hàm main được thực thi
func init() {
	rand.Seed(time.Now().UnixNano()) //tạo nhân gieo giá trị ngẫu nhiên
}

//Hàm mô tả một người chơi quần vợt
func player(name string, court chan int) {
	//Thông báo người này chơi xong ván đấu
	defer wg.Done()

	for {
		//đợi banh từ đối thủ
		ball, err := <-court
		if !err {
			//thắng nếu kênh đã đóng
			fmt.Printf("%s thắng!\n", name)
			return
		}
		//lấy một giá trị ngẫu nhiên từ 0-99
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s đánh hỏng ở lượt đánh thứ %d!\n", name, ball)
			// Đóng kênh khi đánh hỏng
			close(court)
			return
		}
		fmt.Printf("Lượt đánh bóng thành công thứ %d: %s\n", ball, name)
		ball++
		// Đánh banh về lại đối thủ
		court <- ball
	}
}

func main() {
	court := make(chan int)
	wg.Add(2)
	//tạo 2 người chơi
	go player("Federer", court)
	go player("Djokovic", court)
	//bắt đầu phát bóng cho ván đấu
	court <- 1
	wg.Wait() //đợi ván đấu kết thúc
	fmt.Println("Ván đấu kết thúc!")
}
