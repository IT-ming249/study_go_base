/*
Go语言中的基础数据类型可以表示一些事物的基本属性，
但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了，
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
这种数据类型叫结构体，英文名称struct

	type 类型名 struct{
		字段名 字段类型
		字段名 字段类型
		...
	}
*/
package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
} //字段间分隔不用逗号

func main() {
	//结构体实例化
	var p1 Person
	p1.name = "张三"
	p1.age = 20
	p1.gender = "男"
	fmt.Printf("值:%v,类型:%T\n", p1, p1)
	fmt.Printf("值:%#v,类型:%T\n", p1, p1) // %#v完整打印结构体的信息

}
