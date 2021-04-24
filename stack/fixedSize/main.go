package main

type CustomStack struct {
	stack []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, 0, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) == len(this.stack) {

	}
	this.stack = append(this.stack, x)
}

func (this *CustomStack) Pop() int {
	return -1
}

func (this *CustomStack) Increment(k int, val int) {

}
