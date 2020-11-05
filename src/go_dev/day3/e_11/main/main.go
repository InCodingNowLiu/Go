package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var n int
	n = rand.Intn(100)
	// for 死循环让用户一直输入
	for {
		var input int
		// \n 分隔符
		fmt.Scanf("%d\n", &input)
		flag := false
		switch  {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input < n:
			fmt.Println("less")
		case input > n:
			fmt.Println("bigger")
		}
		if flag {
			break
		}
	}
}