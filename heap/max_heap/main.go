package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int
}

type Heap []*Item

//Len gets len of priority queue
func (pq Heap) Len() int {
	return len(pq)
}

//Less checks if i's priority is less than j
func (pq Heap) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Heap) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *Heap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	ksorted()
}

func kthsmallest() {
	// Some items and their priorities.
	items := []int{
		1, 23, 12, 9, 30, 2, 50,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := Heap{}
	for _, value := range items {

		heap.Push(&pq, &Item{
			value: value,
		})
		if pq.Len() > 3 {
			heap.Pop(&pq)
		}
	}
	for pq.Len() > 0 {
		fmt.Printf("Kth element %v\n", heap.Pop(&pq))
	}

}

func ksorted() {
	// Some items and their priorities.
	items := []int{
		6, 5, 3, 2, 8, 10, 9,
	}
	var sorted []int

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := Heap{}
	for _, value := range items {

		heap.Push(&pq, &Item{
			value: value,
		})
		if pq.Len() > 3 {
			sorted = append(sorted, heap.Pop(&pq).(*Item).value)
		}
	}
	for pq.Len() > 0 {
		sorted = append(sorted, heap.Pop(&pq).(*Item).value)
	}
	fmt.Printf("sorted %v\n", sorted)
}
