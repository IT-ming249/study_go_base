package main

import "fmt"

func main() {
	var silceA = make([]int, 4, 8)
	//修改切片的内容（通过下标）
	silceA[0] = 10
	silceA[1] = 20
	silceA[2] = 30
	fmt.Println(silceA)

	var slc = []string{"python", "C", "rust", "golang"}
	slc[0] = "java"
	fmt.Println(slc)

	//golang中没法通过下标的方式给切片扩容
	// golang中给切片扩容的话要用到append()方法
	var sliceB []int
	fmt.Printf("%v-%v-%v", sliceB, len(sliceB), cap(sliceB))

}
