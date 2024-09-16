package main

import (
	"fmt"
	"sort"
)

/*
案例2：

	m1 := map[string]string{
		"username":"zhangsan",
		"age":"20",
		"sex":"男",
		"height":"180",
	}
	输出：age=>20height=>180sex=>男username=>zhangsan

	上面有一个map对象m1,封装一个方法,要求按照key升序排序，最后输出一个key=>valuekey=>value的字符串
*/
func sortMapKey(map_sorted map[string]string) string {
	var key_slc []string
	for key, _ := range map_sorted {
		key_slc = append(key_slc, key)
	}
	sort.Strings(key_slc)
	var str string
	for _, v := range key_slc {
		str = str + fmt.Sprintf("%v=>%v", v, map_sorted[v])
	}
	return str

}

func main() {

	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
		"height":   "180",
	}
	fmt.Println(sortMapKey(m1))
}
