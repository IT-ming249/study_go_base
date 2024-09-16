package main

import "fmt"

//定义函数类型 type 类型名 func(int int) int
type calc func(int, int) int //表示定义一个calc的类型

type myInt int //自定义一个叫myInt的类型

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test...")
}

func main() {
	var c calc
	c = add // 方法赋给c
	fmt.Printf("c的类型%T\n", c)

	//c = test()  该方法必须满足func(int, int) int，才能复制给c

	fmt.Println(c(10, 5)) //c的类型main.calc

	//类型推导方式
	f := sub
	fmt.Printf("f的类型%T\n", f) //f的类型func(int, int) int

	var a int = 10
	var b myInt = 20
	fmt.Printf("a的类型%T,b的类型%T\n", a, b)
	// fmt.Println(a + b) 不同类型不能相加
}
