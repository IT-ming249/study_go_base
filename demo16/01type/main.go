package main

import "fmt"

//自定义类型
type myInt int

//类型别名
type myFloat = float64

func main() {
	var a myInt = 10
	fmt.Printf("%v,%T\n", a, a)
	var b myFloat = 12.56
	fmt.Printf("%v,%T\n", b, b) //打印出来的类型还是float64
}
