package main

import "fmt"

func main()  {
	var test1 int
	var test2 bool
	var c byte = 'a'
	d := 'a'
	fmt.Printf("%v\n", test1)
	fmt.Printf("%v\n", test2)
	fmt.Printf("%v\n", c)
	fmt.Printf("%#v\n", test2)
	fmt.Printf("%T\n", d)
	// 输入bool的值类型
	fmt.Printf("%t\n", test2)
	// 输出二进制
	fmt.Printf("%b\n", 100)
	// 输入浮点型
	fmt.Printf("%f\n", 199.02)
	// 输出双引号的字符串
	fmt.Printf("%q\n", "this is a tesrt")
	// 输出16进制
	fmt.Printf("%x\n", 39123123)
	// 输出指针地址n
	fmt.Printf("%p\n", &test1)
	str := fmt.Sprintf("%d", 100)
	fmt.Printf("%q\n", str)
	fmt.Printf("%T", str)
}