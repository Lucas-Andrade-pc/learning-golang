package main

import "fmt"

func generics() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[R int | float64 | string](a, b R) R {
	return a + b
}
