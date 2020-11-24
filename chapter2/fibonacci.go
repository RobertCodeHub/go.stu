package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func fibn(n int) int {
	if n == 1 || n == 2{
		return 1
	}
	return fibn(n-1)+fibn(n-2)
}

//费波纳茨数列
func main() {

	a := fib(10)
	fmt.Printf("fib value %d\n", a)
}