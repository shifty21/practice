package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack []int
	var minStack []int
	return MinStack{stack: stack, minStack: minStack}
}

func (this *MinStack) Push(x int) {
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
		this.stack = append(this.stack, x)
		return
	}
	if this.minStack[len(this.minStack)-1] > x {
		this.minStack = append(this.minStack, x)
		this.stack = append(this.stack, x)
		return
	}
	this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) > 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return -1
}

func (this *MinStack) SortStack() []int {
	if len(this.stack) == 0 {
		return this.stack
	}
	return sort(this.stack)
}

//sort stack recursively
func sort(stack []int) []int {
	if len(stack) == 1 {
		return stack
	}
	top := stack[len(stack)-1]
	stack = sort(stack[:len(stack)-1])
	stack = insert(stack, top)
	return stack
}

//insert stack recursively - insert at top is its greater else take out top are repeat
func insert(stack []int, top int) []int {
	if stack[len(stack)-1] < top || len(stack) == 0 {
		stack = append(stack, top)
		return stack
	}
	val := stack[len(stack)-1]
	stack = insert(stack[:len(stack)-1], top)
	stack = append(stack, val)
	return stack
}

//DeleteMiddle recursively takes out top and repeate till we find the middle and append
func (this *MinStack) DeleteMiddle() []int {
	if len(this.stack) == 0 {
		return this.stack
	}
	k := len(this.stack) / 2
	stack := delete(this.stack, k)
	return stack
}

func delete(stack []int, k int) []int {
	if k == 0 {
		return stack[:len(stack)-1]
	}
	top := stack[len(stack)-1]
	stack = delete(stack[:len(stack)-1], k-1)
	stack = append(stack, top)
	return stack
}

//ReverseStack - takes out top and (insertEnd inserts at the end) that to reversed stack in
func (this *MinStack) ReverseStack() []int {
	if len(this.stack) == 0 {
		return this.stack
	}

	stack := reverse(this.stack)
	return stack
}

func reverse(stack []int) []int {
	if len(stack) == 1 {
		return stack
	}
	temp := stack[len(stack)-1]
	stack = reverse(stack[:len(stack)-1])
	stack = insertEnd(stack, temp)
	return stack
}

// insert at the end
func insertEnd(stack []int, val int) []int {
	if len(stack) == 0 {
		stack = append(stack, val)
		return stack
	}
	top := stack[len(stack)-1]
	stack = insertEnd(stack[:len(stack)-1], val)
	stack = append(stack, top)
	return stack
}

func main() {
	obj := Constructor()
	obj.Push(-3)
	obj.Push(0)
	obj.Push(6)
	obj.Push(2)
	obj.Push(1)
	stack := obj.SortStack()
	fmt.Printf("Sorted stack %v\n", stack)

	// stack = obj.DeleteMiddle()
	// fmt.Printf("Sorted stack after deleting %v", stack)
	stack = obj.ReverseStack()
	fmt.Printf("Rever stack after deleting %v\n", stack)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
