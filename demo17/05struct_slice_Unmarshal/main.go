package main

import (
	"encoding/json"
	"fmt"
)

// 复杂js字符串，转换回结构体
// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title string
	St    []Student //结构体类型的切片
}

func main() {
	str := `{"Title":"1班","St":[{"ID":0,"Gender":"男","Name":"stu_0"},{"ID":1,"Gender":"男","Name":"stu_1"},{"ID":2,"Gender":"男","Name":"stu_2"},{"ID":3,"Gender":"男","Name":"stu_3"},{"ID":4,"Gender":"男","Name":"stu_4"},{"ID":5,"Gender":"男","Name":"stu_5"},{"ID":6,"Gender":"男","Name":"stu_6"},{"ID":7,"Gender":"男","Name":"stu_7"},{"ID":8,"Gender":"男","Name":"stu_8"},{"ID":9,"Gender":"男","Name":"stu_9"},{"ID":10,"Gender":"男","Name":"stu_10"}]}`
	var c = Class{} // 实例化一个class对象
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
}
