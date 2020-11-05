package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10000)
	ch2 := make(chan int, 10000)
	//go func() {
	//	var i int
	//	for {
	//		ch <- i
	//		time.Sleep(time.Second)
	//		ch2 <- i * i
	//		time.Sleep(time.Second)
	//		i++
	//	}
	//}()

	for i:=0;i<10000; i++  {
		ch <- i
		ch2 <- i*i
		fmt.Println(i)
	}


	for {
		select {
		case v := <-ch:
			fmt.Println(v)
			case v := <- ch2:
				fmt.Println(v)
				// 如果一秒之后没值就超时
				case <-time.After(time.Second):
					fmt.Println("get data timeout")
					//time.Sleep(time.Second)
		}
	}
}