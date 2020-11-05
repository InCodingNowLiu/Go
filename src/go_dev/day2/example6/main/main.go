package main

import (
	_ "fmt"
	_ "time"
)
const (
	man = 1
	female = 2
)

var (
	test1 int // 默认值是0
	test2 string // 默认值是""
	test3 bool // 默认值是false
	test4 int = 1s
	test5 string = "hello world"
	test6 bool = true
)

func main() {
	// 死循环
	//for {
	//	second := time.Now().Unix()
	//	if (second % female == 0) {
	//		fmt.Println("female")
	//	}else{
	//		fmt.Println("man")
	//	}
	//	time.Sleep(100 * time.Microsecond)
	//}
}