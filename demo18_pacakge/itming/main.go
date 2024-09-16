package main

import (
	"fmt"
	"itming/calc"    //通过目录导入自定义的包
	T "itming/tools" //引入包的别名
)

func init() {
	fmt.Println("初始化方法执行....")
} //init方法优先于 main执行  如果引入的包中含有init方法，则最后引入的包的init方法最先执行

func main() {
	//调用calc包里面的方法
	sum := calc.Add(10, 2)
	fmt.Println(sum)
	fmt.Println(calc.Test_var) // 可以访问calc包中的公有变量

	//调用tools包里的方法
	fmt.Println(T.Mul(2, 3))
	T.PrintInfo()
}
