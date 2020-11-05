package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"student_name"`
	Age int `json:"age"`
	Score int `json:"score"`
}

// 私有属性外界获取不到
type Test struct {
	name string
	Age int
}

func main() {
	var stu Student = Student{"stu01", 10, 80}
	var test Test = Test{
		"test",10,
	}
	data, err := json.Marshal(stu)

	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}
	fmt.Println(string(data))

	data1, err1 := json.Marshal(test)
	if err1 != nil {
		fmt.Println("json encode stu failed, err:", err1)
	}
	fmt.Println(string(data1))
}