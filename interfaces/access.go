package main

import "errors"

type Access struct {
	balance int
}

func NewAccess() *Access {
	return &Access{
		balance: 0,
	}
}

func (a *Access) GetBalance() int {
	return a.balance
}

func (a *Access) Deposit(amount int) {
	a.balance += amount
}

func (a *Access) Withdraw(amount int) error {
	newBalance := a.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient balance")
	}
	a.balance = newBalance
	return nil
}
