package main

import "fmt"

func main() {

	// Go 语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素

	// 要删除索引为 2 的元素    注意：append合并切片的时候最后一个元素要加...
	sliceA := []int{1, 2, 3, 4, 5, 32}
	sliceA = append(sliceA[:2], sliceA[3:]...) //通过append删除
	fmt.Println(sliceA)

	slice := []int{10, 20, 30, 40, 50, 32}
	slice = append(slice[:3], slice[5:]...)
	fmt.Println(slice)

}
