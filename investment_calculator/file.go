package main

import (
	"errors"
	"fmt"
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

func getFile(data string) (string, error) {
	file, err := os.ReadFile(data)
	if err != nil {
		return "", errors.New("erro open file")
	}
	return string(file), nil
}
