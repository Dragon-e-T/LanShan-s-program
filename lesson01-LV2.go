package main

import (
	"fmt"
	"math"
)

//编写一个Go函数，接受圆的半径作为参数，然后返回圆的面积。使用 math 包中的常数 Pi。在 main 函数中调用此函数并打印结果。
//提示，引入 Pi 只需要写出math.Pi

func Multiply(r float64) float64 {
	pi := math.Pi
	return pi * r * r
}
func main() {

	var r float64
	fmt.Scan(&r)

	area := Multiply(r)
	fmt.Println(area)
}
