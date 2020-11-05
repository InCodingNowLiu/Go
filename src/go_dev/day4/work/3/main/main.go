package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func add(str1, str2 string) (result string) {
	if len(str1)==0 && len(str2)==0{
		result = "0"
		return
	}
	index1 := len(str1) -1
	index2 := len(str2) -1
	var left int
	for index1 >= 0 && index2 >= 0 {
		num1 := str1[index1] - '0'
		num2 := str2[index2] - '0'
		sum := int(num1) + int(num2) + left
		if sum >= 10 {
			left = 1
		}else{
			left = 0
		}
		num3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", num3,result)
		// 两者相加
		index1--
		index2--
	}
	for index2 >= 0 {
		num2 := str2[index2] - '0'
		sum := int(num2) + left
		if sum >= 10 {
			left = 1
		}else{
			left = 0
		}
		num3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", num3,result)
		// 两者相加
		index2--
	}
	for index1 >= 0 {
		num1 := str1[index1] - '0'
		sum := int(num1) + left
		if sum >= 10 {
			left = 1
		}else{
			left = 0
		}
		num3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", num3,result)
		// 两者相加
		index1--
	}
	if left != 0 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}

func main()  {
	reader := bufio.NewReader(os.Stdin)

	result,_,err := reader.ReadLine()
	if err != nil {
		fmt.Println("Read from console err", err)
	}
	test1 := strings.Split(string(result), "+")
	if len(test1) != 2 {
		fmt.Println("please input format like a+b")
	}

	// 干掉space
	str1 := strings.TrimSpace(test1[0])
	str2 := strings.TrimSpace(test1[1])

	res := add(str1, str2)
	fmt.Println(res)

}