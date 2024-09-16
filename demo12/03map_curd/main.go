package main

import "fmt"

func main() {

	//map类型数据的curd
	//1、创建 map类型的数据
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["gender"] = "男"
	fmt.Println(userinfo)

	//2、创建 修改map类型的数据
	userinfo["username"] = "李四"
	fmt.Println(userinfo)

	//3、获取 查找map类型的数据
	fmt.Println(userinfo["username"]) //获取
	v, ok := userinfo["age"]          //查找 ok为布尔类型，存在的话返回ture
	v1, okk := userinfo["manager"]
	fmt.Println(v, ok)
	fmt.Println(v1, okk)

	//4、删除map数据里面的key以及对于的值 delete(map对象,对应的key)
	delete(userinfo, "gender")
	fmt.Println(userinfo)

}
