package main

import (
	"fmt"
)

type IBankAccount interface {
	GetName() string
	GetBalance() float64
	Deposit(amount float64)
	Withdraw(amount float64) error
}

func main() {
	myAccounts := []IBankAccount{
		NewAccess(),
		NewBitcoinAccount(),
		NewOpay(),
	}

	for _, account := range myAccounts {
		if bitcoinAccount, ok := account.(*BitcoinAccount); ok {
			bitcoinAccount.Deposit(0.5)   // Custom withdraw for BitcoinAccount
			bitcoinAccount.Withdraw(0.02) // Custom deposit for BitcoinAccount
		} else {
			account.Deposit(500)
			account.Withdraw(70)
		}
		balance := account.GetBalance()
		fmt.Printf("Account: %s, balance = %f\n", account.GetName(), balance)
	}
}
