package day1

import "fmt"

// module.export
var Test int

// struct
type Student struct {
	name string
	age int
}


type Teacher struct {
	name string
	age int
}



func main() {
	var test1 Student = Student{
		name: "jiqing",
		age: 10,
	}
	var test2 map[string]string = map[string]string{
		"key": "jiqing",
		"value": "sos",
	}
	fmt.Println(test1)
}