package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	map1[11] = 100
	map1[2] = 13
	fmt.Println(map1)

	//需求：按照key升序输出map的key=>value   (签名算法)

	//1、把map的key放在切片里面
	var map_key []int
	for key, _ := range map1 {
		map_key = append(map_key, key)
	}
	fmt.Println(map_key)
	//2、让key进行升序排序
	sort.Ints(map_key)
	fmt.Println(map_key)
	//3、循环遍历key输出map的值
	for _, v := range map_key {
		fmt.Printf("key:%v,value:%v\n", v, map1[v])
	}

}
