package main

import (
	"fmt"
	"strconv"
)

//func main() {
//	var str1 = "asdffaadf"
//	for i:=0;i<len(str1);i++ {
//		fmt.Println(str1[i])
//	}
//}

// 使用ASCII 值来解决问题
func main() {
	var str string
	fmt.Scanf("%s", &str)
	var result int
	for i:=0;i<len(str);i++ {
		num := int(str[i] - '0')
		result += (num *num*num)
	}
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("can not convert %s to int\n", str)
	}
	if result == number {
		fmt.Printf("%d is 水仙花数", number)
	}else{
		fmt.Printf("%d is not 水仙花数", number)
	}

}
