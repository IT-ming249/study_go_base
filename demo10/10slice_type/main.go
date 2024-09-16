package main

import "fmt"

func main() {
	/*
		值类型 ： 改变变量副本值的时候，不会改变变量本身的值 (数组、基本数据类型)
		引用类型：改变变量副本值的时候，会改变变量本身的值  （切片）
	*/
	s1 := make([]int, 3, 3)
	s3 := s1
	s3[0] = 100
	fmt.Println(s1, s3)

}
