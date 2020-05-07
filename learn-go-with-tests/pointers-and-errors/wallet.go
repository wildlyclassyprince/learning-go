package main

import (
	"errors"
	"fmt"
)

// Stringer on Bitcoin
type Stringer interface {
	String() string
}

// Stringer method for printing string format
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Bitcoin currency
type Bitcoin int

// Wallet captures account transactions
type Wallet struct {
	balance Bitcoin
}

// Deposit method
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance method
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw method
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}
