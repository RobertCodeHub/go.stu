package main

import (
	_ "bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	
	return s
}

func main() {
	fmt.Println(comma("1234435465"))
}