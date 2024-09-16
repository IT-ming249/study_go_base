package main

import "fmt"

func main() {
	/*
		for 初始语句;条件表达式;结束语句{
			循环体语句
		}
	*/

	// 1、打印1-10的所有数据
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
	}

	// 2、   for 循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	fmt.Println()
	j := 1
	for ; j <= 10; j++ {
		fmt.Print(j)
	}
	/*
		3、打印1-10的所有数据

		写for循环的时候要注意死循环

		for 条件 {
			循环体语句
		}

		for 循环的初始语句和结束语句都可以省略，例如
	*/
	fmt.Println()
	k := 1
	for k <= 10 {
		fmt.Print(k)
		k++ //初始语句和结束语句省略是要注意死循环
	}
	fmt.Println()
	/*
		4、打印1-10的所有数据

		for {
			循环体语句
		}
		注意：Go 语言中是没有 while 语句的，我们可以通过 for 代替
	*/
	//下面这样可以达到与while同样的效果
	for { //这样是一个死循环
		if k <= 10 {
			fmt.Print(k)
		} else {
			break //需要添加停止条件跳出死循环
		}
		k++
	}
}
