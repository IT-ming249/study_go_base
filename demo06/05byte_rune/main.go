package main

import "fmt"

func main() {

	//1、golang中定义字符  字符属于int类型
	a := 'a'
	fmt.Printf("%v--%T\n", a, a)

	//2、原样输出字符
	b := 'a'
	fmt.Printf("%c--%T\n", b, b)

	//3、定义一个字符串输出字符串里面的字符
	c := "this"
	fmt.Printf("%v-%c-%T\n", c[0], c[0], c[0])

	//4、一个汉字占用 3个字节(utf-8), 一个字母占用一个字节
	d := "汉字nb"
	fmt.Println(len(d)) //sizeof方法无法获取字符串占用内存大小，可以用len

	// 5、定义一个字符 字符的值是汉字对应的unicode的十进制编码
	var e = '汉'
	fmt.Printf("%v-%c-%T\n", e, e, e)

	//6、通过循环输出字符串里面的字符
	f := "你好 golang"
	for i := 0; i < len(f); i++ { //for循环遍历字符串通过byte类型
		fmt.Printf("%v(%c) ", f[i], f[i]) //英文能对应，汉字不行，但是会看到占三字节
	}
	fmt.Println()
	for _, v := range f { //rune类型
		fmt.Printf("%v(%c)\n", v, v)
	}

	//7、修改字符串(通过索引)
	s1 := "big"
	s2 := "白萝卜"

	byte_s1 := []byte(s1)       //修改字符串之前把数据类型转换为byte（适用于全英文字符串）
	fmt.Printf("%T\n", byte_s1) //[]uint8 也是个类似数组的切片
	fmt.Printf("%c--%T\n", byte_s1[0], byte_s1[0])
	byte_s1[0] = 'p' //替换的是单个字符 byte_s1[0] = "p"(×)
	fmt.Println(byte_s1, string(byte_s1))

	rune_s2 := []rune(s2)       //这样转换后是个数组，byte同理 （适用于带汉字的字符串）
	fmt.Printf("%T\n", rune_s2) //[]int32
	fmt.Printf("%c\n", rune_s2[0])
	rune_s2[0] = '胡'
	fmt.Println(rune_s2, string(rune_s2))

}
