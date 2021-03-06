package main

type MyQueue struct {
	stack1 MyStack
	stack2 MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: StackConstructor(),
		stack2: StackConstructor(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return -1
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return false
}

type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func StackConstructor() MyStack {
	stack := MyStack{
		stack: []int{},
	}
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	temp := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return temp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}
