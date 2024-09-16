package main

import (
	"fmt"
	"sync"
	"time"
)

// 主线程退出后所有的协程无论有没有执行完毕都会退出，所以我们在主进程中可以通过WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func test1() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test2_golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1

}

func test2() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test2_golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1

}

func main() {
	wg.Add(1)  //协程计数器+1
	go test1() //表示开启一个协程，test方法将与main方法并行执行
	wg.Add(1)
	go test2()
	for i := 0; i <= 10; i++ {
		fmt.Println("hello-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait() //主线程需等待协执行完毕
	fmt.Println("主线程退出....")

}
