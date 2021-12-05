package content

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAVL_Insert(t *testing.T) {
	tree := &AVL{}
	r1 := rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0; i < 100; i++{
		fmt.Println("-------")
		value := r1.Int() % 100
		fmt.Println(value)
		fmt.Println(tree.Insert(value))
		tree.PreOrder()
		tree.Order()
		tree.PostOrder()
		fmt.Println("-------")
	}
}

func TestAVL_Insert2(t *testing.T) {
	tree := &AVL{}

	for i:=0; i < 10; i++{
		fmt.Println("-------")
		fmt.Println(i)
		tree.Insert(i)
		tree.PreOrder()
		tree.Order()
		tree.PostOrder()
		fmt.Println("-------")
	}
}