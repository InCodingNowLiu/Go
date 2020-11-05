package balance

import (
	"errors"
	"math/rand"
)

// init 的逻辑, 是应用的时候生效么??????
func init() {
	RegisterBalancer("random", &RandomBalance{})
}


type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string)(inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}