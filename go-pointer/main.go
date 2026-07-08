package main

/*
KIẾN THỨC CỐT LÕI VỀ POINTERS, WAITGROUP & MUTEX TRONG GO

1. Pointers (Con trỏ) - '&' và '*'
   - Bản chất: Là biến đặc biệt lưu "địa chỉ bộ nhớ" của một biến khác.
   - `&x`: Lấy địa chỉ vùng nhớ của biến `x`.
   - `*p`: Lấy giá trị đang nằm tại địa chỉ mà con trỏ `p` trỏ tới (Dereference).
   - Tại sao dùng? 
     + Để thay đổi trực tiếp giá trị biến gốc truyền vào hàm.
     + Để tránh phải copy một Struct/Mảng lớn khi truyền qua lại giữa các hàm (giúp tối ưu memory).

2. sync.WaitGroup (Chờ Goroutines)
   - Dùng để bắt hàm `main()` chờ một nhóm goroutines chạy xong.
   - `wg.Add(1)`: Báo hiệu có thêm 1 goroutine cần chờ (Nên gọi TRƯỚC KHI gọi chữ `go`).
   - `wg.Done()`: Báo hiệu goroutine này đã chạy xong (Tương đương Add(-1), thường bỏ trong `defer`).
   - `wg.Wait()`: Chặn luồng hiện tại cho đến khi bộ đếm WaitGroup về 0.

3. sync.Mutex & sync.RWMutex (Khóa đồng bộ - Tránh Race Condition)
   - Race Condition: Khi nhiều goroutines cùng đọc/ghi vào một biến chung ở cùng thời điểm -> sai lệch dữ liệu.
   - Mutex (Lock/Unlock): Đảm bảo tại một thời điểm chỉ có ĐÚNG 1 goroutine được truy cập vào tài nguyên.
   - RWMutex (RLock/RUnlock): Tối ưu hơn. Cho phép nhiều goroutines CÙNG ĐỌC đồng thời, nhưng chỉ 1 goroutine được GHI (và khi GHI thì cấm mọi luồng ĐỌC/GHI khác).
*/

import (
	"fmt"
	"time"
	"sync"
)

/* ------------------ Poiter in go ------------------*/
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var results = []string{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	/*
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("Pointer p: %v\n", *p)
	fmt.Printf("No pointer i: %v\n", i)

	p = &i
	*p = 1

	fmt.Printf("Pointer p: %v\n", *p)
	fmt.Printf("No pointer i: %v\n", i)

	var k int32= 2
	i = k
	fmt.Printf("No pointer i: %v\n", i)
*/
/*
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// ? [1 2 4] ?
	*/

	/*
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memoryy location of the thing1 array is: %p", &thing1)
	// var result [5]float64 = square(thing1)
	var result [5]float64 = square(&thing1)

	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
*/
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nTotal execution time: %v", results)

}

/* ------------------ Goroutines ------------------*/
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()

}
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memoryy location of the thing2 array is: %p", &thing2)

	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
