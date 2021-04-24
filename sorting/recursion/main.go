package main

import "log"

func sortArray(nums []int) []int {
	//Base condition
	if len(nums) <= 1 {
		return nums
	}
	last := nums[len(nums)-1]
	nums = sortArray(nums[:len(nums)-1])
	nums = insert(nums, last)
	return nums
}

func insert(nums []int, last int) []int {
	//base condition
	if len(nums) == 0 || nums[len(nums)-1] <= last {
		nums = append(nums, last)
		return nums
	}
	val := nums[len(nums)-1]
	nums = insert(nums[:len(nums)-1], last)
	nums = append(nums, val)
	return nums
}

func main() {

	result := sortArray([]int{5, 2, 3, 1})
	log.Printf("result %v", result)
}
