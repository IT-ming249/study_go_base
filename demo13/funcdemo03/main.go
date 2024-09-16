package main

import "fmt"

//int类型升序排序

//int类型降序排序  切片是引用数据类型

func sortIntAsc(slc []int) []int {
	for i := 0; i < len(slc); i++ {
		for j := i + 1; j < len(slc); j++ {
			if slc[i] > slc[j] {
				t := 0
				t = slc[i]
				slc[i] = slc[j]
				slc[j] = t

			}
		}
	}
	return slc
}

//int类型降序排序  切片是引用数据类型
func sortIntDesc(slc []int) []int {
	for i := 0; i < len(slc); i++ {
		for j := i + 1; j < len(slc); j++ {
			if slc[i] < slc[j] {
				t := 0
				t = slc[i]
				slc[i] = slc[j]
				slc[j] = t

			}
		}
	}
	return slc
}

func main() {

	//案例1：把前面讲的选择排序封装成方法，实现整型切片的升序降序排序排列

	var sliceA = []int{12, 34, 37, 55, 556, 36, 2}
	fmt.Println(sortIntAsc(sliceA))
	fmt.Println(sortIntDesc(sliceA))

}
