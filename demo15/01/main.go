package main

import "fmt"

func main() {
	//在计算机底层 a 这个变量其实对应了一个内存地址
	var a = 10
	fmt.Printf("值%v,类型%T,地址%p\n", a, a, &a) //&取地址

	// 指针也是一个变量，但它是一种特殊的变量，它存储的数据不是一个普通的值，而是另一个变量的内存地址
	//golang中每一个变量都有自己的内存地址，指针变量也有
	var p *int = &a //指针变量
	fmt.Printf("值%v,类型%T,地址%p\n", p, p, &p)

}
