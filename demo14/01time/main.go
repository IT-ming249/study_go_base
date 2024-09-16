package main

import (
	"fmt"
	"time"
)

// 1、打印当前日期
func main() {

	//注意：%02d 中的 2 表示宽度，如果整数不够 2 列就补上 0

	timeObj := time.Now()
	fmt.Println(timeObj)

	year := timeObj.Year()
	montn := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%d %02d:%02d:%02d", year, montn, day, hour, minute, second)
}
