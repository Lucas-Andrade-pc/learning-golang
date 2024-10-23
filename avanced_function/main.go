package main

import (
	"fmt"
)

func fibb(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	firt := 0
	second := 1
	var result int
	for i := 2; i <= number; i++ {
		fmt.Println(firt, second)
		result = second
		second = firt + second
		firt = result

	}
	return result
}
func fibRecursion(number int) int {
	if number <= 1 {
		return 1
	}
	return fibRecursion(number-1) + fibRecursion(number-2)
}

func fibdyna(number int) int {
	f := make([]int, number+2)
	f[0] = 1
	f[1] = 1
	var result int
	for i := 3; i <= number; i++ {
		result = f[1]
		f[1] = f[0] + f[1]
		f[0] = result

	}
	return f[0]
}
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

}
func main() {
	numbers := []int{1, 2, 3, 4, 6}
	numberDouble := transformNumber(&numbers, double)
	numberTriple := transformNumber(&numbers, triple)

	fmt.Println(numberDouble)
	fmt.Println(numberTriple)

	fi := fibdyna(13)
	fmt.Println(fi)

	fat := factorial(5)
	fmt.Println(fat)

}

func transformNumber(number *[]int, tranform func(int) int) []int {
	dNumber := []int{}
	for _, value := range *number {
		dNumber = append(dNumber, tranform(value))
	}
	return dNumber
}

func double(number int) int {
	return number * 2

}

func triple(number int) int {
	return number * 3

}
