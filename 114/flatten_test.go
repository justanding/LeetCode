package flatten_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestFlatten(t *testing.T) {

}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	_, _ = flattenRecur(root)
}

func flattenRecur(root *TreeNode) (*TreeNode, *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}

	var rightHead, tail *TreeNode
	if root.Right != nil {
		rightHead, tail = flattenRecur(root.Right)
	}
	if root.Left != nil {
		leftHead, leftTail := flattenRecur(root.Left)
		root.Right = leftHead
		root.Left = nil
		if rightHead != nil {
			leftTail.Right = rightHead
		} else {
			tail = leftTail
		}
	}

	return root, tail
}
