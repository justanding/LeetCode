package _98

import "testing"

func TestRob(t *testing.T) {
    t.Log(rob([]int{3,2,1,10,1,3}))
    t.Log(rob([]int{3,2}))
    t.Log(rob([]int{3}))
    t.Log(rob([]int{226,174,214,16,218,48,153,131,128,17,157,142,88,43,37,157,43,221,191,68,206,23,225,82,54,118,111,46,80,49,245,63,25,194,72,80,143,55,209,18,55,122,65,66,177,101,63,201,172,130,103,225,142,46,86,185,62,138,212,192,125,77,223,188,99,228,90,25,193,211,84,239,119,234,85,83,123,120,131,203,219,10,82,35,120,180,249,106,37,169,225,54,103,55,166,124}))
}

func rob(nums []int) int{
    lenght := len(nums)
    if lenght == 0 {
        return 0
    } else if lenght ==1 {
        return nums[0]
    }

    for i:=0; i<lenght; i++ {
        nums[i] = nums[i]+max(nums, i-2, i-3)
    }
    if nums[lenght-1] > nums[lenght-2] {
        return nums[lenght-1]
    } else {
        return nums[lenght-2]
    }
}

func max(maxSlice []int, max, min int) int {
    if min>=0 {
        if maxSlice[max] > maxSlice[min] {
            return maxSlice[max]
        } else {
            return maxSlice[min]
        }
    } else if max>=0 {
        return maxSlice[max]
    }

    return 0
}


func robRecur(nums []int) int {

    r1 := recurRob(nums, len(nums)-1)
    r2 := recurRob(nums, len(nums)-2)
    if r1>r2 {
        return r1
    }
    return r2
}

func recurRob(nums []int, n int) int {
    if n<0 {
        return 0
    }

    r1 := recurRob(nums, n-2)
    r2 := recurRob(nums, n-3)

    if r1>r2 {
        return r1+nums[n]
    }
    return r2 + nums[n]
}