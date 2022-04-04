package main // 包，表明代码所在的模块（包）；应用程序入口：必须是 main 包

import (
	"fmt"
	"os"
) // 引入代码依赖

// 功能实现
func main() { // 应用程序入口：必须是 main 方法
	// main 函数不支持传入参数；在程序中直接通过 os.Args 获取命令行参数
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
	fmt.Println("Hello World")
	// Go 中 main 函数不支持返回任何值；通过 os.Exit 来返回状态
}
