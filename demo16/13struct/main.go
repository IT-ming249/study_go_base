package main

import "fmt"

/*
嵌套匿名结构体
*/

type user struct {
	username string
	password string
	jointime string
	address
}
type address struct {
	name     string
	phone    string
	city     string
	jointime string
}

func main() {
	var user1 user
	user1.username = "Tom"
	user1.password = "126453"
	user1.address.name = "Tom"
	user1.address.phone = "556300"
	user1.address.city = "深圳"

	user1.city = "上海" // 先在user结构体找city，找不到就在address结构体里面找
	user1.jointime = "2020-1-1"
	user1.address.jointime = "2022-1-1" //允许被嵌套的结构体中有与父结构体同名的字段，不过就是要分清楚

	fmt.Printf("%#v", user1)
}
