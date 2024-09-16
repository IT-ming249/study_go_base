package main

import (
	"fmt"
)

func main() {
	//Go 语言中的常量

	const pi = 3.1415926 //定义常量必须直接给它赋值
	fmt.Println(pi - 1)  //这样是可以的

	//pi = 3 常量的值不可以改变

	//一次声明多个常量
	const a, b, c = 1, 2, "A"
	fmt.Println(a, b, c)

	//const 同时声明多个常量，如果省略了值则表示和上面一行的值相同
	const (
		d = 100
		e
		f
	)
	fmt.Println(d, e, f)

	//iota是一个计数器可与const结合使用
	const g = iota
	fmt.Println(g)

	const (
		g1 = iota
		g2
		g3
	)
	fmt.Println(g1, g2, g3)

	//iota跳过一个值
	const (
		n1 = iota
		_
		n3
		n4
	)
	fmt.Println(n1, n3, n4)

	//iota声明中插队
	const (
		h1 = iota
		h2 = 100
		h3 = iota
		h4
	)
	fmt.Println(h1, h2, h3, h4)

	//多个iota定义在一行
	const (
		k1, k2 = iota + 1, iota + 2 //1 2
		k3, k4
	)
	fmt.Println(k1, k2, k3, k4)

	//定义多单词变量时建议采用驼峰命名规则
	var MaxAge = 100 //大驼峰
	// var maxAge 小驼峰
	fmt.Println(MaxAge)
}
