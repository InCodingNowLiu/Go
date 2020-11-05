package main

import "fmt"

type Student struct {
	Name string
	Age int
	score float32
}

func main() {
	var stu Student
	fmt.Println(stu)
	stu.Age = 18
	stu.Name = "hua"
	stu.score = 80

	var stu1 *Student = &Student{
		Name: "1",
	}
	fmt.Println(stu1)
	fmt.Println(stu)
}