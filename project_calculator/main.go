package main

import (
	"project_calculator/prices"
)

func main() {
	taxRates := []float64{0.07, 0.1, 0.23, 0.09}

	for _, valueTax := range taxRates {
		result := prices.NewTaxRatesPricesJobs(valueTax)
		result.Process()
	}
	result := prices.NewTaxRatesPricesJobs(1)
	result.Process()
	//result.LoadData()
}
