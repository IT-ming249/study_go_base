package main

import "fmt"

func main() {
	//匿名函数  匿名自执行函数
	func() {
		fmt.Println("test")
	}()

	//匿名函数 赋值给变量
	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Print(fn(2, 3))

	// 匿名自执行函数接收参数
	func(x, y int) int {
		return x + y
	}(10, 2)

}
