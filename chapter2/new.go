package main

import "fmt"

func delta(old, new int) int {
	return new - old
}

func main() {

	//两个变量类型是零值，且不携带任何信息。当前的视线中两个变量的地址是一样的
	var a int
	var b int
	fmt.Printf("%t \n", (a==b))
}