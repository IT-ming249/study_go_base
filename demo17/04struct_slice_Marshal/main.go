package main

//复杂嵌套结构体转js字符串
import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []Student //结构体类型的切片
}

func main() {
	c := Class{
		Title:    "1班",
		Students: make([]Student, 0),
	}
	//为c添加students数据
	for i := 0; i <= 10; i++ {
		s := Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}
	//fmt.Println(c)

	//将结构体的嵌套转换为js字符串
	jsonByte, _ := json.Marshal(c)
	js_str := string(jsonByte)
	fmt.Println(js_str)

}
