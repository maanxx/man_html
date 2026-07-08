package main

/*
KIẾN THỨC CỐT LÕI VỀ GOROUTINES & CHANNELS TRONG GO

1. Goroutines (go func)
   - Là các luồng (threads) siêu nhẹ do Go runtime quản lý.
   - Cách dùng: Thêm từ khóa `go` trước một lời gọi hàm. VD: `go checkPrice()`.
   - Vấn đề cốt lõi: Hàm main() không chờ các goroutines khác chạy xong. Nếu main() kết thúc, tất cả goroutines cũng bị ép chết theo.
   
2. Channels (chan)
   - Là "đường ống" để các Goroutines giao tiếp và chia sẻ dữ liệu với nhau một cách an toàn.
   - Khởi tạo: `ch := make(chan string)`
   - Gửi dữ liệu: `ch <- "data"`
   - Nhận dữ liệu: `data := <-ch` (Lệnh này sẽ BLOCK/chặn goroutine hiện tại cho đến khi có dữ liệu gửi đến).

3. Select Statement
   - Dùng để chờ lắng nghe nhiều channel cùng lúc (giống như `switch`, nhưng dành cho channels).
   - Khi có nhiều channel sẵn sàng gửi/nhận, `select` sẽ chọn ngẫu nhiên một case.
   - Thường dùng cho Timeout hoặc worker pool.

4. Unbuffered vs Buffered Channel
   - Unbuffered: `make(chan int)`. Phải có người nhận thì người gửi mới gửi được.
   - Buffered: `make(chan int, 5)`. Người gửi có thể gửi tối đa 5 phần tử vào ống mà chưa cần có người nhận ngay.
*/

import (
	"math/rand"
	"fmt"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i:= range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)

}

// check price of the chicken
func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice <= MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

// check price of the tofu
func checkTofuPrice(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice <= MAX_CHICKEN_PRICE{
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
		case website := <-chickenChannel:
			fmt.Printf("\nText sent: Found deal on chicken at:%v", website)
		case website := <-tofuChannel:
			fmt.Printf("\nEmail sent: Found deal on tofu at %v", website)
	}
}

/*
	var myChan = make(chan int, 5)
	go process(myChan)
	for i:= range myChan {
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
*/

/*

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}
*/
