package main

import (
	"fmt"
	"io"
	"os"
)

func go_file_read() {
	// go_file1()
	// readOps()
	readDir1()
}
func go_file1() {
	//使用open打开的文件是只读的  实际上调用的就是 openfile函数
	f, _ := os.Open("../file/qiu_copy.txt")
	fmt.Println("输出打开的文件名:", f.Name())
	f.Close() //关闭文件

	//可执行的文件  os.O_CREATE 当文件不存在时 会自动创建
	f2, _ := os.OpenFile("../file/qiu_copy.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Println(f2.Name())
	f2.Close()
}

//文件读取  当文件过大时 通过for循环进行文件读取
func readOps() {
	// file, _ := os.OpenFile("../file/qiu_copy.txt",os.O_RDWR,755)
	file, _ := os.Open("../file/qiu_copy.txt")
	s := ""

	data := make([]byte, 10)
	file.ReadAt(data, 5) //从文件的第5个位置读取文件
	fmt.Println("从第5个位置读取文件数据:", string(data))

	for {
		buf := make([]byte, 10)
		n, err := file.Read(buf) //将文件读到字节数组中

		//当文件读取到最后时 就结束循环
		if err == io.EOF {
			break
		}
		fmt.Println("文件长度:", n)
		fmt.Println("文件的内容为:", string(buf))
		s += string(buf)
	}
	fmt.Println("文件的最终内容:", s)

	file.Seek(6, 0)
	file.Read(data)
	fmt.Println("使用文件定位操作:", string(data))
	file.Close()
}

//文件夹的读取与内部文件的遍历
func readDir1() {
	de, _ := os.ReadDir("../main")
	for _, val := range de {
		if val.IsDir() {
			fmt.Println("为文件夹:", val.Name())
		}
		if !val.IsDir() {
			fmt.Println("为文件:", val.Name())
		}
	}
}
