package main

import (
	"fmt"
	"log"
	"os"
)

// 三个级别的日志打印
//1 print 单纯的打印日志
// 2 panic 打印日志 抛出panic异常
// 3 fatal 打印日志 强制结束程序 os.exit(1) defer函数不会执行
func go_log() {
	// go_log1()
	// go_log2()
	// go_log3()
	go_log4()
}
func go_log1() {
	//简单的日志打印
	log.Print("这里是日志打印")
	log.Printf("%d日志输出", 100)
	log.Println("输出", "日志")
}
func go_log2() {
	log.Panic("panic")
}
func go_log3() {
	defer fmt.Println("我要执行...")
	log.Fatal("直接结束程序，不执行 defer")
}
func go_log4() {
	//添加日志输出的格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//设置日志的前缀
	log.SetPrefix("Qiu-Log:")
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	
	//将日志写入文件中去
	f, err := os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err!=nil{
		log.Fatal("日志文件错误")
	}
	//输出当前的日志数据
	log.SetOutput(f)
	//先设置日志输出位置再 输出日志数据
	log.Print("新的log输出格式...")
}
