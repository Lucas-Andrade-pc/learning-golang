package recursion

func FactorialRecursion(number int) int {
	if number == 0 {
		return 1
	}
	return number * FactorialRecursion(number-1)

}
