package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	// 放在这里前序遍历
	trans(root.left)
	// 放在这里中序遍历
	fmt.Println(root)
	//放在最后, 是后序遍历
	trans(root.right)

}



func main() {
	var root *Student = new(Student)
	root.Name = "stu01"
	root.Age = 18
	root.Score = 100

	var left1 *Student = new(Student)
	left1.Name = "left"
	left1.Age = 20
	left1.Score = 90
	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "stu01"
	right1.Age = 21
	right1.Score = 99
	root.right = right1

	var left2 *Student = &Student{
		Name: "left2",
		Age: 23,
		Score: 89,
	}
	left1.left = left2
	trans(root)

}