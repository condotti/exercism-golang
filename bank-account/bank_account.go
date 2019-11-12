//Package account implements a solution of the exercise titled `Bank Account'.
package account

import "sync"

// Account represents the bank account
type Account struct {
	balance int64
	active  bool
	sync.Mutex
}

// Open is a ctor of Account
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	a := Account{balance: deposit, active: true}
	return &a
}

// Close closes the account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.active {
		payout = a.balance
		a.balance = 0
		a.active = false
		return payout, true
	}
	return 0, false
}

// Balance returns the amount of the account
func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.active {
		return a.balance, true
	}
	return 0, false
}

// Deposit deposits money in the account
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.active {
		if a.balance+amount < 0 {
			return 0, false
		}
		a.balance += amount
		return a.balance, true
	}
	return 0, false
}
