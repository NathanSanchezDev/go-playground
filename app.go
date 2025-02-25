package main

import (
	"fmt"
	"strings"

	"github.com/NathanSanchezDev/go-playground/bank"
	"github.com/NathanSanchezDev/go-playground/calculator"
)

const (
	InvestmentCalculator = "invest"
	ProfitCalculator     = "profit"
	Bank                 = "bank"
)

func DecideExecutionType() {
	var executionType string
	fmt.Print("Which program do you want to run? (invest/profit/bank): ")
	fmt.Scan(&executionType)

	executionType = strings.ToLower(strings.TrimSpace(executionType))

	switch executionType {
	case InvestmentCalculator:
		calculator.RunInvestmentCalculator()
	case ProfitCalculator:
		fmt.Println("Adding soon")
	case Bank:
		bank.Bank()
	default:
		fmt.Println("Invalid input. Please enter 'invest' or 'profit'.")
	}
}

func main() {
	DecideExecutionType()
}
