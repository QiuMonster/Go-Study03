package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

//写数据要进行刷新
func go_bufwrite() {
	// go_bufwrite1()
	// go_bufwrite2()
	// go_bufwrite3()
	// go_bufwrite4()
	go_bufwrite5()
}

//写入文件数据
func go_bufwrite1() {
	f, _ := os.OpenFile("lks.txt", os.O_RDWR, os.ModePerm) //文件中实现了 reader 和 writer接口
	w := bufio.NewWriter(f)
	w.WriteString("hello world")
	w.Flush() //进行刷新操作  不然数据在缓冲区中 无法正常写入文件
	f.Close()
}
//将数据写入 缓冲数组中
func go_bufwrite2() {
	b := bytes.NewBuffer(make([]byte, 0)) //实现了 writer接口
	w := bufio.NewWriter(b)
	//将字符串写入缓冲区内
	w.WriteString("hello lks")
	w.Flush() //要进行刷新
	line, _ := b.ReadString('\n')
	fmt.Println(line)
}
//显示缓冲区中剩余大小于已经使用大小
func go_bufwrite3(){
	b := bytes.NewBuffer(make([]byte, 0))//默认是4096
	w := bufio.NewWriter(b)
	w.WriteString("hwllloo pp")
	// b.Reset()//将缓冲区清零
	fmt.Println("缓冲区剩余大小:",w.Available(),"缓冲区已使用大小:",w.Buffered())
	w.Flush()
}
//写入非字母  使用 writerune 只能写入单个字符 函数
func go_bufwrite4(){
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)
	w.WriteRune('你')//相当于 int32   byte 相当于 int8
	w.Flush()
	fmt.Println(b)
}
//构建一个具有写与读的对象
func go_bufwrite5(){
	b := bytes.NewBuffer(make([]byte, 0))
	w := bufio.NewWriter(b)
	r := strings.NewReader("hello")
	r2 := bufio.NewReader(r)
	//获得一个可读可写的对象
	rw := bufio.NewReadWriter(r2, w)
	rw.WriteString("huhu")
	s, _ := rw.ReadString('\n')
	rw.Flush()
	fmt.Println("写入数组中的数据，与读的数据是不同的:",b)
	fmt.Println("读取的内容其实是之前读对象读取的内容:",s)

}
