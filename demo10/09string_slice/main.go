package main

import "fmt"

func main() {
	//切片可用于修改字符串里的字符
	s1 := "hello golang"
	bytes1 := []byte(s1) //将字符串转换为byte类型的切片
	fmt.Println(string(bytes1))
	bytes1[0] = 'p'
	fmt.Println(string(bytes1))

	s2 := "你好 golang"
	bytes2 := []rune(s2) //将字符串转换为rune类型的切片(带汉字的)
	fmt.Println(string(bytes2))
	bytes2[0] = '我'
	fmt.Println(string(bytes2))

}
