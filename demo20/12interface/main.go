package main

import "fmt"

// Golang中空接口和类型断言使用细节

type Addresser struct {
	Name     string
	Phonenum int
}

func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"吃饭", "睡觉", "开摆"} //切片赋值给空接口

	//fmt.Print(userinfo["hobby"][0])  空接口类型不支持索引

	var address1 = Addresser{
		Name:     "李四",
		Phonenum: 123456,
	}

	userinfo["address"] = address1 //结构体赋值给空接口
	//fmt.Println(userinfo["address"].Name) 空接口不支持这样获取对应结构体的属性值

	//获取空接口对应的切片，结构体的值 （采用类型断言）
	hobby2, _ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[1])

	adv, _ := userinfo["address"].(Addresser)
	fmt.Println(adv.Name, adv.Phonenum)

}
