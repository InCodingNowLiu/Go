package main

import (
	"fmt"
	"strconv"
	"strings"
)

// string replace function
func main() {
	str := "  hello world abc   \n"
 	// strings.Replace 字符串替换
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println("repalce: ",result)

	// strings.Count 字符串计数
	count := strings.Count(str, "l")
	fmt.Println("count:", count)

	// strings.Repeat 重复count次str
	result = strings.Repeat(str, 3)
	fmt.Println("repeat:", result)

	// strings.ToLower 转为小写
	result = strings.ToLower(str)
	fmt.Println("to lower", result)

	// strings.ToUpper 转为大写
	result = strings.ToUpper(str)
	fmt.Println("to upper", result)

	// strings.TrimSpace 去掉字符串首尾空白字符
	result = strings.TrimSpace(str)
	fmt.Println("trim space", result)

	// strings.Trim 去掉字符串首尾cut字符
	result = strings.Trim(str, "\n\r")
	fmt.Println("trim", result)

	// strings.trimLeft
	result = strings.TrimLeft(str, " \n\r")
	fmt.Println("trim left", result)


	// strings.TrimRight 去掉字符串首cut字符
	result = strings.TrimRight(str, " \n\r")
	fmt.Println("trim right", result)

	// strings.Field 返回str空格分隔的所有的子串的slice
	splitResult := strings.Fields(str)
	for i:=0;i<len(splitResult);i++ {
		fmt.Println("split element", splitResult[i])
	}

	// strings.Split 返回str split 分隔的所有子串的slice
	splitResult = strings.Split(str, "l")
	for i:=0;i<len(splitResult);i++ {
		fmt.Println(splitResult[i])
	}

	// strings.Join 用sep把s1中的所有元素链接起来
	str2 := strings.Join(splitResult, "l")
	fmt.Println("join:", str2)

	// strings.Itoa 把一个整数转成字符串
	str2 = strconv.Itoa(1000)
	fmt.Println("Itoa:", str2)

	// strings.Atoi 把一个字符串转成整数
	number, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("can not convert string to int")
	}
	fmt.Println("number:", number)
}