package main

import "fmt"

// map slice chan 指针 interface 默认以引用的方式传递

func main() {
	sum, _ := calc(100,200)
	fmt.Println(sum)
}

func add(a,b int) (c int) {
	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a+b
	avg = (a+b)/2
	return
}

// 可变参数

// 0个或者多个参数
func add(arg...int) {

}

// 1个或者多个参数
func add(a int, arg...int) {

}
// 2个或者多个参数
func add(a int, b int, arg...int) {

}