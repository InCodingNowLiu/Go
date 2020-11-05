package main

import "fmt"

// 编写程序，在终端输出九九乘法表
func questionOne() {
	for i:=1;i<=9;i++ {
		for j:=1;j<=i;j++ {
			fmt.Printf("%d*%d=%d ", j,i, i*j)
		}
		fmt.Println()
	}
}

//一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
//编程找出1000以内的所有完数。
func questionTwo() {
	var n int
	fmt.Scanf("%d", &n)
	var element = func(a int) bool {
		var sum int = 0
		for i:=1;i<a;i++ {
			if a%i == 0 {
				sum += i
			}
		}
		if (sum == a) {
			return true
		}else{
			return false
		}
	}
	// 完数的定义 , 除了自己能被自己整除的数
	var result []int
	for i:=1;i<=n;i++ {
		if (element(i)) {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}

//输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
//左读完全相同的字符串。
func questionThree() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println("输入的值是: ", input)
	length := len(input)
	if length %2 ==0 {
		length = length / 2
	}else{
		length = length / 2 + 1
	}
	var flag = true
	for i:=0;i<length;i++ {
		if input[i] != input[len(input)-1-i] {
			flag = false
		}
	}
	if flag {
		fmt.Println("It is huiwen")
	}else{
		fmt.Println("It's not huiwen")
	}
 }


 //输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数
 func questionFour() {
	// 此题暂时不会
 }

 //计算两个大数相加的和，这两个大数会超过int64的表示范围
 func questionFive() {

 }


func main() {
	//questionOne()
	//questionTwo()
	questionThree()
}
