package main

import "fmt"

//一个结构体实现多个接口
type Animaler interface {
	SetName(string)
}

type Animaler2 interface {
	GetName() string
}

//狗
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

	var d1 Animaler = d //Dog实现Animal的接口

	var d2 Animaler2 = d //Dog实现Animal2的接口
	fmt.Println(d2.GetName())
	d1.SetName("哈士奇")
	fmt.Println(d2.GetName())

}
