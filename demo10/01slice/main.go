package main

import "fmt"

func main() {
	/*1、回顾数组的声明以及初始化*/
	// var arr = [3]int{1, 2, 3}
	// arr1 := [3]int{4, 5, 6}
	// arr2 := [...]int{1, 2, 3, 4}
	// fmt.Println(arr, arr1, arr2)
	/*2、切片的声明 初始化*/ //切片：用有相同类型元素的可变长度的序列
	//声明方法1：
	var slc []int
	fmt.Printf("空切片值:%v,类型:%T,长度:%v\n", slc, slc, len(slc))
	fmt.Println(slc == nil) //golang中声明切片后，切片的默认值是nil
	//声明方法以及初始化2：
	var slc1 = []int{1, 2, 3, 4}
	fmt.Printf("切片值:%v,类型:%T,长度:%v\n", slc1, slc1, len(slc1))
	//声明方法以及初始化：
	var slc2 = []int{0: 1, 2: 2, 4: 3, 6: 4}
	fmt.Printf("切片值:%v,类型:%T,长度:%v\n", slc2, slc2, len(slc2))

}
