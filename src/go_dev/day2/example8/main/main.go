package main

import (
	"fmt"
	)

func modify1(a *int) {
	*a = 10
}
func modify(a int) {
	a = 10
}

func main()  {
	test1 := 5
	test2 := make(chan int, 1)
	fmt.Println("test1=", test1)
	fmt.Println("test2=", test2)
	modify(test1)
	fmt.Println("test1=", test1)
	modify1(&test1)
	fmt.Println("test1=", test1)
}