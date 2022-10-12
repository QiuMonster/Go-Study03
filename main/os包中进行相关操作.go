package main

import (
	"fmt"
	"os"
)

//os进程相关
func go_process() {
	go_process1()
}
func go_process1() {
	//获得当前正在运行进程的pid
	i := os.Getpid()
	fmt.Println("当前进程的pid:", i)
	fmt.Println("当前进程的父进程的pid:", os.Getppid())
}
