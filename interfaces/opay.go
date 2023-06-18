package main

import "errors"

type Opay struct {
	balance float64
	fee     float64
}

func NewOpay() *Opay {
	return &Opay{
		balance: 0,
		fee:     10,
	}
}
func (o *Opay) GetBalance() float64 {
	return o.balance
}

func (o *Opay) Deposit(amount float64) {
	o.balance += amount
}

func (o *Opay) Withdraw(amount float64) error {
	newBalance := o.balance - amount - o.fee
	if newBalance < 0 {
		return errors.New("insufficient balance")
	}
	o.balance = newBalance
	return nil
}

func (o *Opay) GetName() string {
	return "Opay"
}
