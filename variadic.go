package main

func square(n int) int {
	return n * n
}

func process(cb func(int) int, nums ...int) { //variadic parameter
	for _, a := range nums {
		println(cb(a))
	}
}
func main() {
	mySlice := []int{4, 5, 6, 7}
	process(square, mySlice...)

	/*num1 := 1
	num2 := 2
	num3 := 3
	num4 := 4
	process (square, num1, num2, num3, num4)*/
}
