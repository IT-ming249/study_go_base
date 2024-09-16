package main

import "fmt"

/*
panic/recover

Go 语言中目前（Go1.12）是没有异常机制，但是使用 panic/recover 模式来处理错误。

panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效
*/

func fn1(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}() //必须是个匿名自执行函数
	return a / b
}

func main() {

	fmt.Println(fn1(10, 0))
	fmt.Println("结束")
	fn1(10, 2)
}
