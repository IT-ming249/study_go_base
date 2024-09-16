package main

import "fmt"

func main() {
	//1、二维数组的定义
	var arr = [3][2]string{ //可以多个[]定义n维数组
		{"北京", "上海"},
		{"广州", "深圳"},
		{"南京", "南昌"},
	}
	fmt.Println(arr[0], arr[0][1])
	//循环遍历多维数组，每一层循环拆解其中一维
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
	//定义二维数组的另一种方式
	var arr1 = [...][2]string{ //只支持[...][3]，不支持[3][...]
		{"北京", "上海"},
		{"广州", "深圳"},
		{"南京", "南昌"},
	}
	fmt.Println(arr1)

}
