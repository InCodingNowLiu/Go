package main

import "fmt"

type add_func func(int, int) int

func sub(a,b int) int {
	return a-b
}

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

func add(a,b,c int) int {
	return a+b
}

func main(){
	var test add_func
	c := sub
	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Println(sum)

}