package main

import "fmt"

func main() {
	var i int = 0
	defer fmt.Println(i)
	defer fmt.Println("second")

	i = 10
	fmt.Println(i)
}

func read() {
	file := open(filename)
	defer  file.Close()
	// 文件操作
}

// 资源释放
func read1() {
	mc.Lock()
	defer mc.Unlock()
	// 其他操作
}

func read2() {
	conn, err := openConnect()
	defer func() {
		if(conn != nil) {
			conn.close()
		}
	}()
}
