package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	priority int
	value    string
	index    int
}

type PriorityQueue []*Item

//Len gets len of priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

//Less checks if i's priority is less than j
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	l := len(*pq)
	item := x.(*Item)
	item.index = l
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	kthsmallest()
}

func kthsmallest() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 7, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	// pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	// for pq.Len() > 0 {
	// 	item := heap.Pop(&pq).(*Item)
	// 	fmt.Printf("%.2d:%s ", item.priority, item.value)
	// }
}
