package main

import "fmt"

type IBankAccount interface {
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
		account.Deposit(500)

		if err := account.Withdraw(70); err != nil {
			fmt.Printf("ERR: %f\n", err)
		}
		balance := account.GetBalance()
		fmt.Printf("balance = %f\n", balance)
	}
}
