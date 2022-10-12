package main

import "os"

//文件的写操作
func go_file_write() {
	// go_file_write1()
	// go_file_write2()
	go_file_write3()
}

//通过字节数组写
func go_file_write1() {
	file, _ := os.OpenFile("./qiu.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	//直接使用类型转换
	file.Write([]byte("dsdsdassdad"))

	data := make([]byte, 100)
	s := "hello world lks and qiu"
	data = []byte(s)
	file.Write(data)
	file.Close()
}
//直接写入字符串
func go_file_write2() {
	file, _ := os.OpenFile("./qiu.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	file.WriteString("hello qiumonster and lks love love")
	file.Close()
}
//随机写 可以实现文件的续写操作
func go_file_write3(){
	file, _ := os.OpenFile("./qiu.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	//文件在4的位置进行续写
	file.WriteAt([]byte("hello lks and qiu and..."),4)
	file.Close()
}