package main

import "fmt"

func main()  {
	var str1 = "hello"
	str2 := "world"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	n := len(str3)
	fmt.Println(str3)
	fmt.Printf("len(str3)=%d\n",n)
	subStr := str3[0:5]
	fmt.Println(subStr)
	// 切片
	subStr = str3[6:]
	fmt.Println(subStr)

	reverse("hello world")
	result := reverse1("hello world")
	fmt.Println(result)
}

func reverse(str string) string {
	var ouputStr string
	length := len(str)
	for i:=0;i<length;i++ {
		var str1 = str[length-i-1]
		ouputStr = ouputStr + fmt.Sprintf("%c", str1);
	}
	fmt.Println(ouputStr)
	return ouputStr
}

func reverse1(str string) string {
	var result []byte
	length := len(str)
	 temp := []byte(str)
	for i:=0;i<length;i++ {
		result = append(result, temp[length-i-1])
	}
	return string(result)
}