package main

import (
	"fmt"
	"time"
)

// 需求：要统计1-120000的数字中那些是素数？for循环实现
func main() {
	start := time.Now().Unix()
	for i := 1; i <= 90000; i++ {
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
	end := time.Now().Unix()
	fmt.Println(end - start)

}
