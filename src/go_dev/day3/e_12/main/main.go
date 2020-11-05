package main

import "fmt"

func Print(n int) {
	for i:=0;i<n+1;i++ {
		for j:=0;j<i;j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
	str := "hello world, 中国"
	for i,v := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n",i,v,len(string(v)))
	}
}

func main() {
	Print(6)
}