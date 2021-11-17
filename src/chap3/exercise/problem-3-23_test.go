package exercise

import (
	"math/rand"
	"testing"
	"time"
)

func TestMultStacks(t *testing.T) {
	stack := &MultStacks{}
	stack.Init()
	r1 := rand.New(rand.NewSource(time.Now().Unix()))
	r2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 1; i < 100 ;i ++ {
		index := r1.Intn(3) + 1
		value := r2.Int() % 10
		if value > 2 {
			stack.Push(index,i)
		}else {
			stack.Pop(index)
		}
	}
}