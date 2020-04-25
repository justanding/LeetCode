package permute__test

import (
    "testing"
)

func TestPermute(t *testing.T) {
    s := []int{1, 2, 3}
    t.Log(permute(s))
}

/**
 * https://leetcode-cn.com/problems/permutations/submissions/
 *
 */
func permute(nums []int) [][]int {
    result := [][]int{}
    if len(nums) == 0 {
        return [][]int{}
    }

    for i, num := range nums {
        newNums := make([]int, len(nums[:i]))
        copy(newNums,nums[:i])
        newNums = append(newNums, nums[i+1:]...)
        ret := permute(newNums)
        if len(ret) == 0 {
            result = append(result,[]int{num})
            continue
        }

        for _, r := range ret {
            newR := make([]int, len(r))
            copy(newR, r)
            n := append(newR, num)
            result = append(result, n)
        }
    }
    return result
}

