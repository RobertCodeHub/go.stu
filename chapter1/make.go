package main

import (
	"fmt"
)

//golang 分配内存主要有内置函数new和make
//map只能为slice, map, channel分配内存，并返回一个初始化的值
func main() {
	a :=  make([]int, 0)
	n := 20
	for i := 0; i < n; i++ {
		a = append(a, 1)
		fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
	}
}
