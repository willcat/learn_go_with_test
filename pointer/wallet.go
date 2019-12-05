package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) WithDraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return errors.New("funds insufficient")
	}

	w.balance -= amount
	return nil
}