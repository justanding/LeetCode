package _5

import "testing"

func Test_searchInsert(t *testing.T) {
    t.Log(searchInsert([]int{1,2,3}, 3))
    t.Log(searchInsert([]int{1,2,3,4,5}, 2))
    t.Log(searchInsert([]int{1,2,3,4,5}, 0))
    t.Log(searchInsert([]int{1,2,3,4,5}, 9))
    t.Log(searchInsert([]int{}, 0))

}


func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }

    for k, v := range nums {
        if v >= target {
            return k
        }
    }
    return len(nums)
}