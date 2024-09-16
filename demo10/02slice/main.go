package main

import "fmt"

func main() {
	//切片的循环遍历
	var strSlice = []string{"php", "java", "node", "golang"}

	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	for k, v := range strSlice {
		fmt.Println(k, v)
	}
}
