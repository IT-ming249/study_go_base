// 赋值运算
package main

import "fmt"

func main() {
	/*	跟Python那些一样的
		=	简单的赋值运算符，将一个表达式的值赋给一个左值
		+=	相加后再赋值
		-=	相减后再赋值
		*=	相乘后再赋值
		/=	相除后再赋值
		%=	求余后再赋值
	*/
	var a = 20
	a = 21
	fmt.Println(a)
	var b = 23 + 2 //这条赋值是从后往前执行
	fmt.Println(b)
	c := b
	fmt.Println(c)
	e := b + c
	fmt.Println(e)

	fmt.Println()

}
