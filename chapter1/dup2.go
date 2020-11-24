package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {

	//map是使用make创建的引用,当一个map被传递给函数时，函数接受到的是引用的副本
	//在被调用函数中改变,调用函数中的map也可见的
	counts := make(map[string]int)
	file := os.Args[1:]

	fmt.Printf("file len=%d\n", len(file))

	if len(file) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, value := range file {

			//os.Open 返回值: 一个是打开的文件(*os.File) 一个是error类型的值,error的值如果为nil标志文件打开成功.
			f, err := os.Open(value)

			if err != nil {
				fmt.Printf("dup: %v\n", err)
				continue
			}

			countLines(f, counts)

			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}