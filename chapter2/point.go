package main

import "fmt"

//函数返回局部变量的地址
func f() *int {
	v := 1
	return &v
}

func link_one() {
	var p = f()
	fmt.Printf("p value %v \n", p)
}

func link_two() {
	var p = f()
	fmt.Printf("p value %v \n", p)
}

func incr(p *int) int {
	*p++
	return *p
}

func main() {

	var cat int = 1
	var str string = "banana"

	//通过&(取地址符) 获取到 *int 和 * string 类型的指针
	fmt.Printf("%p %p \n", &cat, &str)

	var x, y int
	//当且仅当两个指正同时指向同一个变量或者两个指针的值都为nil时相等
	fmt.Println(&x == &x, &y == &x, &x == nil)

	//两次调用返回的指针地址是不相同的
	link_one()
	link_two()

	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}