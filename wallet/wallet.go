// Package main provides ...
package main

import (
	"fmt"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// Implement "Stringer", defined in fmt,
// defines how type is printed in %s format
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
