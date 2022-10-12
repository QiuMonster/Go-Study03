package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//使用ioutil工具包进行数据输出
func go_ioutil() {
	go_ioutil1()
	go_ioutil2()
	go_ioutil3()
	go_ioutil4()
	go_ioutil5()

}

//读取文件数据
func go_ioutil2() {
	file, _ := os.Open("qiu.txt")
	//file已经实现了 reader接口 可以直接传入
	b, _ := ioutil.ReadAll(file)
	fmt.Println(string(b))
	file.Close()
}

//读取字符串
func go_ioutil1() {
	b, _ := ioutil.ReadAll(strings.NewReader("hello world"))
	fmt.Println(string(b))
}

//读取文件夹
func go_ioutil3() {
	fi, _ := ioutil.ReadDir("../main")
	for _, val := range fi {
		fmt.Println("文件名称:", val.Name())
	}
}

//读取文件
func go_ioutil4() {
	b, _ := ioutil.ReadFile("qiu.txt")
	fmt.Println(string(b))
}

//写文件
func go_ioutil5() {
	ioutil.WriteFile("lks.txt", []byte("lks love wlb"), os.ModePerm)
}
