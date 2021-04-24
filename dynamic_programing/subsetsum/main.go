package main

func subsetSum(sum int, arr []int) bool {
	dp := [1001][1001]bool{}
	for i := 0; i <= len(arr); i++ {
		for j := 0; j <= sum; i++ {
			if i == 0 {
				dp[i][j] = false
			}
			if j == 0 {
				dp[i][j] = true
			}

		}
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j < sum; j++ {
			if j > arr[i-1] {
				dp[i][j] = false
			} else {
				dp[i][j] = dp[i-1][j-arr[i]] || dp[i-1][j]
			}
		}
	}
	return dp[len(arr)][sum]
}

func recursion(sum int, arr []int, n int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 {
		return false
	}

	if sum > arr[n-1] {
		return recursion(sum, arr, n-1)
	}
	return recursion(sum-arr[n-1], arr, n-1) || recursion(sum, arr, n-1)
}

func main() {
	subsetSum(0, nil)
}
