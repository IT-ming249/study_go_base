package main

import "fmt"

//定义一个接口
type Usber interface {
	start() //可以规定参数和返回值的类型
	stop()
}

//电脑
type Computer struct {
	Name string
}

func (c Computer) work(usb Usber) {
	fmt.Println("欢迎使用")
	usb.start()
	usb.stop()
}

//手机
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

//照相机
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

	var computer = Computer{
		Name: "联想",
	}

	var phone = Phone{
		Name: "红米K50",
	}

	var camera = Camera{
		Name: "哈苏",
	}

	//手机使用usb接口连接电脑

	computer.work(phone)  //手机实现了Usb接口，用于电脑端的work方法
	computer.work(camera) //照相机实现了Usb接口

	//Golang的接口也是一种数据类型，不需要显式实现，只需要一个变量含有接口类型的所有方法，这个变量就实现了这个接口

}
