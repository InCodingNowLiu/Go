package main

import (
	"fmt"
	"sort"
)

func testMapSort1() {
	var a map[string]int
	var b map[int]string
	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["abc"] = 101
	a["efg"] = 10
	for k,v := range a {
		b[v] = k
	}
	fmt.Println(b)
}

// map 里面的key都是无序的
func main() {
	testMapSort()
	testMapSort1()
}

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[8] =10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for k,v := range a {
		fmt.Println(k,v)
	}

	var keys []int
	for k,_ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _,v:=range keys {
		fmt.Println(v, a[v])
	}
}

