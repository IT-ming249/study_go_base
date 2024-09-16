package main

import (
	"fmt"
	"time"
)

/*
进程：一个正在执行的程序
线程：进程的一个执行实例

并发：多个线程同时竞争一个位置，竞争到才可以执行，每一段时间只有一个线程在执行
并行：多个线程可以同时执行，每一个时间段，可以有多个线程同时执行

golang采用多协程实现并行并发，比其它编程语言占的内存更小，调度更快
*/

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔50毫秒秒输出 "你好golang"
// 在主线程中也每隔50毫输出"你好golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test_golang")
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	go test() //表示开启一个协程，test方法将与main方法并行执行
	for i := 0; i <= 10; i++ {
		fmt.Println("hello")
		time.Sleep(time.Millisecond * 100)
	}
}
