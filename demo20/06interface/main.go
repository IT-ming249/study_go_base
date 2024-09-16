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
	if _, ok := usb.(Phone); ok {
		fmt.Println("欢迎手机用户")
	} else if _, ok := usb.(Camera); ok {
		fmt.Println("欢迎照相机用户")
	} else {
		fmt.Println("请接入正确接口")
	}
	usb.start()
	usb.stop()
}

// 手机
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

// 照相机
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
	var iphone = Phone{
		Name: "苹果",
	}
	var ausu = Computer{
		Name: "华硕",
	}
	var hasu = Camera{
		Name: "哈苏",
	}
	ausu.work(iphone)
	ausu.work(hasu)
}
