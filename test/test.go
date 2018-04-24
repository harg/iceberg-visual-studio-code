package main

import "errors"

type Account struct {
	balance int
	undo    Undo
}

func (account *Account) Credit(x int) error {
	if x < 0 {
		return errors.New("Negative credit disallowed")
	}
	account.balance += x∏
	account.undo.Add(func() { account.Debit(x) })
	return nil
}

func (account *Account) Debit(x int) error {
	if x < 0 {
		return errors.New("Negative credit disallowed")
	}
	if account.balance-x < 0 {
		return errors.New("Overdraft disallowed")
	}
	account.balance -= x
	account.undo.Add(func() { account.Credit(x) })
	return nil
}
func (account *Account) Balance() int {
	return account.balance
}

func (account *Account) Undo() error {
	return account.undo.Undo()
}
