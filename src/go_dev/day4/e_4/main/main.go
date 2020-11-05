package main

import (
	"fmt"
	"strings"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		 x+=d
		 return x
	}
}

func makeSuffixFunc(suffix string) func(string2 string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f:= Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

	func1 := makeSuffixFunc(".jpg")
	func2 := makeSuffixFunc(".bmp")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}