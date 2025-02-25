package bank

import (
	"fmt"
)

var bankBalance float64

func Bank() {

	fmt.Println("Welcome to GO bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		GetBankBalance()
	case 2:
		DepositMoney()
	}
}

func GetBankBalance() {
	fmt.Println("Bank Balance", bankBalance)
}

func DepositMoney() {
	var moneyAdded float64
	fmt.Print("How much to add: ")
	fmt.Scan(&moneyAdded)

	fmt.Println("Added $", moneyAdded)

	bankBalance = moneyAdded + bankBalance
	fmt.Println("New Balance: $", bankBalance)
}
