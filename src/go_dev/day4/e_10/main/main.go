package main

import "fmt"

func testMap() {
	//var a map[string]string
	// 这个size 根据数据的格式预估, 这样增加性能不会自动扩容
	//a = make(map[string]string, 2)
	//a["12"] = "12"
	//a["34"] = "35"
	//a["ad"] = "dc"
	//a["ad1"] = "dc1"
	//a["ad2"] = "dc2"
	//a["ad3"] = "dc3"

	var a map[string]string = map[string]string{
		"key": "value",
	}
	a["12"] = "12"
	a["34"] = "34"
	fmt.Println(a)
}

func testMap2() {
	a := make(map[string]map[string]string, 10)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"
	fmt.Println(a)
}

func modify(a map[string]map[string]string) {
	//val, ok := a["zhangsan"]
	//if ok {
	//	val["username"] = "123456"
	//	val["password"] = "123456"
	//}else {
	//	a["zhangsan"] = make(map[string]string)
	//	a["zhangsan"]["password"] = "123456"
	//	a["zhangsan"]["username"] = "123456"
	//}

	// 优化代码如下
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["password"] = "123456"
	a["zhangsan"]["username"] = "123456"
	return

}

func main() {
	//testMap()
	//testMap2()
	//a := make(map[string]map[string]string)
	//a := map[string]map[string]string{
	//	//	"zhangsan":{
	//	//		"key1": "value1",
	//	//},
	//	//}
	//	//modify(a)
	//	//
	//	//fmt.Println(a)
	//testMap4()
	testMap5()
}



func testMap4() {
	a := map[string]string {
		"1": "2",
		"2":"3",
		"4": "5",
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
	delete(a, "1")
	fmt.Println(len(a))
}

func testMap5() {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] ==nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)
}


