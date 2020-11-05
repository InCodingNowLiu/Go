package main

import "fmt"

func test1() {
	var a [10]int
	a[0] = 100

	fmt.Println(a)

	for i:= 0;i<len(a);i++ {
		fmt.Println(a[i])
	}
	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func main()  {
	var a [10]int

	j := 10
	a[0] = 100
	a[j] = 200
	fmt.Println(a)
}