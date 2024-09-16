package main

import "fmt"

/*
结构体嵌套
*/
type user struct {
	username string
	password string
	address  address //表示user结构体嵌套address结构体
}
type address struct {
	name  string
	phone string
	city  string
}

func main() {
	var user1 user
	user1.username = "Tom"
	user1.password = "126453"
	user1.address.name = "Tom"
	user1.address.phone = "556300"
	user1.address.city = "深圳"

	fmt.Printf("%#v", user1)
}
