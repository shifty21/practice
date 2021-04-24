package main

func knapsack(W int, wt []int, val []int, n int) int {
	momiz := [1002][1002]int{}
	return internal(W, wt, val, n, momiz)
}

func internal(W int, wt []int, val []int, n int, momiz [1002][1002]int) int {
	if n == 0 || W == 0 {
		return 0
	}

	if wt[n-1] <= W {
		inc, noninc := val[n-1]+knapsack(W-wt[n-1], wt, val, n-1), knapsack(W, wt, val, n-1)
		if inc > noninc {
			return inc
		}
		return noninc
	}
	return knapsack(W, wt, val, n-1)
}
