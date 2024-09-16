package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1、定义float类型
	var a float32 = 3.14
	fmt.Printf("值%v 类型%T\n", a, a)
	fmt.Println("float32占用内存", unsafe.Sizeof(a))

	//2、 %f  输出float类型   %.2f 输出数据的时候保留2位小数 保留n位%.4f
	fmt.Printf("%f,%.2f\n", a, a)

	//3、 64位的系统中 浮点数默认是 float64
	var b = 3.33
	fmt.Printf("值%f 类型%T\n", b, b)
	fmt.Println("float64占用内存", unsafe.Sizeof(b))

	//4.Golang中使用科学计数法表示浮点数
	var c = 3.14e2  //3.14 *10^2
	var d = 3.14e-2 //3.14 /10^2
	fmt.Printf("值%v 类型%T\n", c, c)
	fmt.Printf("值%v 类型%T\n", d, d)

	// 5、Golang 中 float 精度丢失问题 可以使用第三方包解决
	e := 1129.6
	fmt.Println(e * 100) //112959.99999999999

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) //4.3999999999999995

	// 6、int类型转换成float类型
	m := 10
	fmt.Println("m的类型：", unsafe.Sizeof(m))
	fmt.Printf("m转换后的类型%T", float32(m))

	//7. float32也能转换为float64

	//8、float类型转换成int类型 （不建议这样做）

	var c1 float32 = 23.45

	c2 := int(c1)

	fmt.Printf("\nc2的值：%v  c2的类型：%T", c2, c2)

}
