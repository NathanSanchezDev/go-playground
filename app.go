package main

import (
	"fmt"
	"strings"
)

const (
	InvestmentCalculator = "invest"
	ProfitCalculator     = "profit"
)

func DecideExecutionType() {
	var executionType string
	fmt.Print("Which program do you want to run? (invest/profit): ")
	fmt.Scan(&executionType)

	executionType = strings.ToLower(strings.TrimSpace(executionType))

	switch executionType {
	case InvestmentCalculator:
		RunInvestmentCalculator()
	case ProfitCalculator:
		fmt.Println("Adding soon")
	default:
		fmt.Println("Invalid input. Please enter 'invest' or 'profit'.")
	}
}

func main() {
	DecideExecutionType()
}
