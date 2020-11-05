package main

import "fmt"

type student struct {
	name string
	age int
}

func main() {
	test := student{
		name: "jiqing",
		age: 18,
	}
	//fmt.Println("操作之前地址:", &test)
	fmt.Printf("%p\n", &test)
	changeValue(&test)

	fmt.Println("操作之后:", test)
}

func changeValue(b *student) {
	//fmt.Printf("changeValue: %p", &b)

	//(*b).name = "dick"
	v := 2
	if b.name == "dick" {
		b.name = "dick"
		fmt.Println(v)
	}

	path := "../../day2"
	a := readFile(path)

}

func readFile(path string) (int, bool) {
	fmt.Println(a, b)
	 return
}