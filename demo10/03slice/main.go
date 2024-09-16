package main

import "fmt"

func main() {

	//1、基于数组定义切片
	a := [5]int{55, 66, 77, 88, 99}
	fmt.Printf("值%v,类型%T,这是数组\n", a, a)
	b := a[:] //获取数组里面的所有值
	fmt.Printf("值%v,类型%T，这是切片\n", b, b)
	c := a[1:4] //左闭右开
	fmt.Printf("值%v,类型%T，这是切片\n", c, c)
	d := a[2:] //2到最后
	fmt.Printf("值%v,类型%T，这是切片\n", d, d)
	e := a[:3]
	fmt.Printf("值%v,类型%T，这是切片\n", e, e)
	fmt.Println("----------")
	//2、基于切片定义切片
	b1 := b[1:]
	fmt.Printf("值%v,类型%T，这是切片\n", b1, b1)
	fmt.Println("----------")

	//3、关于切片的长度和容量
	// 长度：切片的长度就是它所包含的元素个数
	// 容量：切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	fmt.Printf("长度%v,容量%d\n", len(b), cap(b))

	fmt.Printf("长度%v,容量%d\n", len(b1), cap(b1))

	fmt.Printf("长度%v,容量%d\n", len(c), cap(c)) // 底层数组[55,66,77,88,99] 从下标是1的元素开始往后数到末尾刚刚好4个

	fmt.Printf("长度%v,容量%d\n", len(e), cap(e))
}
