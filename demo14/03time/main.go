package main

import (
	"fmt"
	"time"
)

/*
时间戳是自 1970 年 1 月 1 日（08:00:00GMT）至当前时间的总毫秒数。它也被称为 Unix 时 间戳（UnixTimestamp）。
*/
func main() {
	timeObj := time.Now()
	unixTime := timeObj.Unix() //获取当前时间的时间戳
	fmt.Println(unixTime)

	unixnaTime := timeObj.UnixNano() //纳秒时间戳
	fmt.Println(unixnaTime)
}
