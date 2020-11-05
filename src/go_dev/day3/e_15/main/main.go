package main

import "fmt"

func add(a int, arg...int) int {
	var sum int = a
	for i:=0;i<len(arg);i++ {
		sum += arg[i]
	}
	return sum
}

func main() {
	sum := add(1, 2,3,4)
	fmt.Println(sum)

	result := concat("hello", "world", "中国")
	fmt.Println(result)
}

func concat(a string, arg...string)(result string) {
	result = a
	for i:=0;i<len(arg);i++ {
		result += arg[i]
	}
	return
}

// defer 使用 如果申明二个defer 第一个defer1 第二个defer2 那么第二个先执行, 第一个后执行

