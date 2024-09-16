package main

import "fmt"

//函数作为另一个函数参数
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//自定义一个方法类型
type calc func(int, int) int

func clc(x, y int, cb calc) int {
	return cb(x, y)
}
func main() {
	fmt.Println(clc(10, 5, add))

	fmt.Println(clc(10, 5, sub))

	j := clc(3, 4, func(x, y int) int { //可以采用匿名函数的方式
		return x * y
	})

	fmt.Println(j)
}
