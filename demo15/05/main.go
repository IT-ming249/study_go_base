package main

import "fmt"

func main() {

	// 实际想开发中 new 函数不太常用，使用 new 函数得到的是一个指类型针，并且该指针对应的值为该类型的零值

	/*
		错误的写法
		var a *int //指针也是引用数据类型
		*a = 100
		fmt.Println(*a)
	*/
	a := new(int) //a是一个指针变量 类型是*int 对应的值为0（int类型默认值）
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)

	var c *int
	c = new(int) //给指针变量c分配存储空间
	*c = 100
	fmt.Println(*c)

	//make也可以用来分配内存空间，之前用于slice,map的创建，这两都是引用数据类型

	//区别：new返回的是指针，make返回的是引用类型本身

	//Golang指针千万记住：&（取地址） *（根据地址取值）
}
