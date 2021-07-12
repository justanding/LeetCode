package findTargetSumWays_test

import "testing"

func TestFindTargetSumWays(t *testing.T) {

	t.Log(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	t.Log(findTargetSumWays([]int{1, 1, 1, 1, 1}, 4))
}

/*
 * https://leetcode-cn.com/problems/target-sum/
 * 递归
 */
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	ways := 0
	if len(nums) == 1 {
		if S == nums[0] {
			ways++
		}
		if S == -nums[0] {
			ways++
		}
		return ways
	}

	ways += findTargetSumWays(nums[1:], S-nums[0])
	ways += findTargetSumWays(nums[1:], S+nums[0])
	return ways
}
