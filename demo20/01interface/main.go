package main

import "fmt"

//接口是一个规范
/*
type 接口名 interface{
	方法名1 （参数列表1） 返回值列表1
	方法名2 （参数列表2） 返回值列表2
	...
}
*/

type Usber interface {
	start() //可以规定参数和返回值的类型
	stop()
}

//如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

//手机要实现usb接口的话必须得实现usb接口中的所有方法，即手机要定义start和stop方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
	Name string
}

func (c Camera) start() {
	fmt.Println("相机启动")
}
func (c Camera) stop() {
	fmt.Println("相机关机")
}
func (c Camera) run() {
	fmt.Println("咔嚓")
}

func main() {
	p := Phone{
		Name: "红米k50",
	}
	p.start()

	//实现接口
	var p1 Usber //golang中的接口就是一个类型
	p1 = p       //表示手机实现Usb接口

	p1.start()
	p1.stop()

	c := Camera{
		Name: "哈苏",
	}
	//实现相机接口
	var c1 Usber
	c1 = c //表示相机实现了Usb接口
	c1.start()
	c1.stop()

	//接口不能调用结构体除接口方法外的方法
	//c1.run()  （×）

	c.run()
}
