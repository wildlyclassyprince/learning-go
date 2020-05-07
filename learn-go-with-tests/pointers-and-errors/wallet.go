package main

// Wallet captures account transactions
type Wallet struct {
	balance int
}

// Deposit method
func (w *Wallet) Deposit(amount int) int {
	w.balance += amount
	return w.balance
}

// Balance method
func (w *Wallet) Balance() int {
	return w.balance
}
