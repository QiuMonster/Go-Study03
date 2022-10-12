package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	// "strings"
)

//缓冲输出
func go_buf() {
	// go_buf1()
	// go_iobuf3()
	// go_iobuf2()
	// go_iobuf4()
	// go_iobuf5()
	// go_bufio6()
	go_bufio7()
}

//使用一个缓冲区进行输出  默认大小是 4096
func go_buf1() {
	f, _ := os.Open("lks.txt")
	// r1 := strings.NewReader("hello world")
	r := bufio.NewReader(f)
	b, _ := r.ReadString('\n')
	fmt.Println(b)
}
func go_iobuf2() {
	r := strings.NewReader("hello world")
	//将数据存入缓冲区
	r2 := bufio.NewReader(r)

	for {
		p := make([]byte, 10)
		//将缓冲区中的数据读到 字节数组中去
		_, err := r2.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p))
	}
}
func go_iobuf3() {
	s := strings.NewReader("ABCDEF")
	r := bufio.NewReader(s)
	//每次只读取一个字节
	b, _ := r.ReadByte()
	fmt.Println("读取到的字节是:", string(b))

	r.UnreadByte() //吐出一个字节
	b2, _ := r.ReadByte()
	fmt.Println("读取到的字节是:", string(b2))
}

//读取一行数据
func go_iobuf4() {
	s := strings.NewReader("ABC\nDEF\nGHI\nJKN")
	r := bufio.NewReader(s)
	//读取一行的数据
	line, isPrefix, _ := r.ReadLine()
	fmt.Println(string(line), isPrefix) //isPrefix当一行的数据过多时 会省略后面的数据
}

//读取一个切片  需要循环读取才能读出全部数据
func go_iobuf5() {
	s := strings.NewReader("abc,def,ghj")
	r := bufio.NewReader(s)
	//以什么进行分割 读取数据
	line, _ := r.ReadSlice(',')
	fmt.Println("使用,分割出来的数据:", string(line))
}

//以什么方式读取字符串 或字节
func go_bufio6() {
	s := strings.NewReader("asdf,ffgfh,kjkj")
	r := bufio.NewReader(s)
	// r.ReadBytes(',')           //以逗号进行字节读取
	s2, _ := r.ReadString(',') //以,进行字符串读取
	fmt.Println("读取到的字符串:", s2)

}
//数据写入
func go_bufio7(){
	s := strings.NewReader("asdf,ffgfh,kjkj")
	r := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))
	r.WriteTo(b)//也可以向文件中写入数据
	fmt.Println("写入的数据是:",b)
}
