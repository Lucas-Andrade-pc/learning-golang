package main

import "fmt"

type Product struct {
	id      int
	product string
	price   float64
}

func main() {
	var price []float64
	product := []string{"leite", "oleo"}

	var market [2]string = [2]string{"center box", "cometa"}

	fmt.Println(market)
	price = append(price, 10.3)
	product = append(product, "rapadura")
	fmt.Printf("product - %v price - %v\n", product[0], price)

	newArray := product[:1]
	fmt.Printf("new array - %v , size - %v - capacity - %v\n", newArray, len(newArray), cap(newArray))

	a := []Product{
		{
			id:      1,
			product: "rapaduta",
			price:   5.60,
		},
		{
			id:      2,
			product: "leite",
			price:   5.00,
		},
	}
	fmt.Println(a[0])

	newProduct := Product{
		id:      3,
		product: "Farinha",
		price:   6.70,
	}

	a = append(a, newProduct)
	fmt.Println(a)

}
