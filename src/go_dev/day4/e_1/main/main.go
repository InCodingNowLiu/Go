package main

import (
	"errors"
	"fmt"
)

/*
	close: 主要用来关闭channel
	len: 用来求长度, 比如string, array, slice, map, channel
	new: 用来分配内存, 主要用来分配值类型, 比如int,struct, 返回的是指针
	make: 用来分配内存, 主要用来分配引用类型, 比如chan, map, slice
	append: 用来追加元素到数组, slice中
	panic和recover: 用来做错误处理
*/

func test() {
	//defer func() {
	//	if err := recover(); err !=nil {
	//		fmt.Println(err)
	//	}
	//}()

	err := initConfig()
	if err != nil {
		panic(err)
	}
	return

	b := 0
	a := 100/b
	fmt.Println(a)
	return
}

func main() {
	var i int
	fmt.Println(i)

	j:= new(int)
	*j = 100
	fmt.Println(*j)

	// 定义数组
	var a []int
	a = append(a, 10, 20,303)
	a = append(a, a...)
	fmt.Println(a)
	test()
}

func initConfig() (err error) {
	return errors.New("init config failed")
}