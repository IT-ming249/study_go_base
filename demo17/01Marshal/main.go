package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string //若是私有属性就不能被json包访问
}

func main() {
	var s1 = Student{
		ID:     1,
		Gender: "male",
		Name:   "Tom",
		Sno:    "152209001",
	}
	fmt.Printf("%#v\n", s1)

	jsonByte, _ := json.Marshal(s1) //将结构体数据转换为json数据(序列化)
	jsonStr := string(jsonByte)
	fmt.Printf("%v\n", jsonStr)
}
