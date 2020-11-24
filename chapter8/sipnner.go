package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for _,r := range `-\|/` {
		fmt.Printf("\r%c",r)
		time.Sleep(delay)
	}
}

func fib(x int) int{
	if x<2{
		return 1;
	}
	return fib(x-2)+fib(x-1);
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 50
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n",n, fibN)
}