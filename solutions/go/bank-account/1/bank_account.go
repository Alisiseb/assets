package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	closed  bool
}

var Mu sync.Mutex

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, closed: false}
}

func (a *Account) Balance() (int64, bool) {
	if a.closed || a == nil {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed || a == nil {
		return 0, false
	}
	Mu.Lock()
	defer Mu.Unlock()
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	Mu.Lock()
	defer Mu.Unlock()
	if a.closed || a == nil {
		return 0, false
	}
	var closingBalance int64 = 0
	closingBalance, a.closed = a.balance, true
	a.balance = 0
	return closingBalance, true

}
