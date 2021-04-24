package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	value int
	diff  int
}

type PriorityQueue []*Item

//Len gets len of priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

//Less checks if i's priority is less than j
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].diff > pq[j].diff
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.diff)
}

func main() {
	kthsmallest()
}

func kthsmallest() {
	// Some items and their priorities.
	items := []int{
		10, 2, 14, 4, 7, 6,
	}
	x := 5
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := PriorityQueue{}
	for _, value := range items {
		heap.Push(&pq, &Item{
			value: value,
			diff:  int(math.Abs(float64(value - x))),
		})
		if pq.Len() > 3 {
			heap.Pop(&pq)
		}
	}
	for pq.Len() > 0 {
		fmt.Printf("Value %d \n", heap.Pop(&pq).(*Item).value)
	}
}
