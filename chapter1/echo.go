package main

import(
	"fmt"
	"os"
)

func main(){

	//这里如果有os.Args[1:] 表示从数组的第一个元素开始循环.os.Args[0:] 表示从第0个元素开始循环
	//这里默认第0个元素是程序的执行路径
	for index, value := range os.Args[1:] {
		fmt.Printf("index : %d -> value : %s \n", index, value)
	}

	var s, str string
	for i := 1; i < len(os.Args); i++ {
		s += str + os.Args[i]
		str = " "
	}

	fmt.Println(s)
}
