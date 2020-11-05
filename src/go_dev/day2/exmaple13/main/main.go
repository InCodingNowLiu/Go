package main

import "fmt"

func main()  {
	var n int16 = 34
	var m int32
	//m=n
	m = int32(n)
	fmt.Printf("m=%d, n = %d\n", m, n)

	var a int = 19
	if (a > 10) {
		fmt.Println(a)
	}
}