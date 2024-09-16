/*


结构体可以描述现实生活中的任何事物，生活中的任何事物都可以当做结构体对象

我们可以把客观事物封装成结构体 :

	汽车

	 	汽车有属性：颜色   大小   重量  发动机   轮胎 ...

	 	汽车行为 也叫方法：  run  跑

	小狗

		属性：颜色     大小     品种  性别 ..

		行为：叫    、闻一闻  舔一舔  咬一咬

	电风扇

		属性： 颜色 大小  高低...

		行为：转动

	人也是一个结构体对象：

		属性： 名字  年龄 性别 ...

		行为：工作  运动 ..


*/

package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
	height int
}

func (p Person) Printinfo() { //不修改结构体类型的属性，采用非指针类型接收者
	fmt.Printf("姓名：%v,年龄：%v\n", p.Name, p.Age)
}

func (p *Person) Setinfo(name string, age int) { //指针类型的接收者，接收者是指针类型才能修改实例的属性
	p.Name = name
	p.Age = age
}
func main() {
	var p1 = Person{
		Name:   "张三",
		Age:    20,
		Gender: "男",
		height: 180,
	}
	p1.Printinfo() //调用结构体方法

	p1.Setinfo("李四", 18)
	p1.Printinfo()

}
