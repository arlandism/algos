package main

import (
	"fmt"
)

func longestIncreasingSubsequence(x []int) []int {
	cachedSeqs := make(map[int][]int)
	for i := 0; i < len(x); i++ {
		cachedSeqs[x[i]] = []int{x[i]}
		for j := 0; j < i; j++ {
			// if x[j] is less than x[i], then check to see if the longest sequence for x[j] + x[i] would
			// be a new longest sequence for x[i]
			if x[j] < x[i] {
				longestSubForJ := cachedSeqs[x[j]]
				longestSubForI := cachedSeqs[x[i]]
				if len(longestSubForJ)+1 > len(longestSubForI) {
					// need to copy slice here to avoid mutating previous seqs
					tmp := make([]int, len(longestSubForJ))
					copy(tmp, longestSubForJ)
					cachedSeqs[x[i]] = append(tmp, x[i])
				}
			}
		}
	}
	max := 0
	var maxSeq []int
	// NOTE: there's more than one possible match here, just grab whichever
	for _, val := range cachedSeqs {
		if len(val) > max {
			maxSeq = val
			max = len(val)
		}
	}
	return maxSeq
}

func main() {
	testCases := [][]int{
		[]int{2, 4, 3, 5, 1, 7, 6, 9, 8}, // multiple options, all length 5
		[]int{9, 3, 20, 48, 29, 2, 1},    // multiple options, all length 3
		[]int{5, 4, 3},                   // multiple options, all length 1
		[]int{1, 2, 3},                   // 1, 2, 3
		[]int{19, 27, 1, 2, 3, 4},        // 1, 2, 3, 4
	}

	for _, x := range testCases {
		fmt.Printf("A longest subsequence in %v is %v\n", x, longestIncreasingSubsequence(x))
	}
}
