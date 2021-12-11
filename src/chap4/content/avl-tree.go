package content

import (
	"fmt"
)

func absInt(value int) int {
	if value < 0{
		return -value
	}
	return value
}

func max(v1,v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func getLeftHeight(node *TreeNode) int {
	if node == nil{
		return -1
	}
	if node.left == nil{
		return -1
	}
	return node.left.height
}

func getRightHeight(node *TreeNode) int {
	if node == nil{
		return -1
	}
	if node.right == nil{
		return -1
	}
	return node.right.height
}

func setNodeHeight(node *TreeNode){
	if node == nil{
		return
	}
	leftHeight := -1
	if node.left != nil{
		leftHeight = node.left.height
	}
	rightHeight := -1
	if node.right != nil{
		rightHeight = node.right.height
	}
	node.height = max(leftHeight,rightHeight) + 1
}

type TreeNode struct{
	value int
	left *TreeNode
	right *TreeNode
	height int
}

type AVL struct{
	root *TreeNode
}

func (self *AVL) Insert(value int) bool{
	node := self.insert(self.root,value)
	if node == nil{
		return false
	}
	self.root = node
	return true
}

func (self *AVL) insert(root *TreeNode,value int) *TreeNode{
	if root == nil {
		root = &TreeNode{value:value,left:nil,right:nil,height:0}
		return root
	}
	var node *TreeNode
	if root.value < value {
		node = self.insert(root.right,value)
		if node == nil{
			return nil
		}
		root.right = node
		leftHeight := -1
		if root.left != nil{
			leftHeight = root.left.height
		}
		rightHeight := node.height
		//if root.right != nil{
		//	rightHeight = root.right.height
		//}
		if absInt(leftHeight - rightHeight) > 1{
			if value < root.right.value{ // right-left
				return RotateRightLeft(root)
			}else { // right-right
				return RotateRightRight(root)
			}
		}else{
			root.height = max(leftHeight,rightHeight) + 1
		}
		return root
	}else if root.value > value {
		node = self.insert(root.left,value)
		if node == nil{
			return nil
		}
		root.left = node
		leftHeight := node.height
		//if root.left != nil{
		//	leftHeight = root.left.height
		//}
		rightHeight := -1
		if root.right != nil{
			rightHeight = root.right.height
		}
		if absInt(leftHeight - rightHeight) > 1{
			if value > root.left.value{ // left-right
				return RotateLeftRight(root)
			}else { // left-left
				return RotateLeftLeft(root)
			}
		}else{
			root.height = max(leftHeight,rightHeight) + 1
		}
		return root
	}else {
		return nil
	}
}

func (self *AVL) InsertWithStack(value int) bool {
	//if self.root == nil{
	//	self.root = &TreeNode{value: value,height: 0,left: nil,right: nil}
	//	return true
	//}
	stack := make([]*TreeNode,0)
	var node *TreeNode
	root := self.root
	for {
		if root == nil{
			node = &TreeNode{value: value,left: nil,right: nil,height: 0}
			break
		}
		if root.value == value{
			return false
		}
		if root.value < value {
			stack = append(stack,root)
			root = root.right
		}else {
			stack = append(stack,root)
			root = root.left
		}
	}
	for len(stack) != 0{
		top := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if top.value < value{
			top.right = node
			leftHeight := getLeftHeight(top)
			rightHeight := getRightHeight(top)
			if absInt(leftHeight - rightHeight) > 1{
				if value < node.value{ // 右左
					node = RotateRightLeft(top)
				}else { // 右右
					node = RotateRightRight(top)
				}
			}else {
				setNodeHeight(top)
				node = top
			}
		}else {
			top.left = node
			leftHeight := getLeftHeight(top)
			rightHeight := getRightHeight(top)
			if absInt(leftHeight - rightHeight) > 1{
				if value > node.value{ // 左右
					node = RotateLeftRight(top)
				}else { // 左左
					node = RotateLeftLeft(top)
				}
			}else {
				setNodeHeight(top)
				node = top
			}
		}
	}
	self.root = node
	return true
}


func (self *AVL)Order(){
	order(self.root)
	fmt.Println()
}

func (self *AVL)PreOrder(){
	preOrder(self.root)
	fmt.Println()
}

func (self *AVL)PostOrder(){
	postOrder(self.root)
	fmt.Println()
}

func order(root *TreeNode){
	if root == nil{
		return
	}
	order(root.left)
	fmt.Print(root.value)
	fmt.Print("\t")
	order(root.right)
}

func preOrder(root *TreeNode){
	if root == nil{
		return
	}
	fmt.Print(root.value,"\t")
	preOrder(root.left)
	preOrder(root.right)
}

func postOrder(root *TreeNode){
	if root == nil{
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.value,"\t")
}

func RotateLeftLeft(root *TreeNode) *TreeNode{
	fmt.Println("left left")
	left := root.left
	leftRight := left.right
	left.right = root
	root.left = leftRight
	rootLH := -1
	rootRH := -1
	if root.left != nil{
		rootLH = root.left.height
	}
	if root.right != nil{
		rootRH = root.right.height
	}
	root.height = max(rootLH,rootRH) + 1
	leftLeftH := -1
	if left.left != nil{
		leftLeftH = left.right.height
	}
	left.height = max(leftLeftH,left.right.height) + 1
	return left
}

func RotateRightRight(root *TreeNode) *TreeNode{
	fmt.Println("right right")
	right := root.right
	rightLeft := right.left
	right.left = root
	root.right = rightLeft
	rootLH := -1
	rootRH := -1
	if root.left != nil{
		rootLH = root.left.height
	}
	if root.right != nil{
		rootRH = root.right.height
	}
	root.height = max(rootLH,rootRH) + 1
	rightRightH := -1
	if right.right != nil{
		rightRightH = right.right.height
	}
	right.height = max(right.left.height,rightRightH) + 1
	return right
}

func RotateLeftRight(root *TreeNode) *TreeNode{
	fmt.Println("left right")
	root.left = RotateRightRight(root.left)
	return RotateLeftLeft(root)
}

func RotateRightLeft(root *TreeNode) *TreeNode{
	fmt.Println("right left")
	root.right = RotateLeftLeft(root.right)
	return RotateRightRight(root)
}
