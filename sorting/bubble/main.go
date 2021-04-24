package main

import "log"

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	result := sort([]int{1, 7, 2, 9, 35, 22, 3})
	log.Printf("Sorted array %v", result)
}
