package main

import "fmt"

func main() {
	a := 10
	p := &a //p是一个指针变量 类型*int

	fmt.Println(*p) //*p表示取出变量p对应的内存地址的值

	*p = 30
	fmt.Println(a)

}
