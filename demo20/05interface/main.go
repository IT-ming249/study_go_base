package main

import "fmt"

//定义一个方法，可以传入任意数据类型，然后根据不同的类型实现不同的功能
func type_func(itf interface{}) {
	v_int, ok_int := itf.(int)
	if ok_int {
		fmt.Println("这是个整型", v_int)
	}
	v_string, ok_string := itf.(string)
	if ok_string {
		fmt.Println("这是个字符串", v_string)
	}
	v_bool, ok_bool := itf.(bool)
	if ok_bool {
		fmt.Println("这是个布尔", v_bool)
	}

}

//x.(type) 判断一个变量的类型  这个语句只能用在switch语句里面
func type_func_switch(itf interface{}) {
	switch itf.(type) {
	case int:
		fmt.Println("整型")
	case string:
		fmt.Println("字符串")
	case bool:
		fmt.Println("布尔")
	default:
		fmt.Println("断言类型不在范围内")
	}

}

//类型断言
func main() {
	var a interface{}

	a = "hello"
	v, ok := a.(string) //类型断言  值，判断是否符合猜想类型返回bool := 空接口名.（猜想类型）
	if ok {
		fmt.Println("a是string类型,值", v)
	} else {
		fmt.Println("false")
	}

	type_func(15)
	type_func_switch("hihi")

}
