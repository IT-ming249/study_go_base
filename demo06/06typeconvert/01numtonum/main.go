package main

import "fmt"

func main() {

	//1、整型和整型之间的转换
	var a int8 = 20
	var a1 int16 = 40
	fmt.Println(int16(a) + a1) //数据类型不同没法相加，需要转换成同类型

	//2、浮点型和浮点型之间的转换
	var b float32 = 20.1
	var b1 float64 = 30.2
	fmt.Println(b + float32(b1))

	//3、整型和浮点型之间的转换,不建议把浮点型转换成整型
	var c int = 10
	var c1 float32 = 20.1
	fmt.Print(float32(c) + c1)
	//注意：转换的时候建议从低位转换成高位，高位转换成低位的时候如果转换不成功就会溢出，和我们想的结果不一样。

}
