/*
	值类型 ： 改变变量副本值的时候，不会改变变量本身的值 (数组、基本数据类型、结构体)
	引用类型：改变变量副本值的时候，会改变变量本身的值  （切片、map）
*/

package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	var p1 Person
	p1.Name = "嘿嘿"
	p1.Age = 20
	p1.Gender = "男"

	var p2 = p1
	fmt.Println(p1, p2)

	p2.Name = "哈哈"

	fmt.Println(p1, p2) //结构体是值类型
}
