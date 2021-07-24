package offer03

import "testing"

func TestFindRepeatNumber(t *testing.T) {
    t.Log(findRepeatNumber([]int{}))
    t.Log(findRepeatNumber([]int{1,1}))
    t.Log(findRepeatNumber([]int{1,3,3,1}))
    t.Log(findRepeatNumber([]int{1,3,2,1,5,4}))
}


func findRepeatNumberMap(nums []int) int {
    hash := map[int]bool{}

    for _, v := range nums {
        if hash[v] {
            return v
        }
        hash[v] = true
    }

    return -1
}

func findRepeatNumber(nums []int) int {
    sl := make([]int, len(nums))

    for _, v := range nums {
        if sl[v] != 0 {
            return v
        }
        sl[v] = 1
    }
    return -1
}


