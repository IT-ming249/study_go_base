package main

import "fmt"

//结构体的继承，通过嵌套结构体实现

//父亲结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v在运动\n", a.Name)
}

//子结构体
type Dog struct {
	Age    string
	Animal //结构体嵌套(继承)
}

func (d Dog) wang() {
	fmt.Printf("%v旺旺叫\n", d.Name)
}

func main() {
	var dog1 Dog
	dog1.Age = "2"
	dog1.Animal.Name = "哈士奇"
	dog1.run() //现在Dog结构体中寻找run方法，找不到就到嵌套结构体Animal中找run方法
	dog1.wang()

}
