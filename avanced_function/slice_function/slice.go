package slice_function

func SumSlice(number []int) int {
	sum := 0
	for _, value := range number {
		sum = sum + value
	}
	return sum
}

func SumSliceOther(number ...int) int {
	sum := 0
	for _, value := range number {
		sum = sum + value
	}
	return sum
}
