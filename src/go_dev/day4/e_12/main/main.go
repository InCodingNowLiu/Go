package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex

func testMap() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[8] = 10
	a[3] = 10
	for i:=0;i<2;i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}

func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10
	var count int32
	for i:=0;i<2;i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[8] = rand.Intn(100)
			rwLock.Unlock()
		}(a)
	}
	for i :=0;i<100;i++ {
		go func(b map[int]int) {
			rwLock.RLock()
			fmt.Println(a)
			rwLock.RUnlock()
			atomic.AddInt32(&count, 1)
		}(a)
	}
	time.Sleep(time.Second *3)
	fmt.Println(atomic.LoadInt32(&count))
}

// go get 安装第三方包
func main() {
	//testMap()
	testRWLock()
}
