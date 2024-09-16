package main

import "fmt"

func main() {
	//1、创建channel: make (chan 元素类型，容量)  管道（先进先出）
	ch := make(chan int, 3)
	//2、给管道里面存储数据
	ch <- 10
	//3、获取管道里面的内容
	a := <-ch
	fmt.Println(a)
	//4、打印管道的长度和容量

	// 5、管道的类型（引用数据类型）

	//8、管道阻塞

	// ch6 := make(chan int, 1)
	// ch6 <- 34
	// ch6 <- 64 //all goroutines are asleep - deadlock!

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	// ch7 := make(chan string, 2)
	// ch7 <- "数据1"
	// ch7 <- "数据2"
	// m1 := <-ch7
	// m2 := <-ch7
	// m3 := <-ch7
	// fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!

	//正确的写法

}
