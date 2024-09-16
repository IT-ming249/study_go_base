package main

import "fmt"

func main() {
	//如果我们想在map对象中存放一系列的属性的时候，我们就可以把map类型的值定义成切片
	userinfo := map[string][]string{ //map的值是个切片
		"hobby": {"吃饭", "睡觉", "敲代码"},
	}
	userinfo["base_info"] = []string{"张三", "20", "男"}
	fmt.Println(userinfo)

}
