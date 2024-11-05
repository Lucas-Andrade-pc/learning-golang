package prices

import (
	"fmt"
	"project_calculator/filemanager"
)

type TaxRatesPricesJobs struct {
	taxRates        float64
	inputPrices     []float64
	taxIncludePrice map[string]string
}

func (job *TaxRatesPricesJobs) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
	}
	job.inputPrices = lines
}

func (job *TaxRatesPricesJobs) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, valuePrice := range job.inputPrices {
		cal := job.taxRates * valuePrice
		result[fmt.Sprintf("%.2f", valuePrice)] = fmt.Sprintf("%.2f", cal)
	}
	job.taxIncludePrice = result

	//fmt.Println(job.taxIncludePrice)
	err := filemanager.WriteFileJson(fmt.Sprintf("result_%0.f.json", job.taxRates*100), job)

	if err != nil {
		return
	}

}

func NewTaxRatesPricesJobs(taxRates float64) *TaxRatesPricesJobs {
	return &TaxRatesPricesJobs{
		taxRates: taxRates,
	}
}
