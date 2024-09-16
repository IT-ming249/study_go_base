package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//1、定义int类型
	var n1 int = 10 //64位计算机默认int类型为int64占8字节
	fmt.Printf("%v 类型%T\n", n1, n1)

	//2、int8的范围演示
	var n2 int8 = 98
	//n2=130 130超出int8的数据范围了，会报错
	fmt.Printf("%v 类型%T\n", n2, n2)
	fmt.Println("int8占用存储空间 ", unsafe.Sizeof(n2)) //单位是字节 8bit一字节

	//3、uint8的范围(0-255)
	var n3 uint8 = 127
	// n3=-1（×） 负数超出无符号数的范围
	fmt.Printf("%v 类型%T\n", n3, n3)
	fmt.Println("unit8占用存储空间 ", unsafe.Sizeof(n3))

	//5、int不同长度类型转换
	var a1 int32 = 10
	var a2 int64 = 21
	//a1+a2 不同长度直接相加会报错
	fmt.Println(int64(a1) + a2) //强制类型转换
	//高位向低位转换的时候要注意防止溢出 eg:int64 1000 => int8(1000)肯定报错

	//6、数字字面量语法  %d 表示10进制输出 %b表示二进制输出 %o 八进制输出 %x 表示16进制输出(了解)
	num := 12

	fmt.Printf("num=%v\n", num) //%v 原样输出

	fmt.Printf("num=%d\n", num) //%d 表示10进制输出

	fmt.Printf("num=%b\n", num) //%b 表示二进制输出

	fmt.Printf("num=%o\n", num) //%o 表示八进制输出

	fmt.Printf("num=%x\n", num) //%x 表示16进制输出

}
