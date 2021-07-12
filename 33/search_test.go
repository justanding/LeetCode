package search_test

import (
    "testing"
)

func TestSearch(t *testing.T) {
    for _,v := range []int{0,1,2,3,4,5,6} {
       t.Log(search([]int{5,1,2,3,4}, v))
    }

    for _,v := range []int{-1,0,1,2,3,4,5,6, 7, 8,9} {
        t.Log(search([]int{4,5,6,7,8,0,1,2}, v))
    }
}

func search(nums []int, target int) int {
    l := len(nums)
    if l < 1 {
        return -1
    }
    return searchRev(nums, target, 0, l-1)
}

func searchRev(nums []int,  target, start, end int) int {
    if nums[start] == target {
        return start
    }

    if nums[end] == target {
        return end
    }

    if end - start <= 1 {
        return -1
    }

    mid := (start + end)/2
    if nums[end] < nums[start]{
        if nums[end] > target {
            if (nums[mid]<nums[end] && nums[mid]<target) ||
                (nums[mid]>nums[end] && nums[end]>target) {
                return searchRev(nums, target, mid, end)
            } else {
                return searchRev(nums, target, start, mid)
            }
        } else {
            if (nums[mid] > nums[start] &&  nums[mid] > target) ||
                (nums[mid] < nums[start] && nums[start]< target) {
                return searchRev(nums, target, start, mid)
            } else {
                return searchRev(nums, target, mid, end)
            }
        }
    } else {
        if nums[mid] < target {
            return searchRev(nums, target, mid, end)
        } else {
            return searchRev(nums, target, start, mid)
        }
    }
}