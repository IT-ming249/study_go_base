package main

import "fmt"

func main() {

	/*
		Go 语言中 if 条件判断的格式如下：
		if 表达式1 {
			分支 1
		}else if 表达式2 {
			分支 2
		}else{
			分支 3
		}
	*/

	//1、最简单的if语句
	flag := true
	if flag {
		fmt.Println(flag)
	}
	//2、if语句的另一种写法

	//3、探讨上面两种写法的区别

	//4、输入一个人的成绩，如果成绩大于等于90输出A，如果小于90大于75输出B,否则输出C

	if score := 95; score >= 95 { //复制可以在if里面
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//5、if else要注意的细节

	//5.1、if{}不能省略掉

	//5.2、 { 必须紧挨着条件

	//6、求两个数的最大值
	var a = 10
	var b = 200
	var max int
	if a > b {
		max = a
		fmt.Printf("最大值%v", max)
	} else {
		max = b
		fmt.Printf("最大值%v", max)
	}
}
