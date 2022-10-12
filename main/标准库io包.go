package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//io库中主要的接口
type Reader interface {
	Read(p []byte) (int, error)
}
type Writer interface {
	Write(p []byte) (int, error)
}

func go_io() {
	// go_io1()
	go_testcopy()

}
func go_io1() {
	//使用字符串工具类进行操作
	r := strings.NewReader("hello world")
	for {
		buf := make([]byte, 10)
		_, err := r.Read(buf) //将数据读到字节数组中
		//循环读取数据结束标志
		if err == io.EOF {
			break
		}
		fmt.Println("输出读取的数据", string(buf))
	}
}

//测试试用io流进行数据复制操作
func go_testcopy() {
	r := strings.NewReader("hello world")
	// os.Open("  ")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
