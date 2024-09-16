package main

import "fmt"

func main() {
	// 1、make创建map类型的数据  map[KeyType]ValueType map类型只有初始化以后才能使用
	var userinfo = make(map[string]string) //类似Python中的字典
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["gender"] = "男"
	fmt.Println(userinfo)
	fmt.Println(userinfo["username"]) //通过key获取map的值

	// 2、map 也支持在声明的时候填充元素
	var userinfo2 = map[string]string{
		"username": "李四",
		"age":      "20",
		"gender":   "女",
	}
	fmt.Println(userinfo2)
	fmt.Println(userinfo2["username"])

	// 3、第三种创建map类型数据的方法
	userinfo3 := map[string]string{
		"username": "王五码子",
		"age":      "10",
		"gender":   "男",
	}
	fmt.Println(userinfo3)
	fmt.Println(userinfo3["username"])

}
