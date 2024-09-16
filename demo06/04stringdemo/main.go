package main

import (
	"fmt"
	"strings"
)

func main() {
	//1、定义string类型
	var str string = "hello"
	var str2 = "hello2"
	str3 := "hello3"
	fmt.Printf("%T,%T,%T\n", str, str2, str3)

	//2、字符串转义符 \n表示换行
	st1 := "C:\\GO" //输出反斜杠,双引号,多加一个斜杠
	st2 := "C:\"GO"
	fmt.Println(st1, st2)

	// 3、多行字符串   `(反引号)  tab键上面
	st3 := `this is a string. 
And this is the second row.`
	fmt.Println(st3)

	// 4、len(str) 求长度
	fmt.Println(len(st3))

	// 5、+ 或者 fmt.Sprintf拼接字符串
	fmt.Println("hello " + st2)
	st4 := fmt.Sprintf("%v %v %d", st1, st2, 3) //Sprintf() 获取()内输出的值，返回给一个变量
	fmt.Printf("%v %T\n", st4, st4)

	//6、strings.Split 分割字符串  strings需要引入strings包
	st5 := "分割字符串,strings需要引入strings包"
	arr := strings.Split(st5, ",")      //(拆分谁，以什么拆分)
	fmt.Printf("arr切片的数据类型: %T\n", arr) //是一种类似数组的数据类型，但不完全是数组
	fmt.Println(arr)

	//7、strings.Join(a[]string, sep string) join 操作  表示把切片链接成字符串
	stj := strings.Join(arr, "*") //(连接哪个切片，用什么连接)
	fmt.Println(stj)
	arr2 := []string{"php", "java", "python"} //定义一个切片
	stj2 := strings.Join(arr2, " ")
	fmt.Println(stj2)

	//8、 strings.contains 判断是否包含
	st6 := "this is a string"
	str7 := "is"
	result := strings.Contains(st6, str7)
	fmt.Println(result)

	//9、strings.HasPrefix,strings.HasSuffix 前缀/后缀判断
	result2 := strings.HasPrefix(st6, str7)
	result3 := strings.HasSuffix(st6, str7)
	fmt.Println(result2, result3)

	//10、 strings.Index(),strings.LastIndex() 子串出现的位置  查找不到返回-1 查找到返回下标位置 下标是从0开始的
	res := strings.Index(st6, str7)      //从前往后找
	res2 := strings.LastIndex(st6, str7) //从后往前找
	fmt.Println(res, res2)
}
