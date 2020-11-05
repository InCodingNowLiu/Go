package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var str string
	fmt.Scanf("%s", &str)
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("convert failed, err:", err)
	}
	fmt.Println(number)

	// switch case
	var a int = 0;
	switch a {
	case 0:
		fmt.Println("a is equal 0")
		fmt.Println("yes ")
		// 下面代码是穿透作用
		fallthrough
	case 10:
		fmt.Println("a is equal 10")
	default:
		fmt.Println("a is equal default")
	}
}