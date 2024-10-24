//编写一个Go函数，接受两个整数作为参数，然后返回它们的和。在 main 函数中调用此函数并打印结果。

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
