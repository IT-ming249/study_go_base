package main

import "fmt"

func main() {
	//如果我们想在切片里面放一系列用户的信息的时候我们就可以定义一个map类型的切片
	var userinfo = make([]map[string]string, 3, 3)

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
	for _, v := range userinfo {
		//fmt.Println(v)
		for k, v1 := range v {
			fmt.Println(k, v1)
		}
	}
}
