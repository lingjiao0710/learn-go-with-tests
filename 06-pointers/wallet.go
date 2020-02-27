package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

//Bitcoin类型的String方法，可以用于该类型的打印
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//func (w Wallet) Deposit(amount int) {
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

//func (w Wallet) Balance() int {
func (w *Wallet) Balance() Bitcoin {
	//fmt.Println("address of balance in Balance is", &w.balance)
	return w.balance
}

//func (w Wallet) PrintWallet() {
func (w *Wallet) PrintWallet() {
	//fmt.Println("address of balance in PrintWallet is", &w.balance)
	fmt.Printf("%#v has %d\n", w, w.balance)
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
