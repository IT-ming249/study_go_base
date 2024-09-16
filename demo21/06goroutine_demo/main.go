package main

import (
	"fmt"
	"sync"
	"time"
)

//需求：要统计1-120000的数字中那些是素数？goroutine  for循环实现

/*
1 协程  统计  1-30000

2 协程  统计  30001-60000

3 协程  统计  60001-90000

4 协程  统计  90001-120000

// start:(n-1)*30000+1       end:n*30000
*/
var wg sync.WaitGroup

func test(n int) {

	for i := (n-1)*30000 + 1; i <= n*30000; i++ {
		//var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				//flag = false
				break
			}
		}
		// if flag == true {
		// 	fmt.Println("素数", i)
		// }

	}
	wg.Done()
}
func main() {
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println("执行完毕")
}
