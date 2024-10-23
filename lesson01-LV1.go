package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := add(a, b)
	fmt.Println(sum)
}
