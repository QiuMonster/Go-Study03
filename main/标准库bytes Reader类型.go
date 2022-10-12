package main

import (
	"bytes"
	"fmt"
)

func go_byte_reader() {
	go_byte_reader1()
}
func go_byte_reader1() {
	data := "123456789"
	r := bytes.NewReader([]byte(data))

	bu := make([]byte, 2)
	r.Read(bu)
	fmt.Println(string(bu))
	fmt.Println("字符未读取长度:", r.Len(), "字符的总长度:", r.Size())
	
	r.Seek(0,0)//设置偏移量 设置为头部
	r.Read(bu)
	fmt.Println(string(bu))//输出同样的内容

	r.ReadAt(bu,4)//2 偏移量 在哪个位置开始进行读取
	fmt.Println(string(bu))
	//每次读取一个字节
	b, _ := r.ReadByte()
	fmt.Println(string(b))

}
