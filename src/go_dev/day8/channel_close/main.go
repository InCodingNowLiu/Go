package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0;i<10;i++ {
		ch <- i
	}
	close(ch)
	for {
		var b int
		b, ok := <-ch
		if ok == false {
			fmt.Println("chann is close")
			break
		}
		fmt.Println(b)
	}
}
