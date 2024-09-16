package main

import "fmt"

type Usber interface { //接口本身就是一个类型
	start() //可以规定参数和返回值的类型
	stop()
}

// 电脑
type Computer struct {
	Name string
}

//类型断言也可以用于判断结构体，不局限于基本数据类型
func (c Computer) work(usb Usber) {

	usb.start()
	usb.stop()
}

// 手机
type Phone struct {
	Name string
}

func (p Phone) start() { //这种就是值接收者
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

// 照相机
type Camera struct {
	Name string
}

func (c *Camera) start() { //这种就是指针接收者
	fmt.Println("相机启动")
}
func (c *Camera) stop() {
	fmt.Println("相机关机")
}
func (c Camera) run() {
	fmt.Println("咔嚓")
}

/*
结构体值接收者实现接口：

值接收者： 如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量

*/
func main() {

	// 结构体值接收者实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
	var rmk50 = Phone{
		Name: "红米",
	}

	var p1 Usber = rmk50 // 表示让红米手机实现Usb接口
	p1.start()

	var iphone = &Phone{
		Name: "苹果",
	}

	var p2 Usber = iphone
	p2.start()

	//结构体方法是指针类型接收者只接受指针类型结构体赋值给接口变量

	var hasu = &Camera{
		Name: "哈苏",
	}
	var p3 = hasu
	p3.start()
}
