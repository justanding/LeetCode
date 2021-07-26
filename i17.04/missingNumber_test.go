package i17_04

import "testing"

func TestMissingNumber(t *testing.T) {
    t.Log(missingNumber([]int{1,3,2,0}) == 4)
    t.Log(missingNumber([]int{1,3,4,0}) == 2)
    t.Log(missingNumber([]int{}) == 0)
}


func missingNumber(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    hash := make([]bool, len(nums)+1)

    for _, v := range nums {
        hash[v] = true
    }

    i := 0
    for _, f := range hash {
        if f != true {
            return i
        }
        i++
    }

    return len(nums)
}
