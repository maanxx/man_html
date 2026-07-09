package main

import (
	"errors"
	"fmt"
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
	go func() {
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


