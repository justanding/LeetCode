package longestArithSeqLength_test

import (
	"testing"
)

func TestLongestArithSeqLength(t *testing.T) {
	t.Log(longestArithSeqLength([]int{1, 3, 6, 8, 9, 10, 12, 14}))
	t.Log(longestArithSeqLength([]int{12, 28, 13, 6, 34, 36, 53, 24, 29,
		2, 23, 0, 22, 25, 53, 34, 23, 50, 35,
		43, 53, 11, 48, 56, 44, 53, 31, 6, 31, 57, 46,
		6, 17, 42, 48, 28, 5, 24, 0, 14, 43, 12, 21, 6,
		30, 37, 16, 56, 19, 45, 51, 10, 22, 38, 39, 23, 8, 29, 60, 18}))
}

/*
 * https://leetcode-cn.com/problems/longest-arithmetic-sequence/
 * 动态规划 dp[j][seq] = dp[j-seq][seq]+1
 */
func longestArithSeqLength(A []int) int {
	l := len(A)
	if l <= 2 {
		return l
	}

	max := 0
	seqMap := map[int]map[int]int{}
	seqMap[0] = map[int]int{}
	for i := 1; i < l; i++ {
		mi := map[int]int{}
		for j := 0; j < i; j++ {
			seq := A[i] - A[j]
			m := seqMap[j]
			v, ok := m[seq]
			if ok {
				mi[seq] = v + 1
			} else {
				mi[seq] = 2
			}
			if max < mi[seq] {
				max = mi[seq]
			}
		}
		seqMap[i] = mi
	}
	return max
}
