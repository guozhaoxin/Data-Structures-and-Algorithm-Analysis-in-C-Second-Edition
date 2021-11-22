package content

import (
	"chap4"
	"fmt"
)

type LinkedNode struct {
	value *chap4.Node
	next *LinkedNode
}

type Linked struct {
	head *LinkedNode
	tail *LinkedNode
}

func (self *Linked) IsEmpty() bool{
	return self.tail == nil
}

func (self *Linked)RightPush(value *chap4.Node) {
	node := &LinkedNode{value: value,next: nil}
	if self.head == nil{
		self.head = node
		self.tail = node
	}else {
		self.tail.next = node
		self.tail = node
	}
}

func (self *Linked)LeftPop() *chap4.Node{
	if self.IsEmpty(){
		return nil
	}
	res := self.head
	self.head = self.head.next
	if res == self.tail{
		self.tail = nil
	}
	return res.value
}

type Stack struct {
	top *LinkedNode
}

func (self *Stack)IsEmpty() bool{
	return self.top == nil
}

func (self *Stack) Top() *chap4.Node{
	if self.IsEmpty(){
		return nil
	}
	return self.top.value
}

func (self *Stack) Push(value *chap4.Node) {
	node := &LinkedNode{value: value,next: self.top}
	self.top = node
}
func (self *Stack) Pop() *chap4.Node{
	if self.IsEmpty(){
		return nil
	}
	node := self.top
	self.top = self.top.next
	return node.value
}


type BST struct {
	root *chap4.Node
}

func (self *BST) Reset(){
	self.root = nil
}

func (self *BST) Insert(value int) bool {
	if self.root == nil{
		root := &chap4.Node{Value: value,Left: nil,Right: nil}
		self.root = root
		return true
	}
	node := &chap4.Node{Value: value,Left: nil,Right: nil}
	root := self.root
	for root != nil{
		if root.Value == value{
			return false
		}
		if root.Value < value{
			if root.Right == nil{
				root.Right = node
				return true
			}
			root = root.Right
		}else {
			if root.Left == nil{
				root.Left = node
				return true
			}
			root = root.Left
		}
	}
	return false
}

func (self *BST) Find(value int) *chap4.Node {
	root := self.root
	for root != nil{
		if root.Value == value{
			return root
		}else if root.Value < value{
			root = root.Right
		}else {
			root = root.Left
		}
	}
	return nil
}

func (self *BST) Delete(value int) bool {
	var parent *chap4.Node = nil
	node := self.root
	for node != nil {
		if node.Value == value {
			break
		}else if node.Value < value {
			parent = node
			node = node.Right
		}else {
			parent = node
			node = node.Left
		}
	}
	if node == nil{ // cannot find target node
		return false
	}

	if node.Left == nil && node.Right == nil{ // leaf node
		if parent == nil{
			self.root = nil
		}else if parent.Left == node {
			parent.Left = nil
		}else {
			parent.Right = nil
		}
		return true
	}

	if node.Left == nil {
		if parent == nil{
			self.root = node.Right
		}else {
			if parent.Left == node {
				parent.Left = node.Right
			}else {
				parent.Right = node.Right
			}
		}
		return true
	}

	if node.Right == nil{
		if parent == nil{
			self.root = node.Left
		}else {
			if parent.Left == node {
				parent.Left = node.Left
			}else {
				parent.Right = node.Left
			}
		}
		return true
	}

	prev := node
	temp := node.Right
	for temp.Left != nil{
		prev = temp
		temp = temp.Left
	}
	if temp == node.Right {
		temp.Left = node.Left
		if parent == nil{
			self.root = temp
		}else {
			if parent.Left == node{
				parent.Left = temp
			}else {
				parent.Right = temp
			}
		}
		return true
	}

	prev.Left = temp.Right
	temp.Left = node.Left
	temp.Right = node.Right
	if parent == nil{
		self.root = temp
	}else if parent.Left == node{
		parent.Left = temp
	}else {
		parent.Right = temp
	}

	return true
}

func (self *BST) PreOrder(){
	if self.root == nil{
		fmt.Println()
		return
	}
	stack := Stack{}
	stack.Push(self.root)
	for !stack.IsEmpty(){
		node := stack.Top()
		if node != nil{
			stack.Pop()

			if node.Right != nil{
				stack.Push(node.Right)
			}
			if node.Left != nil{
				stack.Push(node.Left)
			}
			stack.Push(node)
			stack.Push(nil)
		}else {
			stack.Pop()
			node = stack.Pop()
			fmt.Print(node.Value,"\t")
		}
	}
	fmt.Println()
}

func (self *BST) InOrder(){
	if self.root == nil{
		fmt.Println()
		return
	}
	stack := Stack{}
	stack.Push(self.root)
	for !stack.IsEmpty(){
		node := stack.Top()
		if node != nil{
			stack.Pop()
			if node.Right != nil{
				stack.Push(node.Right)
			}
			stack.Push(node)
			stack.Push(nil)
			if node.Left != nil{
				stack.Push(node.Left)
			}
		}else {
			stack.Pop()
			node = stack.Pop()
			fmt.Print(node.Value,"\t")
		}
	}
	fmt.Println()
}

func (self *BST) PostOrder(){
	if self.root == nil{
		fmt.Println()
		return
	}
	stack := Stack{}
	stack.Push(self.root)
	for !stack.IsEmpty(){
		node := stack.Top()
		if node != nil{
			//stack.Pop()
			//stack.Push(node)
			stack.Push(nil)
			if node.Right != nil{
				stack.Push(node.Right)
			}
			if node.Left != nil{
				stack.Push(node.Left)
			}

		}else {
			stack.Pop()
			node = stack.Pop()
			fmt.Print(node.Value,"\t")
		}
	}
	fmt.Println()
}

func (self *BST) LevelOrder(){
	if self.root == nil{
		fmt.Println()
		return
	}
	curLink := Linked{}
	curLink.RightPush(self.root)
	for !curLink.IsEmpty(){
		nextLink := Linked{}
		for !curLink.IsEmpty(){
			node := curLink.LeftPop()
			fmt.Print(node.Value,"\t")
			if node.Left != nil{
				nextLink.RightPush(node.Left)
			}
			if node.Right != nil{
				nextLink.RightPush(node.Right)
			}
		}
		fmt.Println()
		curLink = nextLink
	}
	fmt.Println()
}