package main

import (
	"fmt"
)

//student结构体

//判断参数是不是结构体类型

//1、通过类型变量里面的Field可以获取结构体的字段

//2、通过类型变量里面的FieldByName可以获取结构体的字段

//3、通过类型变量里面的NumField获取到该结构体有几个字段

// 打印执行方法
func PrintStructFn(s interface{}) {

	//1、通过类型变量里面的Method可以获取结构体的方法

	fmt.Println("--------------------------")
	//2、通过类型变量获取这个结构体有多少个方法

	fmt.Println("--------------------------")
	//3、通过《值变量》执行方法 （注意需要使用值变量，并且要注意参数） v.Method(0).Call(nil) 或者v.MethodByName("Print").Call(nil)

	//4、执行方法传入参数 （注意需要使用《值变量》，并且要注意参数,接收的参数是[]reflect.Value的切片）

	// 5、获取方法数量

}

func main() {

}
