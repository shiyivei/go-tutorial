package main

import (
	"fmt"
	"os"
	"strings"
	"mastering-go/get_start"
	

)

func main() {

	// fmt.Println("hello world")
	// get_start.StartApp()
	
	// get_start.GetURL()

	get_start.StartFetch()
}

// 获取命令行参数
func geCommandArgs() {

	// 变量声明：隐式才初始化
	var s, sep string

	// fmt.Println("s:",s, "sep:",sep)

	// for循环三段论
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}

	// fmt.Println(os.Args[2])

	// for 循环 range形式
	for _, arg := range os.Args {

		s += sep + arg
		sep = ""
	}
	// fmt.Println("s",s)

	// strings 包的使用
	fmt.Println(strings.Join(os.Args[1:], " "))
}

