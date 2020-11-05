package main

import "fmt"

func main()  {
	result := caculateOne()
	fmt.Println(result)
	result = caculateTwo()
	fmt.Println(result)
	result1 := caculateThree(3)
	fmt.Println(result1)
}

func caculateOne() []int {
	var result []int
	for i:= 101;i<=200;i++ {
		isPrimeNumber := true
		for j:= 2;j<i;j++ {
			if (i%j == 0) {
				isPrimeNumber = false
			}
		}
		if (isPrimeNumber) {
			result = append(result, i)
		}
	}
	return result
}

func caculateTwo() []int {
	var result []int
	for i := 100; i <= 999; i++ {
		gewei := i % 10
		shiwei := (i % 100) / 10
		baiwei := i / 100
		temp := gewei*gewei*gewei+shiwei*shiwei*shiwei+baiwei*baiwei*baiwei
		fmt.Println(temp)
		if (temp == i) {
			result = append(result, i)
		}
	}
	return result
}

func caculateThree(n int) int {
	var total int
	for i:=1;i<=n;i++ {
		total = total + jiechen(i)
	}
	return total
}

func jiechen(n int) int {
	total := 1
	for i:=1;i<=n;i++ {
		total = total *i
	}
	return total
}

//var total int
//// 递归
//func digui(n int) int {
//	if (n == 1){
//		return 1
//	}else{
//		total +=
//	}
//	digui(n-1)
//}