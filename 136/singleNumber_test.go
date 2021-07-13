package _36

import "testing"

func TestMaxSubArray(t *testing.T) {
    t.Log(singleNumber([]int{2,2,1,1,0}))
}


func singleNumber(nums []int) int {
    ret := nums[0]

    for i:=1;i<len(nums);i++ {
        ret ^= nums[i]
    }


    return ret
}
