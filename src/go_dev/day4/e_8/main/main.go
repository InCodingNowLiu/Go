package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1,2,3,4,5}
	s := a[1:]
	s[1] = 100
	fmt.Println(a)
	s[2] = 200
	fmt.Println(a)
}

func testCopy() {
	var a []int = []int{1,2,3,4,5,6,7,8,9}
	b := make([]int, 1)
	copy(b,a)
	fmt.Println(b)
}

func testString() {
	s := "hello world"
	s1 := s[0:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)
}

func testModifyString() {
	s := "hello world"
	s1 := []rune(s)

	s1[1] = '0'
	str := string(s1)
	fmt.Println(str)
}

func main() {
	testSlice()
	testCopy()
	testString()

	// string 底层就是一个byte数组, 因此也可以进行切片操作
	// 字符串在golang里面是不可变的
	testModifyString()
}
