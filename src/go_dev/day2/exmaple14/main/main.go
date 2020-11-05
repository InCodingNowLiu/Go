package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 放在init函数中
func main()  {
	rand.Seed(time.Now().UnixNano())
}

func main()  {
	// 随机生成一个种子, 来保证每次random的都不一样
	rand.Seed(time.Now().Unix()) // 当前时间的秒数, 单位秒
	rand.Seed(time.Now().UnixNano()) // 纳秒

	// 定义一个字符
	var a1 byte = 'a'
	var test1 string = "ab"
	var test2 string = `ad
ad
ads`
 	fmt.Println(a1, test1, test2)

	for i := 0;i<10;i++ {
		a := rand.Int()
		fmt.Println(a)
	}
	for i:= 0;i<10;i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	for i := 0;i<10;i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}