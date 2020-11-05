package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/jliu396/go/src/awesomeProject/src/go_dev/day7/example4/main/test.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("read string failed, err:", err)
		return
	}
	fmt.Printf("read str success, ret: %s \n", str)
}