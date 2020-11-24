package main

import (
	"fmt"
	"strings"
)

//获取到最后一个/之后.之前的字符串
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

//简化代码
func basenametwo(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s;
}

func main() {
	s := "/hom/java/apache-tomcat/bin/catalina.sh"
	fmt.Printf("Process Result: %s\n", basename(s))
}
