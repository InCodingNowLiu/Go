package main

import "fmt"

type Student struct {
	Age int
	Name string
	
}

type StudentArray []Student

func main() {
	var test StudentArray
	fmt.Println(test)
	test = []Student{
		{Age: 1, Name: "age1"},
		{Age: 2, Name: "age2"},
	}
	fmt.Println(test)
}