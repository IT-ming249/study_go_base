package main

import (
	"fmt"
	"strings"
)

func main() {
	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	var str = "how do you do"
	var str_splic = strings.Split(str, " ")

	//我自己想的
	// count_how := 0
	// count_do := 0
	// count_you := 0
	// for _, v := range str_splic {
	// 	if v == "how" {
	// 		count_how++
	// 	} else if v == "do" {
	// 		count_do++
	// 	} else if v == "you" {
	// 		count_you++
	// 	}
	// }
	// fmt.Printf("how:%v,do:%v,you:%v", count_how, count_do, count_you)

	//map方法
	strMap := make(map[string]int)
	for _, v := range str_splic {
		strMap[v]++
	}
	fmt.Println(strMap)
}
