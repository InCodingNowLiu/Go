package main

import "fmt"

func test() {
	s1 := new([]int)
	fmt.Println(s1)

	s2 := make([]int, 10)
	fmt.Println(s2)
	return
}

func factor(n int) int {
	if n == 1{
		return 1
	}
	return factor(n-1)*n
}

func main() {
	test()
}

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

// 斐波那契数
