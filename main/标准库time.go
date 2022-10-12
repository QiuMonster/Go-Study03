package main

import (
	"fmt"
	"time"
)

func go_time() {
	t := time.Now()
	fmt.Println(t.Local().Date())
	fmt.Println(t.Year(), t.Month())
	temp := t.Unix()
	fmt.Println("当前的时间戳:",temp)//从1970 1 1 开始计算
	t2 := time.Unix(temp, 0)//转换为当前的时间
	fmt.Println("当前的时间:",t2)
}
