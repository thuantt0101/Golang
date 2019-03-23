package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Bây giờ chúng ta cùng khảo sát ví dụ về kênh đa.
// Giả sử chúng ta có 10 công việc cần hoàn thành bởi 4 người
//chúng ta sẽ mô phỏng quá trình thực hiện 10
//công việc này dưới dạng 4 goroutine nhận 10 giá trị từ
//kênh đa như sau:

const (
	numWorkers = 4
	numTasks   = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

//Hàm đọc dữ liệu từ channel
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		//chờ nhận công việc
		//đọc dữ liệu từ channel
		//khi mà hết dữ liệu được đọc thì sẽ ngưng loop
		task, err := <-tasks
		if !err {
			//hết việc
			fmt.Printf("Người thứ %d: hoàn thành công việc được giao!\n", worker)
			return
		}
		fmt.Printf("Người thứ %d: Bắt đầu công việc được giao!\n", worker)

		//giả lập thời gian thực hiện công việc
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Người thứ %d: thực hiện xong %s\n", worker, task)
	}
}

func main() {

	//tạo kênh đa chứa các công việc
	tasks := make(chan string, numTasks)
	wg.Add(numWorkers)
	for gr := 1; gr < numWorkers; gr++ {
		go worker(tasks, gr)
	}

	//tạo chuổi các công việc cần thực thi cho vào kênh
	for t := 1; t <= numTasks; t++ {
		tasks <- fmt.Sprintf("công việc thứ %d.", t)
	}

	close(tasks)
	wg.Wait()
}
