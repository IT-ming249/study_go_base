package main

import "fmt"

func main() {

	//1、练习：打印 0-50 所有的偶数
	i := 0
	for ; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
	fmt.Println()
	//2、练习：求 1+2+3+4 +...100 的和
	sum := 0
	for j := 1; j <= 100; j++ {
		sum = sum + j
	}
	fmt.Println(sum)
	//3、练习：打印 1~100 之间所有是 9 的倍数的整数的个数及总和
	count := 0
	sum9 := 0
	for k := 0; k <= 100; k++ {
		if k%9 == 0 {
			count++
			sum9 = sum9 + k
		}
	}
	fmt.Println(count, sum9)

	//4、练习：计算 5 的阶乘 (12345 n 的阶乘 12……n)   1*2*3*4*5
	sump := 0
	for y := 1; y <= 5; y++ {
		if y == 1 {
			sump = y
		} else {
			sump = sump * y
		}
	}
	fmt.Printf("5的阶乘%v\n", sump)
	/*
		1、sum=1*1
		2、sum=sum*2   1*1*2
		3、sum=sum*3   1*1*2*3
		4、sum=sum*4   1*1*2*3*4
		5、sum=sum*5   1*1*2*3*4*5
	*/

	//5、练习： 打印一个矩形 (for循环的嵌套)
	i1 := 0
	for {
		if i1 < 4 {
			fmt.Println("****")
		} else {
			break
		}
		i1++
	}
	fmt.Println()
	/*

	****
	****
	****

	 */

	//6、for循环的嵌套

	/*
		for循环嵌套的一个执行流程
		1、i=0   打印4个*   一个换行
		2、i=1   打印4个*   一个换行
		3、i=2   打印4个*   一个换行
		4、i=3  跳出循环
	*/

	//7、练习： 打印一个三角形
	for i3 := 1; i3 <= 9; i3++ {
		for j3 := 1; j3 <= i3; j3++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	/*

	*
	**
	***
	****
	*****

	 */

	// 8、练习：打印出九九乘法表
	for i2 := 1; i2 <= 9; i2++ {
		for j2 := 1; j2 <= i2; j2++ {
			fmt.Printf("%v×%v=%v ", i2, j2, i2*j2)
		}
		fmt.Println()
	}

}
