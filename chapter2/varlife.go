package main

import (
	"fmt"
)

var global *int

func fun() {
	var x int
	x = 1
	global = &x
}

func main() {
	fun()
	fmt.Printf("%d", *global)
}
