package get_start


import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io/ioutil"
	
)

// 2.1 查找重复行

func GetDuplicateLines() {
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

// 2.2 查找重复行

func ReadContentFromFile() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		CountLines(os.Stdin,counts)
	}else {
		for _, arg := range files {
			f,err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr,"dup2: %v\n", err)
				continue
			}

			CountLines(f,counts)

			f.Close()
		}
	}

	for line,n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

func CountLines(f *os.File,counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

// 2.3 查找重复行

func readFile() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data,err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)

			continue
		}

		for _, line := range strings.Split(string(data),"\n") {
			counts[line] ++ 
		}
	}

	for line,n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}

}