// 关系运算
package main

import "fmt"

func main() {

	/*
		==	检查两个值是否相等，如果相等返回 True 否则返回 False。
		!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。
		>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。
		>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
		<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。
		<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
	*/
	//演示关系运算符的使用
	var a1 = 9
	var a2 = 8
	fmt.Println(a1 == a2) //false
	fmt.Println(a1 != a2) //true
	fmt.Println(a1 > a2)  //true
	fmt.Println(a1 >= a2) //true
	fmt.Println(a1 < a2)  //flase
	fmt.Println(a1 <= a2) //flase
	flag := a1 > a2
	fmt.Println("flag=", flag)

	var b1 = 9
	var b2 = 14
	flag1 := b1 < b2
	if flag1 {
		fmt.Println("b1<b2")
	}
}
