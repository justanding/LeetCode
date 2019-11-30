package isBalanced_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 *https://leetcode-cn.com/problems/balanced-binary-tree/
 */
func isBalanced(root *TreeNode) bool {
	_, ok := treeLevel(root)
	return ok
}

func treeLevel(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, leftOk := treeLevel(root.Left)
	right, rightOk := treeLevel(root.Right)
	if leftOk == false || rightOk == false ||
		left > right+1 || left < right-1 {
		return 0, false
	}
	if left < right {
		left = right
	}
	return left + 1, true
}
