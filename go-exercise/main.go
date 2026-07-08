package main


import (
	"io"
	"sync"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// ==========================================
// PHẦN 1: GENERICS & MUTEX (Lõi lưu trữ)
// ==========================================
// Cache là cấu trúc dữ liệu lưu trữ dùng Generics (nhận kiểu T bất kỳ)
type Cache[T any] struct {
	data map[string]T
// Nơi lưu trữ dữ liệu thực tế
	mu sync.RWMutex
// Chìa khóa bảo vệ data khi có nhiều Goroutine cùng truy cập
	notifyChan chan string	
// Ống (Channel) để báo tin khi có dữ liệu mới
}

// NewCache là hàm khởi tạo bộ nhớ cho Cache
func NewCache[T any]() *Cache[T] {
	return &Cache[T]{
		data: make(map[string]T),
		notifyChan: make(chan string, 100),
		// Buffered channel (đệm 100 thông báo)
	}
}

// Set lưu dữ liệu vào Cache
func (c *Cache[T])Set(key string, value T) {
	// Khóa lại: "Chỉ mình tôi được ghi lúc này"
	c.mu.Lock()
	// Đảm bảo mở khóa khi chạy xong hàm
	defer c.mu.Unlock()
	c.data[key] = value

	// Bắn tín hiệu qua Channel (để Goroutine chạy ngầm báo log)
	c.notifyChan <- key

}

// Get lấy dữ liệu từ Cache ra
func (c *Cache[T])Get(key string) (T, bool) {
	c.mu.RLock()
	// Khóa đọc: "Nhiều người cùng đọc được, nhưng không ai được ghi"
	defer c.mu.RUnlock()
	value, exists := c.data[key]

	return value, exists
}

// ==========================================
// PHẦN 2: GOROUTINE & CHANNELS (Chạy ngầm)
// ==========================================

// logWorker liên tục đứng chờ tin nhắn gửi vào ống `ch`
func logWorker(ch <- chan string) {
	// Vòng lặp này sẽ bị BLOCK (dừng lại chờ) cho đến khi có dữ liệu chui vào ống
	for key:= range ch {
		logrus.Infof("📝 [Background Worker] Vừa có người lưu dữ liệu mới vào key: %s", key)
	}
}

// ==========================================
// PHẦN 3: API ROUTER CHI (Giao tiếp Web)
// ==========================================

func main() {
	// 1. Khởi tạo kho lưu trữ (Chọn T là string cho dễ làm API)
	myCache:= NewCache[string]()

	// 2. Kích hoạt Công nhân chạy ngầm (Chạy song song bằng Goroutine)
	go logWorker(myCache.notifyChan)

	// 3. Khởi tạo Router (Chi)
	r := chi.NewRouter()

	// API 1: Lưu dữ liệu (POST)
	// Ví dụ dùng cURL: curl -X POST -d "Du lieu bi mat" http://localhost:8080/cache/secret
	r.Post("/cache/{key}", func(w http.ResponseWriter, req *http.Request) {
		// Lấy chữ "secret" từ URL
		key := chi.URLParam(req, "key")
		// Đọc nội dung gửi kèm từ Body (chữ "Du lieu bi mat")
		body, _:= io.ReadAll(req.Body)

		value := string(body)

		// Lưu vào Cache
		myCache.Set(key, value)

		// Phản hồi cho người dùng
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Đã lưu thành công!\n"))
	}) // LỖI 1: Phải đóng hàm POST ở đây!

	// API 2: Lấy dữ liệu (GET)
	// Ví dụ gõ lên trình duyệt: http://localhost:8080/cache/secret
	r.Get("/cache/{key}", func(w http.ResponseWriter, req *http.Request) {
		key := chi.URLParam(req, "key")

		// Đọc từ Cache
		value, exists := myCache.Get(key)

		if !exists {
			// LỖI 2: Hàm http.Error bắt buộc phải có 3 tham số (w, string_lỗi, status_code). Bạn bị thiếu mã code 404!
			http.Error(w, "Không tìm thấy dữ liệu (404 Not Found)", http.StatusNotFound)
			return // LỖI 3: Chữ return không được mở ngoặc nhọn `{` ở đằng sau như Java/JS
		}
		
		// Trả kết quả về cho người dùng
		// LỖI 4: Ép kiểu mảng byte trong Go là `[]byte`, chứ không phải `[byte]`
		w.Write([]byte("Kết quả: " + value + "\n"))
	})

	// 4. Mở cửa Server đón khách
	// LỖI 5: Viết sai chính tả logrus thành logus!
	logrus.Info("🚀 Server đang khởi chạy tại http://localhost:8080 ...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logrus.Fatal("Server bị sập")
	}
} // LỖI 6: Toàn bộ code Router ở trên phải nằm BÊN TRONG hàm main, bạn bị đóng ngoặc `}` sớm quá ở dòng 80!