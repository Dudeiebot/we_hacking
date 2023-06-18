package main

import "errors"

type BitcoinAccount struct {
	balance float64
	fee     float64
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     50,
	}
}

func (b *BitcoinAccount) GetBalance() float64 {
	return b.balance
}

func (b *BitcoinAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BitcoinAccount) Withdraw(amount float64) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient balance")
	}
	b.balance = newBalance
	return nil
}
