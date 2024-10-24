//编写一个Go函数，接受一个整数作为参数，然后判断它是否为素数（质数）。
//在 main 函数中调用此函数并打印结果。
//提示：一个素数是只能被 1 和自身整除的正整数。

package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {

	var num int
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf(":%d is prime\n\n", num)
	} else {
		fmt.Printf("%d is not prime\n", num)
	}
}
