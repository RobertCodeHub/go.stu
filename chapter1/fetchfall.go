package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s :%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs  %7d  %s ", secs, nbytes, url)
}

//main 函数在goroutine执行,go语句创建额外的goroutine
func main() {

	//打开日志文件，不存在则创建
	file, _ := os.OpenFile("E:\\error.log", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0777)
	defer file.Close()

	logger:=log.New(file,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		//fmt.Printf(<-ch)
		logger.Println(<-ch)
	}

	//设置输出流
	logger.SetOutput(file)

	//日志前缀
	logger.SetPrefix("[Error]")

	//日志输出样式
	logger.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}