package main

import "fmt"

/*
func 函数名(参数)(返回值){
	函数体
}
*/

//求两个数的和
func add(a int, b int) int {
	return a + b
}

//求两个数的差
func sub(a int, b int) int {
	return a - b
}

//函数参数的简写
func sub1(x, y int) int { // 前边的x类型和后边的y类型相同就可以这样写,只能简写前边的
	return x - y
}

//函数的可变参数，可变参数是指函数的参数数量不固定。Go 语言中的可变参数通过在参数名后加...来标识

func sumFunc(x ...int) int {
	fmt.Printf("%v-%T\n", x, x) //这里边x是一个切片[]int
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//可变参数和固定参数可以同时使用,可变参数要放在后边
func sumFunc2(x int, y ...int) int {
	fmt.Printf("x:%v-%T----y:%v-%T\n", x, x, y, y) //这里边y是一个切片[]int
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(sub(3, 4))
	a := 10
	b := 20
	fmt.Println(sub1(a, b))
	sum3 := sumFunc(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum3)
	sum4 := sumFunc(10, 12, 23, 25)
	fmt.Println(sum4)

	sum5 := sumFunc2(100, 1, 2, 5, 6, 8, 7)
	fmt.Println(sum5)

}
