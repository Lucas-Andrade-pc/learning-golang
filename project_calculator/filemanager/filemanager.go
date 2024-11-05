package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"project_calculator/convertion"
)

func ReadLines(path string) ([]float64, error) {
	file, err := os.Open("prices.txt")
	if err != nil {
		return nil, errors.New("failed open file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("failed open file")
	}
	prices, err := convertion.StringToFloat(lines)
	if err != nil {
		return nil, errors.New("failed convert string to float")
	}
	return prices, nil
}

func WriteFileJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed create file")
	}
	fmt.Println(data)

	enconder := json.NewEncoder(file)
	err = enconder.Encode(data)
	if err != nil {
		return errors.New("failed convert file")
	}
	file.Close()
	return nil
}
