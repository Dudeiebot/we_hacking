package main

import "fmt"

// with these polymorphism you can add more income for the company without tampering ewith the code or it affecting code
type WorkIncome interface {
	source() string
	calculateTotal() int
}

type advertisment struct { // here is a source of income
	comName       string
	totalhour     int
	hourPerClick  int
	pricePerClick int
}

type allShop struct { // here is another source of income
	projectName        string
	noOfProduct        int
	priceOfProductSold int
}

func (a advertisment) source() string {
	return a.comName
}

func (b allShop) source() string {
	return b.projectName
}

func (a advertisment) calculateTotal() int {
	return a.totalhour * a.hourPerClick * a.pricePerClick
}

func (b allShop) calculateTotal() int {
	return b.noOfProduct * b.priceOfProductSold
}

func CalculateTotalInc(wi []WorkIncome) {
	totalIncome := 0
	for _, income := range wi {
		fmt.Println("Income from %s = #%d", income.source(), income.calculateTotal())
		totalIncome += income.calculateTotal()
	}
	fmt.Println("Here is the net income of our company: %d", totalIncome)
}

func main() {
	OluSons := advertisment{
		"OluSons",
		1000,
		500,
		1,
	}
	LukaPro := advertisment{
		"luka pro",
		5000,
		100,
		2,
	}
	SholaAD := allShop{
		"SholaAD",
		500,
		100,
	}
	tempIncome := []WorkIncome{OluSons, LukaPro, SholaAD}
	CalculateTotalInc(tempIncome)
}
