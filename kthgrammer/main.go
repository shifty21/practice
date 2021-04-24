package main

import "math"

func kthGrammar(N int, K int) int {
	if N == 1 && K == 1 {
		return 0
	}
	mid := int(math.Pow(2, float64(N-1)) / 2)
	if K <= mid {
		return kthGrammar(N-1, K) 
	}
	Out := kthGrammar(N-1, K-mid)
	Out = 1 - Out
	return Out
}
