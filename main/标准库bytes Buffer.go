package main

import (
	"bytes"
	"fmt"
	"io"
)

func go_byte_buffer() {
	// go_byte_buffer1()
	// go_byte_buffer2()
	go_byte_buffer3()
}

//从缓冲区读取字节到指定字节数组中去
func go_byte_buffer3() {
	var b bytes.Buffer
	b.WriteString("hello world")
	for {
		b1 := make([]byte,4)
		_, err := b.Read(b1) //将数据读到字节数组中
		if err==io.EOF{
			break
		}
		fmt.Println(string(b1))
	}
}

//向缓冲区写入指定数据
func go_byte_buffer2() {
	var b bytes.Buffer
	//b.Bytes()才是数组内容对象
	b1 := make([]byte, 10) //单纯的数组是无法直接操作写入数据的
	fmt.Println(b1)
	n, _ := b.WriteString("hello")
	fmt.Println("字符长度:", n)
	fmt.Println("写入的数据是:", string(b.Bytes()))
}

//创建一个缓冲区
func go_byte_buffer1() {
	//创建的其实是一个字节数组
	var b bytes.Buffer
	fmt.Println(b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Println(b1)
	b2 := bytes.NewBuffer([]byte("hello"))
	fmt.Println(b2)
}
