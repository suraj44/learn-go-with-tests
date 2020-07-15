package fintech

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= w.balance {
		w.balance -= amount
		return nil
	} else {
		return errors.New("ruhroh")
	}

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
