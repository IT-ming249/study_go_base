package main

import "fmt"

func main() {
	// 1、请求出一个数组里面元素的和以及这些元素的平均值。分别用for和for-range实现
	var intArr = [...]int{1, -1, 12, 65, 11}
	//for 方法
	sum := 0
	for i := 0; i < len(intArr); i++ {
		sum = sum + intArr[i]
	}
	avg_for := sum / len(intArr)
	fmt.Printf("for方法求出平均值%v\n", avg_for)

	sum_r := 0
	for _, v := range intArr {
		sum_r = sum_r + v
	}
	avg_for_range := sum_r / len(intArr)
	fmt.Printf("for_range方法求出平均值%v\n", avg_for_range)
	/*
		2、请求出一个数组的最大值，并得到对应的下标

				思路
					1、 声明一个数组 var intArr = [...]int {1, -1, 12, 65, 11}
					2、假定第一个元素就是最大值，下标就 0
					3、然后从第二个元素开始循环比较，如果发现有更大，则交换
	*/
	var max = intArr[0]
	for j := 1; j < len(intArr); j++ {
		if intArr[j] > max {
			max = intArr[j]
		}
	}
	fmt.Printf("intArr中的最大值为%v\n", max)

	//3、从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	var arr = [...]int{1, 3, 5, 7, 8}
	for a := 0; a < len(arr); a++ {
		for b := a + 1; b < len(arr); b++ {
			if arr[a]+arr[b] == 8 {
				fmt.Printf("和为8的元素下标(%v,%v)\n", a, b)
			}
		}
	}
}
