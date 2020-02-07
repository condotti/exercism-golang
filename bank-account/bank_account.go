//Package account implements a solution of the exercise titled `Bank Account'.
package account

import "sync"

// Account represents the bank account
type Account struct {
	sync.Mutex
	balance int64
	closed  bool
}

// Open creates an account with given deposit.
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{balance: deposit}
}

// Close returns the remaining money and closes the account.
// This fails if the account has already been closed (returns false).
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	payout = a.balance
	a.balance = 0
	a.closed = true
	return payout, true
}

// Balance returns the amount of money in the account.
// This fails if the account has already been closed (returns false).
func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit charges or withdraws given amount of money into/from the account.
// This fails if the account has already been closed (returns false),
// also fails when one tries to withdraw more than the balance.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.closed {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

/*
BenchmarkAccountOperations-4           	14155567	        86.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAccountOperationsParallel-4   	 8847091	       148 ns/op	       0 B/op	       0 allocs/op
*/
