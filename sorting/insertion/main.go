package main

import "log"

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for arr[j] > key && j >= 0 {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key

	}
	return arr
}
func main() {
	result := sort([]int{1, 7, 2, 9, 3})
	log.Printf("Sorted array %v", result)
}
