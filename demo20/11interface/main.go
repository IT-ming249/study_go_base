package main

import "fmt"

type A interface {
	SetName(string)
}

type B interface {
	GetName() string
}

// 接口嵌套
type Animaler interface { //要实现Animaler 就要同时实现接口A和B中的所有方法
	A
	B
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(d_name string) { //需要修改结构体属性内容的，采用指针接收者
	d.Name = d_name
}
func (d Dog) GetName() string {
	return d.Name
}
func main() {
	d := &Dog{
		Name: "旺旺",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("哈哈")
	fmt.Println(d1.GetName())

	// 接口只定义不实现，是个规范
}
