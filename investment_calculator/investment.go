package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5

	TL := getInputUserTL("Total investment = ")
	ERR := getInputUserERR("Expected Result Rate = ")
	YEARS := getInputUserYears("Years = ")

	futureResultInvestment, futureValueWithInflation := calculateFutureValueInvestments(TL, ERR, YEARS, inflationRate)

	formatterFutureResultInvestment := fmt.Sprintf("investment result = %.2f\n", futureResultInvestment)
	formatterFutureResultInflation := fmt.Sprintf("investment result with inflation = %.2f\n", futureValueWithInflation)

	fmt.Print(formatterFutureResultInvestment, formatterFutureResultInflation)
}

func getInputUserTL(text string) float64 {
	var investTotal float64
	fmt.Print(text)
	fmt.Scan(&investTotal)
	return investTotal
}

func getInputUserERR(text string) float64 {
	var expectedResultRate float64
	fmt.Print(text)
	fmt.Scan(&expectedResultRate)
	return expectedResultRate
}

func getInputUserYears(text string) float64 {
	var Years float64
	fmt.Print(text)
	fmt.Scan(&Years)
	return Years
}

func calculateFutureValueInvestments(investTotal, expectedResultRate, year, inflationRate float64) (fv float64, fvi float64) {
	fv = investTotal * math.Pow(1+expectedResultRate/100, year)
	fvi = fv / math.Pow(1+inflationRate/100, year)
	//return fv, fvi
	return
}
