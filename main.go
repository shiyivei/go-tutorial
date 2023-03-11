package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	
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

// 查找重复行

func getDuplicateLines() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++

	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}