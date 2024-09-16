// 逻辑运算
package main

import "fmt"

//定义一个方法
func test() bool {
	fmt.Println("test...")
	return true
}
func main() {

	/*
		&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。
		||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。
		!	逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。
	*/
	var a = 23
	var b = 8
	fmt.Println(a > 10 && b < 10)
	fmt.Println(a > 24 && b < 10)
	fmt.Println(a > 24 || b < 10)
	fmt.Println(!false)
	fmt.Println(a != b)
	fmt.Println()
	//逻辑与和逻辑或短路
	//逻辑与：前面是false后面就不会执行
	if a > 9 && test() {
		fmt.Println("执行与")
	}
	if a > 100 && test() {
		fmt.Println("执行与")
	}

	//逻辑或：前面是true后面就不会执行
	if b > 10 || test() {
		fmt.Println("执行或")
	}
	if b < 10 || test() {
		fmt.Println("执行或")
	}
}
