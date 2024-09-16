package main

import "fmt"

func main() {

	/*
		Go 语言中 break 语句用于以下几个方面：
			• 用于循环语句中跳出循环，并开始执行循环之后的语句。
			• break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
			• 在多重循环中，可以用标号 label 标出想 break 的循环。
	*/

	//1、表示当i=2的时候跳出当前循环
	i := 0
	// for {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Print(i)
	// 	i++

	// }

	//2、 break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。

	//3、在多重循环中，可以用标号 label(名字自取) 标出想 break 的循环。

	/*
		i=0 j=0
		i=0 j=1
		i=0 j=2
	*/
abc:
	for ; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				//break //这样只能跳出内循环
				break abc //这样能一次跳出多层循环
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

}
