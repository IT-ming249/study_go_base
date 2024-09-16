package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、string类型转换成整型
	str := "123456"
	var num int64
	num, _ = strconv.ParseInt(str, 10, 64) //转换失败的话num为0，error也会有信息
	fmt.Printf("值%v--类型%T\n", num, num)
	/*
		ParseInt
		参数1：string数据
		参数2：进制
		参数3：位数 32 64 16

	*/
	str2 := "12.56"
	var numFloat float64
	numFloat, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("值%v--类型%T", numFloat, numFloat)
	/*
		ParseFloat
		参数1：string数据
		参数2：位数 32 64

	*/

	//string转另外两种数据类型意义不大

}
