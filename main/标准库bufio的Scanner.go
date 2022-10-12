package main

import (
	"bufio"
	"fmt"
	"strings"
)

func go_scanner() {
	// go_scanner1()
	go_scanner2()
}

//扫描分割
func go_scanner1() {
	//先创建一个reader
	s := strings.NewReader("hello qiu lks and ...")
	s2 := bufio.NewScanner(s)
	//将字符进行分割  分割的方法  以空格进行分割标志
	s2.Split(bufio.ScanWords)
	for s2.Scan() {
		//输出分割出来的字符
		fmt.Println(s2.Text())
	}
}
func go_scanner2() {
	s := strings.NewReader("hello qiu lks and ...")
	s2 := bufio.NewScanner(s)
	//将字符进行分割  分割的方法  以空格进行分割标志
	s2.Split(bufio.ScanBytes)
	for s2.Scan() {
		//输出分割出来的字符
		fmt.Println(s2.Text())
	}
}
