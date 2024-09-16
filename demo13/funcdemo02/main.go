package main

import "fmt"

//return 关键词一次可以返回多个值
func sumFunc(x, y int) (int, int, int) {
	sum := 0
	sum = x + y
	return x, y, sum
}

//返回值命名: 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字 返回。
func sumFunc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y //命名的返回值可以直接用，不用再次声明
	return
}

func main() {
	fmt.Println(sumFunc(10, 2))

	a, b, c := sumFunc(10, 2)
	fmt.Println(a, b, c)

	d, e := sumFunc1(10, 2)
	fmt.Println(d, e)
}
