package main

import "log"

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
func main() {
	result := sort([]int{1, 7, 2, 9, 3})
	log.Printf("Sorted array %v", result)
}
