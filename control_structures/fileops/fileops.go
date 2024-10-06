package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadyBalanceFile(dataFile string) (float64, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return 1000, errors.New("failed open file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed convert string -> float64")
	}
	return balance, nil

}

func WriteBalanceToFile(data string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(data, []byte(balanceText), 0644)
}
