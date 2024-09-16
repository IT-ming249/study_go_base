package main

import "fmt"

//定义一个Animal的接口，Animal中定义两个方法，分别是SetName和GetName。分别让Dog结构体和Cat结构体实现这个方法
type Animaler interface {
	SetName(string)
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

//猫
type Cat struct {
	Name string
}

func (c *Cat) SetName(c_name string) { //接口方法的参数个数要和定义接口时一样
	c.Name = c_name
}
func (c Cat) GetName() string {
	return c.Name
}

func main() {

	//Dog实现Animal的接口
	var hasqi = &Dog{ //指针类型接收者
		Name: "哈士奇",
	}
	var a1 Animaler = hasqi
	fmt.Println(a1.GetName())
	a1.SetName("阿奇")
	fmt.Println(a1.GetName())

}
