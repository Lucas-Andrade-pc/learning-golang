package main

import "fmt"

type mapStringFloat map[string]float64

func (d mapStringFloat) display() {
	fmt.Println(d)
}

func main() {
	courseRate := make(mapStringFloat, 2)

	courseRate["React"] = 4.6
	courseRate["Angular"] = 3.9

	mapStringFloat.display(courseRate)
}
