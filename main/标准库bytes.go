package main

import (
	"bytes"
	"fmt"
)

func go_bytes() {
	// go_bytes1()
	go_bytes2()
}

//强制类型转换
func go_bytes1() {
	var i int = 21
	var b byte = 30
	b = byte(i)
	i = int(b)
	fmt.Println(b, i)

	s := "hello world"
	b1 := []byte{66, 67, 68, 69}
	s = string(b1)
	fmt.Println(s, b1)
}
func go_bytes2() {
	s := "qiu and lks"
	b := []byte(s)
	b1 := []byte("qiu and")
	// b2 := []byte("QIU AND LKS")
	//查看字符是否存在包含关系
	b3 := bytes.Contains(b, b1)
	fmt.Println(b3)
	//字符包含的个数
	i := bytes.Count(b, b1)
	fmt.Printf("i: %v\n", i)
	//比较两个字符
	b2 := bytes.Equal(b, b1)
	fmt.Printf("b2: %v\n", b2)
	//生成重复的字符
	b4 := bytes.Repeat(b, 2) //1目标字符 2 重复次数
	fmt.Println(string(b4))
	//字符替换
	b5 := bytes.Replace(b, []byte("qiu"), []byte("wlb"), 0) //1 目标字符 2 被换的字符 3 换成的字符 4 可交换的最大次数
	fmt.Println(string(b5))
	//Runs 转换为 Runs的切片 便于知道具体存在多少个中文字符
	ss := []byte("你好世界") //一个中文字符 占 3个字节
	r := bytes.Runes(ss) //转换字符长度  3 -> 1
	fmt.Println(len(r), len(ss))
	fmt.Println(string(r))
	//Join 字节切片的连接 用什么字符进行连接
	barr := [][]byte{[]byte("hello"), []byte("world")}
	b6 := bytes.Join(barr, []byte(",."))
	fmt.Println("切片连接后的数据:",string(b6))
}
