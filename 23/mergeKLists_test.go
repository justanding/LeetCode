package mergeKLists_test

import "testing"

type ListNode struct {
    Val int
    Next *ListNode
}

func testMergeKLists(t *testing.T) {

}

/**
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    l := len(lists)
    if l == 0 {
        return nil
    }

    if l==1 {
        return lists[0]
    }

    ll := l / 2
    for i := 0; i < ll; i++ {
        lists[i] = merge(lists[i], lists[l-i-1])
    }

    if l % 2 > 0 {
        ll += 1
    }

    result := mergeKLists(lists[:ll])

    return result
}


func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    result := l1
    if l1.Val > l2.Val {
        result = l2
        l2 = l2.Next
    } else {
        l1 = l1.Next
    }

    curent := result
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            curent.Next = l2
            l2 = l2.Next
        } else {
            curent.Next = l1
            l1 = l1.Next
        }
        curent = curent.Next
    }
    if l1 != nil {
        curent.Next = l1
    } else {
        curent.Next = l2
    }

    return result
}
