package main

import "fmt"

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr= new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	fmt.Println(s.ptr)
	(*s.ptr)[1] = 1000
}

func testSlice() {
	//var slice []int
	//var arr [5]int = [...]int{1,2,3,4,5}
	//slice = arr[2:5]
	//fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//slice = slice[0:1]
	//fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//slice = arr[:]
	//fmt.Println(slice)
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	var test1 [100]int
	var test2 *[100]int
	fmt.Println(test1, test2)
	var test3 slice
	make1(test3,100)
}

func main() {
	//testSlice()
	//var test3 slice
	//test3 = make1(test3,100)
	//fmt.Println(test3.ptr)
	//(*test3.ptr)[0] = 100
	//modify(test3)
	//fmt.Println(*test3.ptr)

	//testSlice1()
	testSlice4()
}

func changeValue(a *[10]int) {
	a[0] = 100
	return
}

func testSlice4() {
	var a = [10]int{1,2,3,4}
	b := a[1:5]
	fmt.Printf("%p\n",b)
	fmt.Printf("%p\n", &a[1])
}

func testSlice1() {
	var a *[10]int
	var b = [10]int{1:1, 2:2}

	a = &b
	changeValue(a)
	//a[0] = 1
	fmt.Println(a)
}