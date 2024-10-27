//编写一个Go函数，接受圆的半径作为参数，然后返回圆的面积。使用 math 包中的常数 Pi。在 main 函数中调用此函数并打印结果。
//提示，引入 Pi 只需要写出math.Pi

package main

//var r float64

//func main() {

//fmt.Scan(&r)

//pi := math.Pi
//area := pi * r * r

//fmt.Printf("area is %f", area)

//}

//模块化写法

import (
	"fmt"
	"math"
)

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
