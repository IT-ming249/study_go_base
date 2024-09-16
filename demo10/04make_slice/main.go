package main

import "fmt"

func main() {
	/*1、切片声明和初始化的几种方法：*/
	var slc = []string{"python", "C", "rust", "golang"}
	fmt.Printf("%v-%T-%d", slc, slc, cap(slc))
	//2、make()函数创建一个切片  make([]T, size, cap)
	var silceA = make([]int, 4, 8)
	fmt.Println(silceA)
	fmt.Printf("%v-%d", len(silceA), cap(silceA))
}
