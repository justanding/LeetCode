package singleNumbers__test

import "testing"

func TestSingleNumbers(t *testing.T) {

    t.Log(singleNumbers([]int{4,1,4,7}))
}


/**
  * https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
 */
func singleNumbers(nums []int) []int {
    ret := 0;
    for _, v := range nums {
        ret ^=v
    }

    div :=1
    for (div & ret) == 0 {
        div <<= 1
    }

    a, b := ret, ret
    for _, v := range nums {
        if v & div == 0 {
            a ^= v
        } else {
            b ^= v
        }
    }
    return []int{a, b}
}
