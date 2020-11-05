package main

import (
	a "awesomeProject/src/go_dev/day2/example2/add"
	"fmt"
)

func main()  {
	const (
		what = 0
		b = 1
		c = 2
	)
	const (
		ab = iota
		ac
		ad
	)
	fmt.Println(ac)
	//a.Test()
	fmt.Println("Name=", a.Name)
	fmt.Println("age=", a.Age)
}