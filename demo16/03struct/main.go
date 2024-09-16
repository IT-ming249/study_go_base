/*

Go语言中的基础数据类型可以表示一些事物的基本属性，
但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了，
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
这种数据类型叫结构体，英文名称struct

注意：结构体首字母可以大写也可以小写，大写表示这个结构体是公有的，在其他的包里面 可以使用。小写表示这个结构体是私有的，只有这个包里面才能使用
*/

package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	//实例化方法2
	var p2 = new(Person) //这个结构体是个结构体指针
	p2.Name = "李四"
	p2.Age = 20
	p2.Gender = "女"
	fmt.Printf("值:%v,类型:%T\n", p2, p2)
	fmt.Printf("值:%#v,类型:%T\n", p2, p2) // %#v完整打印结构体的信息
	//注意：在 Golang 中支持对结构体指针直接使用.来访问结构体的成员。p2.name = "张三" 其 实在底层是(*p2).name = "张三"

	//实例化方法3
	var p3 = &Person{} //直接将结构体地址赋值,也是个指针类型，与方法2相似
	p3.Name = "王五码子"
	p3.Age = 10
	p3.Gender = "女"
	fmt.Printf("值:%v,类型:%T\n", p3, p3)
	fmt.Printf("值:%#v,类型:%T\n", p3, p3) // %#v完整打印结构体的信息

	//实例化方法4
	var p4 = Person{
		Name:   "朱八八",
		Age:    20,
		Gender: "男",
	}
	fmt.Printf("值:%v,类型:%T\n", p4, p4)
	fmt.Printf("值:%#v,类型:%T\n", p4, p4) // %#v完整打印结构体的信息

	//实例化方法5
	var p5 = &Person{
		Name: "朱五四",
		Age:  20,
	} //不给某个字段赋值则使用默认值
	fmt.Printf("值:%v,类型:%T\n", p5, p5)
	fmt.Printf("值:%#v,类型:%T\n", p5, p5) // %#v完整打印结构体的信息

	//实例化方法6
	var p6 = &Person{
		"大老刘",
		20,
		"男",
	} //不说明键则需要按照定义结构体的顺序来
	fmt.Printf("值:%v,类型:%T\n", p6, p6)
	fmt.Printf("值:%#v,类型:%T\n", p6, p6) // %#v完整打印结构体的信息

}
