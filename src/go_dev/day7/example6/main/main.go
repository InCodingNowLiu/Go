package main

import (
	"fmt"
	"os"
)

func main()  {
	// os执行命令的参数, 以空格分隔
	fmt.Printf("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}
}
