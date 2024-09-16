package main

import "fmt"

//空接口  表示没有任何约束  任意的类型都可以实现空接口
type Null_jk interface{}

func main() {
	var a Null_jk
	var str = "hello"
	a = str //让字符串实现Null_jk 接口
	fmt.Printf("值：%v,类型%T\n", a, a)

	var num = 10
	a = num //整型实现Null_jk接口
	fmt.Printf("值：%v,类型%T\n", a, a)
}
