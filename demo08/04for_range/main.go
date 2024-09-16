package main

import "fmt"

func main() {
	var str = "你好 Golang"
	for k, v := range str {
		fmt.Printf("key=%v,value=%c\n", k, v)
	}

	var arr = []string{"php", "java", "nodejs", "golang"} //定义一个切片
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	for _, val := range arr {
		fmt.Println(val)
	}

}
