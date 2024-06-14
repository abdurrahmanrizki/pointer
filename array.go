package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	//cara 1
	//var arr1 [5]int
	//arr1 = [5]int{10, 20, 30, 40, 50}
	//fmt.Println(arr1)

	//cara 1 (gabung)
	var arr1 [5]int = [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr1)
	//println(len(arr1), cap(arr1))

	mySlice := arr1[1:4]
	fmt.Println(mySlice)

	total := sum(mySlice...)
	fmt.Println("Sum of slice elements:", total)

	arr1[2] = 100
	fmt.Println(arr1)

	p := &arr1[2]
	*p = 100
	fmt.Println("Array setelah di ubah melalui pointer :", arr1)

	//println(len(mySlice), cap(mySlice))

	//mySlice [1] = 5
	//mySlice =append(mySlice, 8, 9)//kapasitas melebihi array1, rujukan bukan array lagi
	//fmt.Println(arr1)
	//fmt.Println(mySlice)

	//var arr2 [] int
	//arr2 = append(arr2, 1, 2, 3, 4, 5, 6, 7, 8)
	//var arr3 [] int = [] int {11, 23, 45}
	//arr2 = append(arr2, arr3...)
	//fmt.Println(arr2, arr3)

}
