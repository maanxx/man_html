package main

/*
KIẾN THỨC CỐT LÕI VỀ GENERICS TRONG GO

1. Bản chất:
   - Generics cho phép viết code không phụ thuộc vào một kiểu dữ liệu cụ thể ngay từ đầu.
   - Trì hoãn việc xác định kiểu cho đến khi hàm/struct thực sự được gọi.
   - Giúp tránh việc lặp lại code (ví dụ: không cần viết riêng hàm tính tổng cho int và float).

2. Ba khái niệm nền tảng:
   A. Type Parameter (Tham số kiểu) - [T any]:
      - Nằm trong cặp ngoặc vuông `[...]` ngay sau tên hàm/struct.
      - `T` (hay V, Key...) đại diện cho một kiểu dữ liệu.
      - VD: func Print[T any](s []T) -> nhận vào slice của kiểu T bất kỳ.

   B. Type Constraints (Ràng buộc kiểu):
      - `any` (interface{}) nghĩa là T có thể là bất kỳ kiểu gì.
      - Nếu cần dùng toán tử (như +), phải dùng interface để thu hẹp ràng buộc (Union types).
        VD: type Number interface { int | float64 } -> T chỉ được phép là int hoặc float64.
      - `comparable`: Ràng buộc cho phép các kiểu có thể so sánh `==`, `!=` (Thường làm Key cho Map).

   C. Type Inference (Tự động suy luận kiểu):
      - Khi gọi hàm, trình biên dịch tự hiểu kiểu T dựa vào tham số bạn truyền.
      - VD: Thay vì gọi Sum[int](10, 20), bạn chỉ cần gọi Sum(10, 20).
      - Lưu ý: Khởi tạo Generic Struct bắt buộc phải truyền rõ kiểu (không được suy luận). VD: Stack[int]{}.

3. Khi nào DÙNG / KHÔNG DÙNG:
   - NÊN: Viết cấu trúc dữ liệu đa năng (Stack, Queue, Tree...) hoặc các hàm xử lý chung (Map, Filter, Reverse...).
   - KHÔNG NÊN: Khi một Interface thông thường đã giải quyết được bài toán. Lạm dụng sẽ làm file build nặng và chậm.
*/

import "fmt"

/* ================================*/
/*
func main() {
	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1,2,3}
   fmt.Println(sumSlice[float32](float32Slice))

	fmt.Println("Ready to code!")
}

func sumSlice[T int | float32 | float64](slice []T) T {
   var sum T
   for _, v := range slice{
      sum += v
   }
   return sum
}
*/
/* ================================*/

/*
func main() {
   var intSlice = []int{}
   fmt.Println(isEmpty[int](intSlice))

   var float32Slice = []float32{1,2,3}
   fmt.Println(isEmpty(float32Slice))

}

func isEmpty[T any](slice []T) bool {
   return len(slice) == 0
}
*/
type gasEngine struct {
   gallons float32
   mpg float32
}

type electricEngine struct {
   kwh float32
   mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
   carMake string
   carModel string
   engine T
}
func main() {
   var gasCar = car[gasEngine]{
      carMake: "Honda",
      carModel: "Civic",
      engine: gasEngine{
         gallons: 12.4,
         mpg: 40,
      },
   }
   var electricCar = car[electricEngine]{
      carMake: "Tesla",
      carModel: "Model 3",
      engine: electricEngine{
         kwh: 12.4,
         mpkwh: 40,
      },
   }
   fmt.Println(gasCar)
   fmt.Println(electricCar)

}
