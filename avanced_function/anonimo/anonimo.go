package anonimo

import "fmt"

func anonimo() {
	numbers := []int{1, 2, 3, 4, 6}

	anonimosDoubled := createTransform(2)
	anonimostripled := createTransform(3)

	doubled := transformNumber(&numbers, anonimosDoubled)
	tripled := transformNumber(&numbers, anonimostripled)

	anonimos := transformNumber(&numbers, func(number int) int {
		return number * 4
	})
	fmt.Println("Anonimos - ", anonimos)
	fmt.Println("return function anonimo", doubled)
	fmt.Println("return function anonimo", tripled)
}

func transformNumber(number *[]int, tranform func(int) int) []int {
	dNumber := []int{}
	for _, value := range *number {
		dNumber = append(dNumber, tranform(value))
	}
	return dNumber
}

func createTransform(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
