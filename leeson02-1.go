//编写一个函数，接收一个整数数组，返回一个map
// 其中键是数组中的元素，值是该元素在数组中出现的次数。

package main

import "fmt"

func countOccurrences(arr []int) map[int]int {
	countMap := make(map[int]int)

	for _, num := range arr {
		countMap[num]++
	}
	return countMap
}

func main() {
	arr := []int{2006, 1, 17, 1, 7, 11, 1, 17, 2006, 1}
	result := countOccurrences(arr)
	fmt.Println(result)
}
