package main

import "fmt"

func main()  {
	b := [...]int{8,7,5,4,3,10,15}
	qsort(b[:], 0, len(b)-1)
	fmt.Println(b)
}



// 实现一个冒泡排序
// [1, 123,43, 12313,41254,312412,12314]
func maopao() {
	a := []int{
		4,5,2,6,3,7,8,10,9,
	}
	for i:=0;i<len(a);i++ {
		for j:=1;j<len(a) - i;j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// 实现一个选择排序
// [7,8,2,3,5]
func ssort(a []int) {
	// [2,8,7,3,5]
	for i:=0;i<len(a);i++ {
		var min int = i
		for j := i+1;j<len(a);j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}


// 实现一个插入排序
func isort(a []int) {
	for i:=1;i<len(a);i++ {
		for j:=i;j>0;j--{
			if a[j] > a[j-1] {
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}


// 实现一个快速排序
// [5,3,2,8,7,6]
// 8,7,5,4,3,10,15
// 0 7
func qsort(a []int, left, right int) {
	// 0 7
	if left >= right {
		return
	}
	val := a[left] // 8
	k := left // 0
	// 确定val所在的位置
	// i 2
	for i:= left + 1; i<=right;i++ {
		if a[i] < val {
			// a[1] = 7 < 8

			// a[0] = a[1] = 7
			a[k] = a[i]

			// a[1] = a[1] = 7
			a[i] = a[k+1]

			// k =  2
			k++
		}
	}
	// a[1] = 8
	a[k] = val
	qsort(a, left, k-1)
	qsort(a, k+1, right)
}