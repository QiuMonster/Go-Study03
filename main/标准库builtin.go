package main

import "fmt"

// func main() {
// 	go_builtin()
// }

//可以不用显示的导入包
func go_builtin() {
	// go_b1()
	// go_b2()
	go_b3()
}
func go_b1() {
	//定义切片
	s := []int{1}
	s = append(s, 100)
	fmt.Printf("s: %v\n", s)
	s1 := []int{1, 2}
	//一个切片添加另一个切片时需要添加 ... 进行解包操作
	s1 = append(s1, s...)
	fmt.Println(s1)

}

//使用new创建对象时 返回的是对象的指针
//使用make创建时 返回对象引用 也就是对象本身
func go_b2() {
	a := new(int) //默认是0值
	// *a=0
	fmt.Printf("a: %T\n", a)
	fmt.Println(a, *a, &a)
	b := new(bool)
	fmt.Println(b, *b)
}
func go_b3() {
	//使用new来创建一个int数组
	var p *[]int = new([]int)
	fmt.Println(*p)
	//使用切片来创建一个int数组
	arr := make([]int, 10, 100)
	fmt.Println(arr)
}
