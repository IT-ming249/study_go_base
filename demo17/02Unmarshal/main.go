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
	str := `{"ID":1,"Gender":"male","Name":"Tom","Sno":"152209001"}` //json字符串
	var s1 Student
	err := json.Unmarshal([]byte(str), &s1) //（反序列化）因为要将json字符串内的信息放入s1中，相当于修改，所以要取s1的地址

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s1)
}
