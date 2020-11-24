package main

import (
	"fmt"
	"strings"
)

func main() {

	//string
	s := "Hello World!"
	fmt.Println("string length:",len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[6:])
	fmt.Println(s[:6])

	j := "プログラム";
	//%x 已16进制格式输出,其中x前面有一个空格表示没两个字符分割一个空格
	fmt.Printf("% x\n", j)

	//将UTF-8编码转为字符串
	fmt.Printf("UTF-8 Change To String: %s\n", string(rune(11907))+string(0x2E81))

	//将文字符号类型的切片转化为字符串
	r := []rune(j)
	fmt.Printf("%s\n", string(r))

	//是否包含子串
	fmt.Printf("Has Cotanint SubString:%t\n",strings.ContainsAny("2020 United States presidential election", "ele"))

	//子串出次数: Rabin-Karp 算法实现
	fmt.Println(strings.Count("Hello World", "l"))

	//Fields 用一个或多个连续的空格分隔字符串 s，返回子字符串的数组（slice）
	fmt.Printf("Fields are: %q\n", strings.Fields(" 2020 United States Presidential election "))

	//Split 分割
	fmt.Printf("Spilt are: %q\n", strings.Split("Those,numbers,represent,a,tectonic,shift,away,from,one-day,voting,the,staple,of,the,American,electoral,system,for,centuries",","))
	fmt.Printf("Spilt After are: %q\n", strings.SplitAfter("Those,numbers,represent,a,tectonic,shift,away,from,one-day,voting,the,staple,of,the,American,electoral,system,for,centuries",","))

}