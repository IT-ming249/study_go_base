package main

import "fmt"

func main() {
	//var 数组变量名 [元素数量]数据类型
	//1、数组的长度是类型的一部分,数组长度定义好以后就不可改变了
	var arr1 [3]int
	var strarr [2]string
	fmt.Printf("%T %T\n", arr1, strarr)

	//2、数组的初始化 第一种方法
	fmt.Println(arr1) //查看默认值
	fmt.Println(strarr)

	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	strarr[0] = "php"
	strarr[1] = "java"
	fmt.Println(strarr)

	//、数组的初始化 第二种方法 var 数组变量名 = [元素数量]数据类型{初始化元素}
	var arr2 = [3]int{3, 4, 5}
	fmt.Println(arr2)
	strarr2 := [2]string{"golang", "C++"} //类型推导
	fmt.Println(strarr2)

	//3、数组的初始化 第三种方法  一般情况下我们可以让编译器 根据初始值的个数自行推断数组的长度
	var numArry = [...]int{7, 8, 9}
	strArry := [...]string{"golang", "rust"}
	fmt.Println(numArry, len(numArry))
	fmt.Println(strArry, len(strArry))
	//改变数组里面的值,直接通过索引
	numArry[0] = 10
	fmt.Println(numArry)

	// 4、数组的初始化 第四种方法   我们还可以使用指定索引值的方式来初始化数组（了解）
	arr3 := [...]int{0: 1, 2: 20, 5: 50}
	fmt.Println(arr3, len(arr3))

	//5、数组的循环遍历 for  for range
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for k, v := range strArry {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}

}
