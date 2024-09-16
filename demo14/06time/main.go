package main

import (
	"fmt"
	"time"
)

// 日期字符串转换成时间戳
func main() {

	/*
		1、time包中定义的时间间隔类型的常量如下：
			const (
			    Nanosecond  Duration = 1
			    Microsecond          = 1000 * Nanosecond
			    Millisecond          = 1000 * Microsecond
			    Second               = 1000 * Millisecond
			    Minute               = 60 * Second
			    Hour                 = 60 * Minute
			)
	*/
	fmt.Println(time.Microsecond)
	fmt.Println(time.Second)
	/*
		2、时间操作函数
	*/

	var time_obj = time.Now()
	fmt.Println(time_obj)
	time_obj = time_obj.Add(time.Hour) //当前时间增加一小时 ，Add可以换成算术运算符

	fmt.Println(time_obj)

}
