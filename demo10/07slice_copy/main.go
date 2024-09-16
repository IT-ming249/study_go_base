package main

import "fmt"

func main() {
	/*
		值类型 ： 改变变量副本值的时候，不会改变变量本身的值
		引用类型：改变变量副本值的时候，会改变变量本身的值
	*/

	// 切片就是引用数据类型

	//1、copy()函数复制切片
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := make([]int, 5, 5) //复制切片要创建一个容量和长度都适合的空切片
	copy(sliceB, sliceA)
	fmt.Println(sliceA)
	fmt.Println(sliceB)
	sliceB[0] = 100
	fmt.Println(sliceB)
}
