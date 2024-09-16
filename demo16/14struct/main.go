package main

/*
关于嵌套结构体的字段名冲突
*/
type user struct {
	username string
	password string
	jointime string
	address
	email
}
type address struct {
	name     string
	phone    string
	city     string
	jointime string
}

type email struct {
	account  string
	jointime string
}

func main() {
	//命名冲突在获取值时，需要指明是哪个结构体
}
