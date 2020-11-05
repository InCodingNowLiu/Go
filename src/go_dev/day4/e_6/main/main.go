package main

import "fmt"

func fab(n int)  {
	var a []int64
	a = make([]int64, n)

	a[0] = 1
	a[1] = 1
	for i :=2;i<n;i++ {
		a[i] = a[i-1] + a[i-2]
	}
	for _, v := range a {
		fmt.Println(v)
	}
}


func main()  {
	var age [5]int = [5]int{1,2,3,4,5}
	var age1 = [5]int{1,2,3,4,5}
	fab(10)
	var age2 = [...]int{1,2,3,4,5}
	var str = [5]string{3: "hello world", 4:"tom"}

	fmt.Println(age, age1, age2, str)

	// 多维数组
	var age23 [5][3]int
	var f [2][3]int = [...][3]int{{1,2,3}, {2,3,4}}
	fmt.Println(age23, f)
	testArray2()
}

func testArray2() {
	var a [2][5]int = [...][5]int{{1,2,3,4,5}, {6,7,8,9,10}}
	for row,v := range a {
		for col, v1 := range v {
			fmt.Printf("(%d,%d)=%d", row, col, v1)
		}
		fmt.Println()
	}
}

