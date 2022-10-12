package main

import (
	"fmt"
	"os"
)

func go_os() {

	// //创建一个字节数据切片
	// data := make([]byte, 100)
	// file, _ := os.Open("../file/qiu.txt")
	// count, _ := file.Read(data) //1 字节数 2 错误
	// fmt.Println("文件总字节数:", count)
	// // fmt.Println(os.)

	// go_os_1()
	// go_os_2()
	// go_os_3()
	// go_os_4()
	// go_os_5()
	// go_os_6()
	go_os_7()

}

//创建文件
func go_os_1() {
	file, err := os.Create("../file/qiu1.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("输出文件名称:", file.Name()) //输出创建的全部名称 包括目录
	}
}

//创建文件夹或目录
func go_os_2() {
	//创建一个文件夹
	err := os.Mkdir("../test", os.ModePerm) //1 文件夹名称 2 权限 os.ModePerm 为 Linux中的最高权限
	if err != nil {
		fmt.Println(err)
	}
	//创建级联的文件夹
	os.MkdirAll("../test/testa", os.ModePerm)
}

//删除文件或文件夹
func go_os_3() {
	//删除一个问价或文件夹
	os.Remove("../file/qiu1.txt")

	//删除test目录下的所有内容
	os.RemoveAll("../test")
}

//工作目录相关操作
func go_os_4() {
	dir, _ := os.Getwd() //当前工工作空间的完整路径

	fmt.Println("当前的工作目录是:", dir)
	// os.Chdir("d:/") //修改当前的工作目录
	s := os.TempDir() //获得当前的缓存目录
	fmt.Println("当前的缓存目录:", s)
}

//重命名文件
func go_os_5() {
	//文件名的修改
	os.Rename("../file/qiu.txt", "../file/qiu_copy.txt") //1 原本的文件名 2 修改后的文件名
	os.Mkdir("../test", os.ModePerm)
	//重命名文件夹
	os.Rename("../test", "../test_copy")
}

//读文件 输出文件中的字符串
func go_os_6() {
	b, _ := os.ReadFile("../file/qiu_copy.txt")
	//使用类型转换 将字节转为字符串 进行输出
	fmt.Println(string(b))
}

//写文件
func go_os_7() {
	s := "hello qiu and lks"
	//[]byte(s) 将字符串转为字节
	os.WriteFile("../file/qiu_copy.txt", []byte(s), os.ModePerm) //1 文件名 2 数据字节数组 3 权限
}
