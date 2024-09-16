package main

import "fmt"

func main() {
	userinfo3 := map[string]string{
		"username": "王五码子",
		"age":      "10",
		"gender":   "男",
	}
	//for range循环遍历map类型的数据

	for k, v := range userinfo3 {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}

}
