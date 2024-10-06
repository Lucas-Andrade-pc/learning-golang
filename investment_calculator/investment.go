package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func salveResultFile(investResult, investResultWithInflation string) {
	data := fmt.Sprint(investResult, investResultWithInflation)
	err := os.WriteFile("investResult.txt", []byte(data), 0644)
	if err != nil {
		fmt.Printf("Erro in write file")
		os.Exit(1)
	}
}

func main() {
	const inflationRate float64 = 2.5

	TL := getInputUserTL("Total investment = ")
	ERR, err := getInputUserERR("Expected Result Rate = ")
	if err != nil {
		panic(err)
	}
	YEARS := getInputUserYears("Years = ")

	futureResultInvestment, futureValueWithInflation := calculateFutureValueInvestments(TL, ERR, YEARS, inflationRate)

	formatterFutureResultInvestment := fmt.Sprintf("investment result = %.2f\n", futureResultInvestment)
	formatterFutureResultInflation := fmt.Sprintf("investment result with inflation = %.2f\n", futureValueWithInflation)

	salveResultFile(formatterFutureResultInvestment, formatterFutureResultInflation)

}

func getInputUserTL(text string) float64 {
	var investTotal float64
	fmt.Print(text)
	fmt.Scan(&investTotal)

	if investTotal <= 0 {
		panic("invalid number! number less or = 0")
	}

	return investTotal
}

func getInputUserERR(text string) (float64, error) {
	var expectedResultRate float64
	fmt.Print(text)
	fmt.Scan(&expectedResultRate)
	if expectedResultRate <= 0 {
		return 1000, errors.New("invalid number! number less or = 0")
	}
	return expectedResultRate, nil
}

func getInputUserYears(text string) float64 {
	var Years float64
	fmt.Print(text)
	fmt.Scan(&Years)
	if Years <= 0 {
		panic("invalid number! number less or = 0")
	}
	return Years
}

func calculateFutureValueInvestments(investTotal, expectedResultRate, year, inflationRate float64) (fv float64, fvi float64) {
	fv = investTotal * math.Pow(1+expectedResultRate/100, year)
	fvi = fv / math.Pow(1+inflationRate/100, year)
	//return fv, fvi
	return
}
