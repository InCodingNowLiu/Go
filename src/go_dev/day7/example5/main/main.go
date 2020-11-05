package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main() {
	// 如何产看open的路径
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(`get relate path error: `, err)
		return
	}
	file, err := os.Open(pwd + "/src/go_dev/day7/example4/main/test.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v", err)
			break
		}

		runeArr := []rune(str)
		for _, v := range runeArr {
			switch  {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v>= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)
}