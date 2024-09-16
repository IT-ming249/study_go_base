package main

import "fmt"

func main() {
	//我们想在切片里面放一系列用户的信息,这时候我们就可以定义一个元素为map类型的切片
	var userinfo = make([]map[string]string, 3, 3)
	fmt.Println(userinfo[0]) //map不初始化的默认值是nil

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["gender"] = "男"
	}

	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "19"
		userinfo[1]["gender"] = "女"
	}

	fmt.Println(userinfo)

}
