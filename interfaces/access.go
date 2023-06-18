package main

import "errors"

type Access struct {
	balance float64
}

func NewAccess() *Access {
	return &Access{
		balance: 0,
	}
}

func (a *Access) GetBalance() float64 {
	return a.balance
}

func (a *Access) Deposit(amount float64) {
	a.balance += amount
}

func (a *Access) Withdraw(amount float64) error {
	newBalance := a.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient balance")
	}
	a.balance = newBalance
	return nil
}

func (a *Access) GetName() string {
	return "Access"
}
