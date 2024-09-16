package main

import (
	"fmt"
)

func main() {
	// var userinfo map[string]string
	// userinfo["username"] = "张三" //× 引用数据类型只有分配了内存空间才能初始化
	// fmt.Println(userinfo)
	//正确写法
	userinfo := make(map[string]string)
	userinfo["username"] = "张三"
	fmt.Println(userinfo)

	// var a *int
	// *a = 100 //×  指针变量也是引用数据类型
}
