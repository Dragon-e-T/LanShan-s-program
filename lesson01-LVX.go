//编写一个Go函数，使用rand包随机选择一个1-100的数（必须每次执行的随机数都不一样），然后使用二分法找到这个数。
//（电脑出题电脑做）

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return binarySearch(arr, target)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i + 1
	}

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf(" Got it! %d!\n", result)
	}

}
