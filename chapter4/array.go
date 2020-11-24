package main

import "fmt"

type Currency int

const(
	USD Currency = iota
	EUP
	GBP
	RMB
	BDT
)

func main() {

	//定义一个数组除了最后一个是-1其他都为0
	r := [...]int{99:-1}
	for _, v := range r {
		fmt.Printf("% \n",v)
	}

	moths := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "March", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"};
	Q2 := moths[4:7]
	fmt.Println(Q2)
}