package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//首字母必须大写 才能实现json对象的转换
type Student struct {
	Name string
	Age  int
}

func go_json() {
	// go_json1()
	// go_json2()
	// go_json3()
	go_json4()
}
//读取json文件
func go_json4() {
	f, _ := os.Open("qiu.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var v map[string]interface{} //key string value any

	d.Decode(&v)

	fmt.Println(v)
	fmt.Println(v["Name"])
}

//嵌套类型的解析
func go_json3() {
	s := []byte(`{"Name":"qiuandlks","Children":{"Name":"wls","Age":1},"Age":23}`)
	//使用了一个对象
	var p interface{}
	json.Unmarshal(s, &p)
	fmt.Println(p)
}

//将json转为结构体
func go_json2() {
	s := []byte(`{"Name":"qiu","Age":23}`)
	var p Student
	json.Unmarshal(s, &p) //注意转换的时候应该使用结构体对象的地址
	fmt.Println(p)

}

//将结构体转为json
func go_json1() {
	s := Student{"lks", 22}
	b, _ := json.Marshal(s)
	// for _,val:=range b{
	// 	fmt.Println(string(val))
	// }
	fmt.Println(string(b))
}
