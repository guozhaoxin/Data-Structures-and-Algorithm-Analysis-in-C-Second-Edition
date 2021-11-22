package content

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_Tree(t *testing.T){
	tree := BST{}
	r1 := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Second)
	r2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0; i < 100; i++{
		num := r1.Int() % 10
		value := r2.Int() % 10
		fmt.Println("-------------------------------")
		fmt.Print("前: ")
		tree.PreOrder()
		fmt.Print("中: ")
		tree.InOrder()
		fmt.Print("后: ")
		tree.PostOrder()
		fmt.Println("层: ")
		tree.LevelOrder()
		if value > 5 {
			fmt.Println("进入树: ",num)
			fmt.Println(tree.Insert(num))
		}else {
			fmt.Println("删除: ",num)
			fmt.Println(tree.Delete(num))
		}
		fmt.Print("前: ")
		tree.PreOrder()
		fmt.Print("中: ")
		tree.InOrder()
		fmt.Print("后: ")
		tree.PostOrder()
		fmt.Println("层: ")
		tree.LevelOrder()
		fmt.Println("-------------------------------")
	}
}