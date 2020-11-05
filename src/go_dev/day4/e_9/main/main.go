package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1,8,38,2,345,2}
	// 要改变值需要传一个切片进去
	sort.Ints(a[:])
	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "as","sdfr", "asr", "gwe"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.2134, 2123.123, 0,123, 12.123, 52.123}
	sort.Float64s(a[:])
	fmt.Println(a)
}

// 必须先进行排序
func testIntSearch()  {
	var a = [...]int{1,8,38,2,345,2}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 38)
	fmt.Println(index)
}


func main()  {
	testIntSort()
	testStrings()
	testFloat()
	testIntSearch()
}