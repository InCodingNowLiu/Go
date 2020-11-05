package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	//fmt.Println(int(math.Sqrt(float64(n))))
	//for i:=2;i<int(math.Sqrt(float64((n))));i++ {
	for i := 2;i<int(math.Sqrt(float64(n)));i++ {
		if n%i ==0 {
			return false
		}
	}
	return true
	//return false
}


func  main()  {
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m)
	for i := n;i<m;i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}

// 优化之后的代码
func isPrimise1()  {

}