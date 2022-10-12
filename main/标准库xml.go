package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

//将结构体转为xml标签
func go_xml() {
	// go_xml1()
	go_xml2()

}
func go_xml1() {
	p := Person{"qiu", 22}
	b, _ := xml.Marshal(&p) //直接使用对象的引用
	fmt.Println(string(b))
}

//将xml变为结构体类型
func go_xml2() {
	s := []byte(`<Person><Name>qiu and lks</Name><Age>12.99</Age></Person>`)
	var p Person //表示这是一个任意类型
	xml.Unmarshal(s, &p)
	fmt.Println(p)
}
