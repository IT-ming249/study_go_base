/*
导入第三方包的流程

1. go mod init 项目名称  项目名称建议跟目录名称统一
2. 配置第三方包
3. go mod tidy 下载当前项目缺少依赖项
4. 运行项目
*/
package main

import (
	"fmt"

	"github.com/shopspring/decimal" //引入第三方包
)

func main() {
	fmt.Println()

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
