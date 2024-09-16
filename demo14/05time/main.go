package main

import (
	"fmt"
	"time"
)

// 日期字符串转换成时间戳
func main() {
	var str = "2022-09-04 15:38:34"
	//layout string要与被转换的字符串是相同的模板,下面只是把字符串转换成了日期对象
	time_obj, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	fmt.Println(time_obj)
	unixtime := time_obj.Unix()
	fmt.Println(unixtime)
}
