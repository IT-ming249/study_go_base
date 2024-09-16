package main

import (
	"fmt"
	"strconv"
)

func main() {

	//1、通fmt.Sprintf() 把其他类型转换成 String 类型
	var a int64 = 20
	var b float32 = 12.56
	var c bool = true
	var d byte = 'a'

	s1 := fmt.Sprintf("%d", a) //"%d" 是一个格式化占位符，用于将整数转换为十进制数字字符串
	fmt.Printf("值%v--类型%T\n", s1, s1)

	s2 := fmt.Sprintf("%f", b)
	fmt.Printf("值%v--类型%T\n", s2, s2)

	s3 := fmt.Sprintf("%t", c)
	fmt.Printf("值%v--类型%T\n", s3, s3)

	s4 := fmt.Sprintf("%c", d)
	fmt.Printf("值%v--类型%T\n", s4, s4)
	//注意：Sprintf 使用中需要注意转换的格式 int 为%d  float 为%f  bool 为%t   byte 为%c

	//2、通过strconv  把其他类型转换成string类型

	/*
		FormatInt
		参数1：int64 的数值
		参数2：传值int类型的进制
	*/
	fmt.Println()
	st1 := strconv.FormatInt(a, 10) //FormatInt只接受int64类型的int进行转换
	fmt.Printf("值%v--类型%T\n", st1, st1)

	st2 := strconv.FormatFloat(float64(b), 'f', 3, 64)
	fmt.Printf("值%v--类型%T\n", st2, st2)
	/*
		参数 1：要转换的值
		参数 2：格式化类型 'f'（-ddd.dddd）、
			 'b'（-ddddp±ddd，指数为二进制）、
			 'e'（-d.dddde±dd，十进制指数）、
			 'E'（-d.ddddE±dd，十进制指数）、
			 'g'（指数很大时用'e'格式，否则'f'格式）、
			 'G'（指数很大时用'E'格式，否则'f'格式）。
		 参数 3: 保留的小数点 -1（不对小数点格式化）
		 参数 4：格式化的类型 传入 64  32
	*/

	st3 := strconv.FormatBool(c) //根本没有意义
	fmt.Printf("值%v--类型%T\n", st3, st3)

	//字符型转字符串也没啥意义

}
