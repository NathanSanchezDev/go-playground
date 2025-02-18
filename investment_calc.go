package main

import (
	"fmt"
	"math"
)

func RunInvestmentCalculator() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	readData("How much to invest?: ", &investmentAmount)
	readData("Expected return rate?: ", &expectedReturnRate)
	readData("How many years to invest?: ", &years)

	CalculateInvestment(investmentAmount, expectedReturnRate, years)
}

func readData(printValue string, scanPointer *float64) {
	fmt.Print(printValue)
	fmt.Scan(scanPointer)
}

func CalculateInvestment(investmentAmount float64, expectedReturnRate float64, years float64) {
	calc := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realCalc := calc / math.Pow(1+2.5/100, years)

	fmt.Println("Investment", calc)
	fmt.Println("Investment counting inflation", realCalc)
}
