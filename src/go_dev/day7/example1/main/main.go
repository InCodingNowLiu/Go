package main

import (
	"awesomeProject/src/go_dev/day7/example1/balance"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i<len(insts); i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}