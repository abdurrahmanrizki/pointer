package main

type customeFunc func(int, int) bool
type operationFunc func(int, int) int

func myCallback(f operationFunc) {
	println(f(4, 3))
}

func main() {
	println(add(2, 3))
	println(add(5, 4))
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
