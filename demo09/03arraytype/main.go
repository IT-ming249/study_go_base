package main

import "fmt"

func main() {
	/*
		值类型 ：改变变量副本值的时候，不会改变变量本身的值
		引用类型：改变变量副本值的时候，会改变变量本身的值
	*/

	//值类型:基本数据类型 和 数组都是值类型
	var a = 10
	b := a //a后面变了也不会影响b
	a = 20
	fmt.Println(a, b)

	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1 //arr1发生改变不会影响arr2
	fmt.Println(arr1, arr2)
	arr1[0] = 10
	fmt.Println(arr1, arr2)

	//引用类型: 切片
	var ar1 = []int{1, 2, 3} //没有指定长度的数组为切片
	var ar2 = ar1            //切片ar1发生改变会影响切片ar2
	fmt.Println(ar1, ar2)
	ar1[0] = 10
	fmt.Println(ar1, ar2)
	ar2[0] = 100 //切片ar2的值改变也会影响ar1，因为它两指向通通以地址
	fmt.Println(ar1, ar2)

}
