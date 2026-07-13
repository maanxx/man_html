/*
===========================================================================
📚 TỔNG HỢP KIẾN THỨC GO ĐÃ HỌC TỪ PROJECT BANKING API
===========================================================================

1. STRUCT & INTERFACE (Nền tảng OOP trong Go)
  - Go không có Class. Dữ liệu được lưu trong Struct (ví dụ: BaseAccount).
  - Interface chỉ định nghĩa "Hành vi" (các hàm), không chứa dữ liệu.
  - Bất kỳ Struct nào viết đủ các hàm của Interface thì tự động được coi là implement Interface đó (không cần từ khóa implements như Java).

2. COMPOSITION (Nhúng Struct - Thay thế cho Kế thừa)
  - Go không có Kế thừa (Inheritance). Để tái sử dụng code, ta dùng Composition (Nhúng).
  - Nhúng `BaseAccount` vào `NormalAccount` giúp NormalAccount sở hữu luôn Owner, Balance và các hàm Deposit, GetBalance mà không cần viết lại code.

3. POLYMORPHISM (Đa hình)
  - Nhờ Interface, ta có thể tạo 1 mảng chứa nhiều kiểu dữ liệu khác nhau.
  - Ví dụ: `[]IAccount` trong Bank có thể chứa cả NormalAccount và CreditAccount.

4. POINTER RECEIVER (Hàm nhận Con trỏ)
  - Khi định nghĩa hàm: `func (b *BaseAccount) Deposit(...)` ta phải dùng dấu `*` (con trỏ).
  - Lý do: Để Go truyền "Bản gốc" của account vào hàm. Nếu không có `*`, Go sẽ copy ra một bản nháp, hàm sửa trên bản nháp thì số dư thực tế sẽ không đổi.

5. CONCURRENCY & DATA RACE (Xử lý đồng thời & Xung đột dữ liệu)
  - Goroutine (`go func()`) giúp chạy nhiều tác vụ song song cùng lúc cực kỳ nhẹ.
  - Khi nhiều Goroutine cùng sửa 1 biến (như Balance), sẽ gây ra Data Race (tiền bị trừ sai).

6. SYNC.MUTEX (Khóa bảo vệ dữ liệu)
  - Dùng `sync.Mutex` để tạo ổ khóa.
  - `Mu.Lock()`: Khóa cửa, chặn các Goroutine khác lại.
  - `defer Mu.Unlock()`: Đảm bảo luôn mở khóa khi hàm chạy xong. Đây là "Thần chú" chống Data Race.

7. RESTful API (Xây dựng Web Server với net/http)
  - `http.HandleFunc("/path", Handler)`: Định tuyến URL tới hàm xử lý.
  - `r.URL.Query().Get("key")`: Lấy dữ liệu từ URL (VD: ?amount=500).
  - `strconv.ParseFloat()`: Chuyển chữ (string) thành số.
  - `json.NewEncoder(w).Encode(...)`: Trả dữ liệu về cho người dùng dưới chuẩn JSON.
  - `http.ListenAndServe(":8080", nil)`: Bật server chạy liên tục.

===========================================================================
*/
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Điều kiện: "1 tài khoản phải có 2 yếu tố này => hợp lệ"
type IAccount interface {
	Deposit(amount float64)
	WithDraw(amount float64) error
	getBalance() float64
}

type Bank struct {
	Accounts []IAccount
}

type NormalAccount struct {
	BaseAccount
}

type CreditAccount struct {
	BaseAccount
	CreditLimit float64
}

type BaseAccount struct {
	Owner string
	Balance float64
	History []string
	Mu sync.Mutex
}


// create a new global account for testing
var bank = Bank{}
var globalAcc = NormalAccount{
	BaseAccount: BaseAccount{
		Owner: "NGUYEN PHAN MINH MAN",
		Balance: 1000.0,
	},
}


// normal account
func (b *NormalAccount)WithDraw(amount float64) error {
	b.Mu.Lock()

	defer b.Mu.Unlock()

	if amount > b.Balance {
		return errors.New("Số dư không đủ")
	}

	time.Sleep(time.Millisecond * 10)

	b.Balance -= amount
	b.History = append(b.History, fmt.Sprintf("Rút thành công %.2f", amount))
	return nil
}
// credit account
func (c *CreditAccount)WithDraw(amount float64) error {

	c.Mu.Lock()

	defer c.Mu.Unlock()

	if (c.Balance - amount) < -c.CreditLimit {
		return errors.New("Vượt quá hạn mức tín dụng")
	}

	c.Balance -= amount
	c.History = append(c.History, fmt.Sprintf("Rút thành công %.2f", amount))
	return nil
}

func (b *BaseAccount)Deposit(amount float64) {

	b.Mu.Lock()

	defer b.Mu.Unlock()
	
	b.Balance += amount
	b.History = append(b.History, fmt.Sprintf("Nạp thành công %.2f", amount))

}

func (b *BaseAccount)PrintHistory() {
	fmt.Println("\n\n\n====================== LỊCH SỬ GIAO DỊCH ===================")
	for _, i:= range b.History {
		fmt.Println("\n", i)
	}
}

func TestTransaction(acc IAccount) {
	acc.Deposit(100.0)
	acc.WithDraw(1500.0)
}

// lấy số dư ứng với 1 acc
func (b *BaseAccount)getBalance() float64{  
	return b.Balance
}

// thêm tài khoản vào ngân hàng
func (b *Bank)addAccount(acc IAccount) {
	b.Accounts = append(b.Accounts, acc)
}

// tính tổng tài sản trong ngân hàng
func (b *Bank)totalAssets() float64 {
	total := 0.0

	for _, v:= range b.Accounts {
		total += v.getBalance()
	}
	return total
}

// mô phỏng 2 người withdraw the same time
func ConcurrentWithDraw(acc IAccount) {
	var wg sync.WaitGroup

	wg.Add(2)

	// first guy
	go func() {
		defer wg.Done()
		fmt.Println("Người 1 rút 200")
		err := acc.WithDraw(200.0)
		if err != nil {
			fmt.Println("Người 1 lỗi", err)
		} else {
			fmt.Println("Người 1 rút thành công")
		}
	}()
	// second guy
	go  func() {
		defer wg.Done()
		fmt.Println("Người 2 rút 300")
		err := acc.WithDraw(300.0)
		if err != nil {
			fmt.Println("Người 2 lỗi", err)
		} else {
			fmt.Println("Người 2 rút thành công")
		}
	}()
	wg.Wait()
}


func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	// tính tổng tiền ngân hàng
	total := bank.totalAssets()

	// trả về json cho client
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"owner": globalAcc.Owner,
		"balance": total,
		"message": "Tra cứu số dư thành công",
	})
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	// lấy số tiền từ URL: deposit?amount=500
	amountStr := r.URL.Query().Get("amount")

	amount, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		http.Error(w, "Số tiền không hợp lệ", http.StatusBadRequest)
		return
	}

	globalAcc.Deposit(amount)

	// trả về thông báo thành công dưới dạng json
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Nạp tiền thành công!",
		"deposited_amount": amount,
		"new_balance": bank.totalAssets(), // lấy số dư mới để hiển thị
	})

	
}

func main() {
	// add account in the bank
	bank.addAccount(&globalAcc)

	http.HandleFunc("/balance", GetBalanceHandler)
	http.HandleFunc("/deposit", DepositHandler)

	fmt.Println("Bank API run at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Lỗi khởi động Server:", err)
	}
}
/*
func main() {
	myAcc := NormalAccount{
		BaseAccount: BaseAccount{
			Owner: "NGUYEN PHAN MINH MAN",
			Balance: 0.0,
		},
	}
	myAcc2 := CreditAccount{
		BaseAccount: BaseAccount{
			Owner: "THAN HOANG THIEN THIEN",
			Balance: 0.0,
		},
		CreditLimit: 1000.0,
	}

	fmt.Printf("Tài khoản của %s vừa tạo. Số dư: %.2f\n", myAcc.Owner, myAcc.Balance)

	// Nạp tiền
	fmt.Println("------- Đang nạp 500K ---------")
	myAcc.Deposit(500.0)
	fmt.Printf("Số dư hiện tại: %.2f\n", myAcc.Balance)

	// Rút tiền
	fmt.Println("------- Đang rút 200K ---------")
	err := myAcc.WithDraw(200.0)
	if err != nil {
		fmt.Println("Lỗi: ", err)
	} else {
		fmt.Printf("Rút tiền thành công. Số dư hiện tại là: %.2f\n", myAcc.Balance)
	}

	// Rút tiền bị lỗi
	fmt.Println("------- Đang rút 1000K ---------")
	err2 := myAcc.WithDraw(1000.0)
	if err2 != nil {
		fmt.Println("Lỗi: ", err2)
	}

	// Số dư hiện tại
	fmt.Printf("Số dư hiện tại: %.2f\n", myAcc.Balance)

	myAcc.PrintHistory()

	fmt.Println("\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	fmt.Printf("\n[Trước] Số dư hiện tại của %s là: %v\n", myAcc.Owner, myAcc.Balance)
	TestTransaction(&myAcc)
	fmt.Printf("[Sau] Số dư hiện tại của %s là: %v\n", myAcc.Owner, myAcc.Balance)

	fmt.Printf("\n[Trước] Số dư hiện tại của %s là: %v\n", myAcc2.Owner, myAcc2.Balance)
	TestTransaction(&myAcc2)
	fmt.Printf("[Sau] Số dư hiện tại của %s là: %v\n", myAcc2.Owner, myAcc2.Balance)


	fmt.Println("\nTỔNG NGÂN SÁCH CÓ TRONG NGÂN HÀNG")

	myBank := Bank{}

	myBank.addAccount(&myAcc)
	myBank.addAccount(&myAcc2)
	fmt.Printf("Tổng tài sản của ngân hàng: %.2f\n", myBank.totalAssets())

	fmt.Println("\nBẮT ĐẦU TEST RÚT TIỀN ĐỒNG THỜI")

	ConcurrentWithDraw(&myAcc)
	fmt.Printf("Số dư hiện tại cuối myAcc: %.2f\n", myAcc.getBalance())


}

*/


