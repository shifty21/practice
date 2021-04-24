package main

import "fmt"

// func longestPalindrome(s string) string {
// 	if len(s) <= 1 {
// 		return s
// 	}
// 	var max int
// 	var result string
// 	store := make(map[string]bool)
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s)+1; j++ {
// 			temp := s[i:j]
// 			if status, ok := store[temp]; ok {
// 				if status && len(temp) >= max {
// 					result = temp
// 					max = len(temp)
// 				}
// 			} else {
// 				// store doesnt contain this string
// 				status := check(temp)
// 				if status && len(temp) > max {
// 					result = temp
// 					max = len(temp)

// 				}
// 				store[temp] = status
// 			}
// 		}
// 	}

// 	return result
// }

// func check(s string) bool {
// 	var result string

// 	for i := len(s) - 1; i >= 0; i-- {
// 		result = result + string(s[i])
// 	}
// 	fmt.Printf("result=%v,s=%v\n", result, s)
// 	return result == s
// }

func longestPalindrome(s string) string {
	lo, maxLen := 0, 0
	if len(s) < 2 {
		return s
	}

	for i := range s {
		extendPal(s, i, i, &lo, &maxLen)
		extendPal(s, i, i+1, &lo, &maxLen)
	}

	return s[lo : lo+maxLen]
}

func extendPal(s string, left, right int, lo, maxLen *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	if *maxLen < right-left-1 {
		*lo = left + 1
		*maxLen = right - left - 1
	}
}

func main() {
	// "babad"
	result := longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth")
	fmt.Printf("Result=%v\n", result)
}
