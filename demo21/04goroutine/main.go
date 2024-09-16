package main

import (
	"fmt"
	"sync"
	"time"
)

// 多协程的一个案例
var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("协程", num, i)
		time.Sleep(time.Millisecond * 50)
	}

}
func main() {
	for j := 1; j <= 10; j++ {
		wg.Add(1)
		go test(j)
	}
	wg.Wait()
}
