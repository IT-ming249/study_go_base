package main

import (
	"fmt"
	"time"
)

/*
	golang定时器
*/

func main() {

	ticker := time.NewTicker(time.Second) // 定义一个间隔一秒的计时器
	n := 5

	for t := range ticker.C { //ticker.C只返回一个参数
		n--
		fmt.Println(t)
		if n == 0 {
			ticker.Stop() //终止计时器，释放内存，break只跳出循环，不释放计时器内存
			break
		}
	}
	//休眠方法
	fmt.Println("aaa")
	time.Sleep(time.Second)
	fmt.Println("aaa")
	time.Sleep(time.Second)
	fmt.Println("aaa")
	time.Sleep(time.Second * 2)
	fmt.Println("aaa")

}
