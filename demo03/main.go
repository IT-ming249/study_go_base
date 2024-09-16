package main

import "fmt"

var g = "全局变量"

func getinfo() (string, int) { //(string,int)表示函数要返回的数据类型
	return "明明", 10

}

func main() {
	fmt.Println(g)

	/* 1. var 声明变量
	var 变量名称 类型
	*/

	//var username string
	//fmt.Print(username) //变量声明以后不赋值就是空值

	//go 语言变量定义以及初始化
	var username string
	username = "张三"
	fmt.Println(username)

	//第二种
	var username1 string = "李四" //var username string = "李四"（×） 同一作用域内不支持重复声明变量
	fmt.Println(username1)

	//第三种 类型推导 (推荐)
	var username2 = "王五码子"
	fmt.Println(username2)

	username = "陈六" //变量重新赋值是可以的,重复声明是不行的
	fmt.Println(username)

	/* 2.一次声明多个变量
	var 变量名，变量名 类型

	var（
		变量名 类型
		变量名 类型
	）
	*/
	var a1, a2 string
	a1 = "aa"
	a2 = "aaa"
	fmt.Println(a1, a2)

	var b1, b2, b3 string = "b", "bb", "bbb"
	fmt.Println(b1, b2, b3)

	var (
		uname  string
		age    int = 10
		gender     = "男" //go自带的类型推导

	)
	uname = "朱八八"

	fmt.Println(uname, age, gender)

	/* 3.短变量声明法 只能用于声明局部变量， (用的多)
	 */

	uname1 := "陈九"
	fmt.Println(uname1)
	fmt.Printf("uname1类型 %T\n", uname1)

	//短变量声明也能一次声明多个变量，并初始化变量,可以是不同的数据类型
	q, w, e := 1, 2, "c"
	println(q, w, e)

	/* 匿名变量 使用多重赋值时，想忽略某个值时使用，用_表示
	 */
	var un1, _ = getinfo()
	fmt.Println(un1)

	var _, age1 = getinfo()
	fmt.Println(age1)
}
