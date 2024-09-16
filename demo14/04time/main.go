package main

import (
	"fmt"
	"time"
)

// 时间戳转换成日期字符串
func main() {
	// unixTime: 1587888473
	var unixTime = 1587888473
	str := time.Unix(int64(unixTime), 0)
	fmt.Println(str.Format("2006-01-02 15:04:05"))

}
