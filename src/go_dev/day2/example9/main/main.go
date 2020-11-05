package main

import "fmt"

//test11 := 100

//var test11 int = 100
func temp(a int, b int)  {
	temp := a
	a = b
	b = temp
	return
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	return
}

func swap2(a int, b int) (int, int) {
	return b, a
}

func test() {
	var a = 100
	fmt.Println(a)
	for i:=0; i< 100;i++ {
		var b = i*2
		fmt.Println(b)
	}
	//fmt.Println(b)
}

func test23() {
	var test21 int8 = 100
	var b int16 = int16(test21)
	fmt.Printf("test21=%d b=%d\n", test21, b)
}

func main()  {
	first := 100
	second := 200
	temp(first, second)
	fmt.Println("first=", first)
	fmt.Println("second=", second)
	swap(&first, &second)
	fmt.Println("first=", first)
	fmt.Println("second=", second)
	first, second = swap2(first, second)
	fmt.Println("first=", first)
	fmt.Println("second=", second)

}
var test2 = "G"
func n() {
	fmt.Println(test2)
}

func m() {
	test2 := "O"
	fmt.Println(test2)
}