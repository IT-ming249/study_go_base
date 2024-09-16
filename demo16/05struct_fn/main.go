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
	age    int
	Gender string
	height int
} //属性定义完成
/*添加结构体的自定义方法
func (接收者变量 接收者类型) 方法名（参数列表）（返回参数）{ //接收者类型一般是该结构体的名称
	函数体
}
*/
func (p Person) Printinfo() {
	fmt.Printf("姓名：%v,年龄：%v\n", p.Name, p.age)
}

func main() {
	var p1 = Person{
		Name:   "张三",
		age:    20,
		Gender: "男",
		height: 180,
	}
	p1.Printinfo() //调用结构体方法

	var p2 = Person{
		Name:   "李四",
		age:    20,
		Gender: "女",
		height: 160,
	}
	p2.Printinfo()

}
