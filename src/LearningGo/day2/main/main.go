//package main
//
//import "fmt"
//
//type student struct {
//	name string
//	age int
//}
//
//func main() {
//	var b int = 1
//	//var c student = student{
//	//	name: "jiqing",
//	//	age: 18,
//	//}
//	fmt.Println(b)
//	changeValue(&b)
//	fmt.Println(b)
//
//	// & 取地址
//	// *(type) 指针变量类型
//	// *(*int) 取指针指向的值
//}
//
//func changeValue(b *int) {
//	*b = *b + 1
//}