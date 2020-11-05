package main

import "fmt"

type integer int

type Student struct {
	Number int
	hello string
}

type Stu Student // alias

func main() {
	var i integer = 1000
	var j int = 100
	j = int(i)
	fmt.Println(j)

	var a Student
	a = Student{30,"1"}
	fmt.Println(a)
}