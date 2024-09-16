package main

import "fmt"

func main() {
	/*
		Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

		注意：
		1.布尔类型变量的默认值为false。
		2.Go 语言中不允许将整型强制转换为布尔型.
		3.布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/

	var flag bool = true
	fmt.Printf("%v--%T\n", flag, flag)

	//1.布尔类型变量的默认值为false
	var debool bool
	fmt.Printf("%v--%T\n", debool, debool)

	//2.string类型变量的默认值是空
	var s string
	fmt.Printf("%v--%T\n", s, s)

	//3.int类型变量默认值是0
	var i int
	fmt.Printf("%v--%T\n", i, i)

	//4.float类型变量默认值是0
	var f float32
	fmt.Printf("%v--%T\n", f, f)

	//5.Go 语言中不允许将整型强制转换为布尔型.
	/* 这是错误写法
	var a = 1
	if a {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	*/

	var a = true
	if a {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	//go中 bool型就是bool型，不参与计算，不接受转换

}
