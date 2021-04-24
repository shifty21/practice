package main

import "reflect"

func subsets(nums []int) [][]int {
	var result *[][]int
	var output *[]int
	ret := solve(nums, output, result)
	return *ret
}
func solve(nums []int, output *[]int, result *[][]int) *[][]int {
	if len(nums) == 0 {
		Add(result, output)
		return result
	}
	op1 := output
	AddToArray(output, nums[0])
	op2 := output
	Add(result, op1)
	Add(result, op2)
	solve(nums, op1, result)
	solve(nums, op2, result)
	return result
}
func AddToArray(output *[]int, val int) {
	for x := range *output {
		if x == val {
			return
		}
	}
	*output = append(*output, val)
}
func Add(result *[][]int, add *[]int) {
	for _, v := range *result {
		// fmt.Printf("Checking equality for %v and add %v\n",v, add)
		if reflect.DeepEqual(v, *add) {
			return
		}
	}
	*result = append(*result, *add)
	return
}
