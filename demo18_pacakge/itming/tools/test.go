package tools //同一个包下可以有多个go文件

import "fmt" //自定义的包可以引入别的包

func PrintInfo() {
	fmt.Println("打印信息")
}
