package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"` // 结构体标签，这表示转换为js格式以后，该属性的键为" "中的内容
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
	jsonByte, _ := json.Marshal(s1)
	jsonstr := string(jsonByte)
	fmt.Println(jsonstr)
}
