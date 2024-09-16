package main

import "fmt"

func main() {
	/*

		switch expression {
		case condition:


		}
	*/

	// 练习：判断文件类型,如果后缀名是.html 输入 text/html, 如果后缀名.css 输出 text/css ,如果 后缀名是.js 输出 text/javascript

	// 1、switch case的基本使用
	var extname = ".html"
	switch extname {
	case ".html":
		fmt.Println("text/html")
	case ".java":
		fmt.Println("text/java")
	}
	// 2、switch case的另一种写法(与if,for,同理)
	// switch exname:=".html";exname{

	// }
	// 3、一个分支可以有多个值，多个 case 值中间使用英文逗号分隔

	// 判断一个数是不是偶数
	n := 5
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	//4、分支还可以使用表达式，这时候 switch 语句后面不需要再跟判断变量。例如
	// var age = 30
	// switch {
	// case age < 20:
	// 	fmt.Println("年轻人")
	// case age > 24 && age <= 60:
	// 	fmt.Println("牛马")
	// case age > 60:
	// 	fmt.Println("老登")
	// default:
	// 	fmt.Println("tan90")
	// }
	//5、 switch 的穿透 fallthrought
	var age = 30
	switch {
	case age < 20:
		fmt.Println("年轻人")
	case age > 24 && age <= 60:
		fmt.Println("牛马")
		fallthrough //只会穿透一层，基本用不上
	case age > 60:
		fmt.Println("老登")
	default:
		fmt.Println("tan90")
	}
	//fallthrough`语法可以执行满足条件的 case 的下一个 case，是为了兼容 C 语言中的 case 设计 的

}
