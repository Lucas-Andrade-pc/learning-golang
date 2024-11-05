package convertion

import (
	"errors"
	"strconv"
)

func StringToFloat(number []string) ([]float64, error) {
	resultConvertion := []float64{}
	for _, value := range number {
		convertion, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.New("error convertion string to float64")
		}
		resultConvertion = append(resultConvertion, convertion)
	}
	return resultConvertion, nil
}
