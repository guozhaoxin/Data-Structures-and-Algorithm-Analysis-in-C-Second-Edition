package exercise

import "fmt"

/*
this problem asks you to create 3 stacks within a list.
problem 3.21 asks you to create 2 stacks within a list.
for 3.21, we can use 2 indexs,such as i1 and i2; at anytime, i1 points to top of first stack, while i2 points to top of the second stack.
image the list has N elements, so at first,i1 is -1, while i2 is N, shown as below;
i1													i2
	_	_	_	_	_	_	_	_	_	_	_	_
when one elements is pused to any stacks, we check whether i1 == i2 - 1.
if so,shown as below:
				i1	i2
	_	_	_	_	_	_	_	_	_	_	_	_
	no space for any element, so the stack is full;
else:
	if ele is pushed to first stack,we only need to add 1 to i1 and set list[i1] = ele::
		i1					i2
	_	_	_	_	_	_	_	_	_	_	_	_
		    i1				i2
	_	_	ele	_	_	_	_	_	_	_	_	_
the same as above for stack 2 but minus 1 from i2 and set the value.

when pop one ele from stack 1,we need to check if i1 is -1,in which case stack 1 is empty.
if not empty,pop list[i1] and minus 1 from i1:
			i1				i2
	_	_	ele	_	_	_	_	_	_	_	_	_
		i1					i2
	_	_	_	_	_	_	_	_	_	_	_	_
the same as above for stack 2 but add 1 from i2.

but 3 stacks cannot be done by this method.
we can still use 2 index i1 and i3 to sign the top of the first and third stack, while for the second stack, we must use 2 indexs i2s and i2e to sign
the start and end of the stack,as shown below:

		i1			i2s        i2e		i3
	_	_	_	_	_	_	_	_	_	_	_	_
but attention, i1 and i3 points to the top ele of stack respectively, while i2e not,
i2e points to the position of the top ele + 1, while i2s points to the bottom of the stack 2.
for example:
		i1			i2s        i2e		i3
	1	2	3	4	5	6	7	8	9	10	11	12
in this case,
stack 1 is [1,2],
stack 2 is [5,6,7],
stack 3 is [12,11,10]
stack 1 is empty if i1 is 0, stack 3 is empty if i3 is N, what about stack 2?
we claim stack 2 is empty when i2s == i2e.
at first, the array is as shown:
i1													i2
	_	_	_	_	_	_	_	_	_	_	_	_
i2s
i2e
we can divide all situations as below:
case 1:
	i1 == i2s
		i1
	_	_	_	_	_	_	_	_	_	_	_	_
		i2s
		i2e
	in this case, stack 2 is empty(becase i1 i3 points to valid top stack ele,while any ele cannot be in 2 different stacks at the same time,so stack 2 must be empty),
	as mentioned above, i2e == i2s when stack 2 is empty.
	if i3 == i2e + 1, as shown below:
				i1	i3
			_	_	_	_	_	_	_	_	_	_	_	_
				i2s
				i2e
		no any space to store any new eles to any one stack,so return;
	else there must be at least one position to store the new ele.
				i1				i3
			_	_	_	_	_	_	_	_	_	_	_	_
				i2s
				i2e
		if the target stack is stack 1, we directly add 1 to i1 and set list[i1] to new ele; meanwhile, add 1 to i2s and i2e,so the result will be below:
					i1			i3
			_	_	ele	_	_	_	_	_	_	_	_	_
					i2s
					i2e
		if the target stack is stack 2, as the stack 2 is empty, so we must add 1 to i2s and set list[i2s] to new ele, but add 2 to i2e(remember i2e points to next
			position of the top ele of stack 2), so the result will be below:
				i1				i3
			_	_	ele	_	_	_	_	_	_	_	_	_
					i2s
						i2e
		if the target stack is stack 1, we directly minus 1 to i3 and set value, so the result will be below:
				i1			i3
			_	_	_	_	ele	_	_	_	_	_	_	_
				i2s
				i2e
case 2:
	in this case, stack 2 is not empty, so i2s != i2e.
				i1					?	?	?	?	?	?
			_	_	_	_	_	_	_	_	_	_	_	_
					?	?
							?	?   ?
	but this case is very complex, so we must continue to discuss different case.
		case i3 == i2e, in this case there is no free space between stack 2 and stack 3:
					i1					i3
			_	_	_	_	_	_	_	_	_	_	_	_
						?	?	?	?
										i2e
			if i1 == i2s - 1:
						i1					i3
				_	_	_	_	_	_	_	_	_	_	_	_
							i2s
											i2e
			we can find no any space to store the new ele,so return.
			else there is at least one position to store the ele as below:
						i1					i3
					_	_	_	_	_	_	_	_	_	_	_	_
								i2s
											i2e
				if the target stack is stack1,we just add 1 to s1 and set list[s1] = value and return as below:
						i1					i3
					_	_	_	_	_	_	_	_	_	_	_	_
								i2s
											i2e
							i1				i3
					_	_	ele	_	_	_	_	_	_	_	_	_
								i2s
											i2e

				if the target stack is stack2, as between stack2 and stack3 there is no space, so we must move the stack 2 left 1 position wholly as below:
						i1					i3
					_	_	_	x	y	z	_	_	_	_	_	_
								i2s
											i2e
						i1					i3
					_	_	x	y	z	ele	_	_	_	_	_	_
							i2s
											i2e
				after operation,i2s = i2s - 1, while i2e not change.

				if the target stack stack3, we must still move the stack to left one position to leave a free position for the new ele, as shown below:
						i1					i3
					_	_	_	x	y	z	_	_	_	_	_	_
								i2s
											i2e
						i1				i3
					_	_	x	y	z	ele	_	_	_	_	_	_
							i2s
										i2e
				after operation, i2s = i2s - 1, i2e = i2e - 1.

		case i1 == i2s - 1, so there is no free space between stack1 and stack2 while there is at least one pos between stack2 and stack3.
					i1					?	?	?	?	?
			_	_	_	_	_	_	_	_	_	_	_	_
						i2s
							?	?   ?
			if the target stack is stack1,we must move stack2 to right 1 position,as shown below:
						i1						i3
				_	_	_	x	y	_	_	_	_	_	_	_
							i2s
									i2e
						    i1					i3
				_	_	_	ele	x	y	_	_	_	_	_	_
								i2s
										i2e

			if the target stack is stack2,just set list[i2e] = ele and add 1 to i2e:
						i1						i3
				_	_	_	_	_	_	_	_	_	_	_	_
							i2s
									i2e
						i1						i3
				_	_	_	_	_	_  ele	_	_	_	_	_	_
							i2s
										i2e

			if the target stack is stack3, juset minus 1 from i3, and set list[i3] = ele:
						i1						i3
				_	_	_	_	_	_	_	_	_	_	_	_
							i2s
									i2e
						i1					i3
				_	_	_	_	_	_  _	ele	_	_	_	_	_
							i2s
									i2e

		not any above case, then the array must like below:
				i1							i3
			_	_	_	_	_	_	_	_	_	_	_	_
						i2s
								i2e
		i1 < i2s - 1 and i2e < i3, so each of the 2 gaps has at least one free space, we can push 1 ele to each of the 3 stacks.
*/

const SIZE = 10 // length of the array

func moveArray(array []int,left,right,index int) {
	if index == 1{ // right move
		for right > left{
			array[right] = array[right - 1]
			right--
		}
	}else { // left move
		for left < right {
			array[left - 1] = array[left]
			left++
		}
	}
}


type MultStacks struct {
	array []int
	i1  int
	i2s int
	i2e int
	i3  int
}

func (self *MultStacks) Init(){
	self.array = make([]int, SIZE)
	for i := range self.array {
		self.array[i] = 0
	}
	self.i1 = -1
	self.i2s = -1
	self.i2e = -1
	self.i3 = SIZE
}

func (self *MultStacks) ShowStacks() {
	for i := 0; i < len(self.array); i++ {
		fmt.Print(self.array[i],"\t\t")
	}
	fmt.Println()
	hehe := make([]string,len(self.array))
	for i := range hehe {
		hehe[i] = ""
	}
	for i := 0; i <= self.i1; i++{
		hehe[i] = "s1"
	}
	if self.i2s != self.i2e {
		for i := self.i2s; i < self.i2e; i++{
			if hehe[i] != ""{
				hehe[i] += ",s2"
			}else {
				hehe[i] =  "s2"
			}
		}
	}
	for i := self.i3; i < len(self.array); i++{
		if hehe[i] != ""{
			hehe[i] += ",s3"
		}else {
			hehe[i] = "s3"
		}
	}

	for _,item := range hehe{
		fmt.Print(item,"\t\t")
	}
	fmt.Println()
}

func (self *MultStacks) Push(index int,value int) (res bool){
	if index != 1 && index != 2 && index != 3{
		fmt.Println("wrong index ",index)
		res = false
		return
	}
	res = false
	fmt.Println("-------------------------")
	fmt.Println("element ", value, " is pushed to stack ",index)
	fmt.Println("stacks currently shown as below:")
	self.ShowStacks()
	defer func() {
		fmt.Println("after push stacks are shown as below:")
		self.ShowStacks()
		if res{
			fmt.Println("success")
		}else {
			fmt.Println("stacks full, fail.")
		}
	}()
	if self.i1 == self.i2s { // stack 2 is empty
		if self.i3 == self.i2e + 1 { // all 3 stacks are full
			res = false
			return
		}
		if index == 1 {
			fmt.Println("push to stack 1, but move stack 2 right.")
			self.i1++
			self.array[self.i1] = value
			self.i2s++
			self.i2e++
			res = true
			return
		}
		if index == 2{
			fmt.Println("push to stack 2,i2s moved right 1 pos,while i2e moved right 2 pos.")
			self.i2s++
			self.i2e += 2
			self.array[self.i2s] = value
			res = true
			return
		}
		if index == 3{
			fmt.Println("pushed to stack 3 normally.")
			self.i3--
			self.array[self.i3] = value
			res = true
			return
		}
		return
	}

	if self.i3 == self.i2e {
		if self.i1 == self.i2s - 1 {
			res = false
			return
		}
		if index == 1 {
			fmt.Println("push to stack 1 normally.")
			self.i1++
			self.array[self.i1] = value
			res = true
			return
		}

		if index == 2 {
			fmt.Println("move stack 2 left.")
			moveArray(self.array,self.i2s,self.i2e,-1)
			self.i2s--
			self.array[self.i2e - 1] = value
			res = true
			return
		}

		if index == 3{
			fmt.Println("move stack 2 left.")
			moveArray(self.array,self.i2s,self.i2e,-1)
			self.i2s--
			self.i2e--
			self.i3--
			self.array[self.i3] = value
			res = true
			return
		}

		return
	}

	if self.i1 == self.i2s - 1 {
		if index == 1 {
			fmt.Println("move stack 2 right")
			moveArray(self.array,self.i2s,self.i2e,1)
			self.i2s++
			self.i2e++
			self.i1++
			self.array[self.i1] = value
			res = true
			return
		}

		if index == 2{
			fmt.Println("pusth to stack 2 normally.")
			self.array[self.i2e] = value
			self.i2e++
			res = true
			return
		}

		if index == 3{
			fmt.Println("push to stack 2 normally")
			self.i3--
			self.array[self.i3] = value
			res = true
			return
		}

		return
	}

	if index == 1{
		fmt.Println("push to stack 1 normally.")
		self.i1++
		self.array[self.i1] = value
	}else if index == 2{
		fmt.Println("push to stack 2 normally.")
		self.array[self.i2e] = value
		self.i2e++
	}else {
		fmt.Println("push to stack 3 normally.")
		self.i3--
		self.array[self.i3] = value
	}

	res = true
	return
}

func (self *MultStacks) Pop(index int) (value int,res bool) {
	if index != 1 && index != 2 && index != 3 {
		fmt.Println("wrong index ",index)
		res = false
		return
	}
	fmt.Println("-------------------------")
	fmt.Println("stack ",index, " pop")
	fmt.Println("stacks are shown below:")
	self.ShowStacks()
	res = false
	defer func() {
		if res {
			fmt.Println("popd value is ",value)
			self.ShowStacks()
		}else {
			fmt.Println("fail to pop")
		}
	}()
	if index == 1{
		if self.i1 == -1 {
			res = false
			return
		}
		value = self.array[self.i1]
		self.array[self.i1] = 0
		self.i1--
		if self.i2s == self.i1 + 1{
			self.i2s = self.i1
			self.i2e = self.i1
		}
		res = true
	}

	if index == 2{
		if self.i2s == self.i2e {
			res = false
			return
		}
		self.i2e--
		value = self.array[self.i2e]
		self.array[self.i2e] = 0
		res = true
		if self.i2s == self.i2e { // stack2 is empty
			fmt.Println("reset i2")
			self.i2s = self.i1
			self.i2e = self.i1
		}
		return
	}

	if index == 3{
		if self.i3 == SIZE {
			res = false
			return
		}
		value = self.array[self.i3]
		self.array[self.i3] = 0
		self.i3++
		res = true
		return
	}
	return
}
