package main

import "fmt" //参考文档 https://www.cnblogs.com/jiangchunsheng/p/14329024.html

func main() {
	//  /*内容*/ 多行 ctrl+/快速住址
	//fmt.Println("hello", "heyhey") // fmt包的其中一种输出语句Println会自动换行，空格

	//fmt.Print("hello", "heyhey") //Print不会自动空格/换行

	var a = "aaa"        //go中变量定义了就必须使用
	fmt.Printf(" %v", a) //格式化输出

	var b int = 10
	var c int = 3
	fmt.Println("a", a, "b=", b, "c=", c)
	fmt.Printf("a=%v,b=%v ,c=%v\n", a, b, c)

	//短变量声明法定义变量
	d := 10
	e := 20
	f := "a"
	fmt.Printf("d=%v,e=%v,f=%v", d, e, f)

	//使用printf打印变量类型
	fmt.Printf("a=%v a的类型是%T", a, a)

}
