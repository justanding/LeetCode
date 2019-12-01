package sortedListToBST_test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSortedListToBST(t *testing.T) {
	head := &ListNode{}
	tmp := head
	for _, v := range []int{-10, -3, 0, 5, 9} {
		tmp.Next = &ListNode{v, nil}
		tmp = tmp.Next
	}

	t.Log(isValidBST(sortedListToBST(head.Next)))
}

/*
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
 * 递归找到每个部分的中间值
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	mid, right := findMiddleNode(head)
	treeHead := &TreeNode{
		Val:   mid.Val,
		Left:  nil,
		Right: nil,
	}
	if head != mid {
		treeHead.Left = sortedListToBST(head)
	}
	if right != nil {
		treeHead.Right = sortedListToBST(right)
	}

	return treeHead
}

func findMiddleNode(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	pre := head
	mid := head
	for fast := head; fast != nil; {
		if fast.Next == nil {
			break
		}
		pre = mid
		mid = mid.Next
		if fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	if pre != mid {
		pre.Next = nil
	}
	return mid, mid.Next
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sortSlice := midOrder(root)
	fmt.Println(sortSlice)
	for i := 0; i < len(sortSlice)-1; i++ {
		if sortSlice[i] >= sortSlice[i+1] {
			return false
		}
	}
	return true
}

func midOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := midOrder(root.Left)
	left = append(left, root.Val)
	right := midOrder(root.Right)
	return append(left, right...)
}
