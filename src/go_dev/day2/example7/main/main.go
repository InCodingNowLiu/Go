package main

import (
	"os"
	"fmt"
)

func main()  {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s", path)
}