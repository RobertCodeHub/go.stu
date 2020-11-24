package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}

//匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
//例子中 a,_ 获取的是第一返回值，第二个返回值没有返回因为是匿名变量.
func main(){
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}