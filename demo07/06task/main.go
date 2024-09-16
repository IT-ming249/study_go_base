// 位运算符
package main

import "fmt"

func main() {
	//练习1：有两个变量，a和b，要求将其进行交换，最终打印结果
	a := 10
	b := 20
	fmt.Printf("交换前a=%v,b=%v\n", a, b)
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("交换后a=%v,b=%v\n", a, b)

	//练习2：有两个变量，c和d，要求将其进行交换(不能使用中间变量)，最终打印结果
	var c = 11
	var d = 22
	fmt.Printf("交换前c=%v,d=%v\n", c, d)
	c = c + d
	d = c - d
	c = c - d
	fmt.Printf("交换前c=%v,d=%v\n", c, d)

	// 练习3：假如还有100天放假，问：xx个星期零xx天
	days := 100
	weeks := days / 7
	lat_days := 100 % 7
	fmt.Printf("离放假还有%v周零%v天\n", weeks, lat_days)
	/*
	 练习4：定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：
	 C=(F-32)÷1.8,摄氏温标（°C）和华氏温标（°F），请求出华氏温度对应的摄氏温度
	*/

	var F = 100.0
	var C = (F - 32) / 1.8
	fmt.Printf("华氏%v度对应摄氏%.2f度", F, C)
}
