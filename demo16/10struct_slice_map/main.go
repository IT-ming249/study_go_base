package main

import "fmt"

/*

结构体的字段类型可以是：基本数据类型、也可以是切片、Map 以及结构体

如果结构体的字段类型是: 指针，slice，和map的零值都是 nil ，即还没有分配空间

如果需要使用这样的字段，需要先make，才能使用.

*/
type Person struct {
	name  string
	age   int
	hobby []string
	map1  map[string]string
}

func main() {
	var p Person
	p.name = "张三"
	p.age = 20
	p.hobby = make([]string, 6, 6)
	p.hobby[0] = "写代码"
	p.hobby[2] = "打篮球"
	p.map1 = make(map[string]string)
	p.map1["address"] = "北京"
	p.map1["phone"] = "12345"
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v\n", p.hobby)

}
