package maxSubArray_test

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Log(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

/*
 * https://leetcode-cn.com/problems/maximum-subarray/
 * 采用动态规划思路
 * 每个节点最大值的关系:f(n) = max(num[n]+f(n-1), num[n])
 */
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	max := nums[0]
	for i, maxCurrrent := 1, max; i < l; i++ {
		maxCurrrent += nums[i]
		if maxCurrrent < nums[i] {
			maxCurrrent = nums[i]
		}
		if max < maxCurrrent {
			max = maxCurrrent
		}
	}
	return max
}
