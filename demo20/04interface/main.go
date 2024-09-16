package main

import "fmt"

//golang中空接口也可以直接当做类型来使用，可以表示任意类型

//空接口作为函数的参数
func show(b interface{}) { //若函数项以任意类型作为参数，则可以定义该参数为空接口
	fmt.Printf("值：%v,类型%T\n", b, b)
}

func main() {
	var a interface{}
	a = 20
	fmt.Printf("值：%v,类型%T\n", a, a)
	a = "hel"
	fmt.Printf("值：%v,类型%T\n", a, a)
	a = true
	fmt.Printf("值：%v,类型%T\n", a, a)

	show(20)
	show("hello")
	show([]int{1, 2, 3, 4, 5})

	//空接口在map和切片类型中也有大用处
	m1 := make(map[string]interface{}) //表示这个map的值可以是任意类型
	m1["username"] = "张三"
	m1["age"] = 20
	m1["isactive"] = true
	fmt.Println(m1)

	var slc = []interface{}{"空接口", true, 12}
	fmt.Println(slc)

}
