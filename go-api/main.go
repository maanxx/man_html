/*
GIẢI THÍCH Ý NGHĨA 3 THƯ VIỆN ĐƯỢC IMPORT:

1. "github.com/go-chi/chi"
   - Phân loại: Thư viện bên ngoài (Third-party library).
   - Ý nghĩa: Đây là một thư viện Router (Bộ định tuyến) cực kỳ phổ biến và nhẹ trong Go.
   - Tác dụng: Giúp bạn dễ dàng điều hướng các request web (VD: khi người dùng gõ `GET /users`, nó sẽ chỉ đường cho Go biết phải chạy vào hàm nào).
   - Điểm mạnh: Nó tương thích 100% với chuẩn `net/http` mặc định của Go và quản lý Middleware (các lớp lọc như check Token) rất xịn.

2. "github.com/avukadin/goapi/internal/handlers"
   - Phân loại: Package nội bộ (Internal package) của một project.
   - Ý nghĩa: Đây thực chất là code logic nằm trong thư mục `internal/handlers` của dự án mẫu `goapi` (do tác giả avukadin viết).
   - Tác dụng: Nơi này thường chứa các hàm Controller/Handler (chứa logic phản hồi dữ liệu về cho client).
   - Lưu ý: Vì bạn import thẳng tên repo của người khác, Go đã tự động lên Github tải code repo đó về.

3. log "github.com/sirupsen/logrus"
   - Chữ `log` đứng trước: Nghĩa là đặt tên giả (alias). Gọi `log.Info()` thay vì `logrus.Info()`.
   - Tác dụng: Đây là thư viện Ghi log có cấu trúc (Structured logger) nổi tiếng nhất nhì của Go.
   - Điểm mạnh: Có thể in log ra thành định dạng JSON, có chia màu sắc đẹp mắt khi báo lỗi, và phân loại mức độ nghiêm trọng (Info, Warn, Error, Fatal, Panic...).
*/

package main
