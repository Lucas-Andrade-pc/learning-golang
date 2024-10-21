package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age
	// agePointer := &age

	fmt.Println("Address variable age - ", agePointer)

	fmt.Println("Age: ", age)

	adultYear := getAdultYear(agePointer)
	fmt.Println("adult year, ", adultYear)
	fmt.Println("age, ", age)

	editAge(agePointer)
	fmt.Println("age change, ", age)

}

func getAdultYear(age *int) int {
	fmt.Println(age)
	return *age - 18
}

func editAge(age *int) {
	*age = *age - 18

}
