package main

import "fmt"

func main()  {
	var str1 = "hellow world\n"
	var str2 = `
	床前明月光,
	疑是地上霜.
	举头望明月,
	低头思故乡.
	`
	var b byte = 'c'
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(b)
	fmt.Printf("%c\n", b)

}
