package hasPathSum_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * https://leetcode-cn.com/problems/path-sum/
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return true
		} else {
			return false
		}
	}

	if root.Left != nil {
		flag := hasPathSum(root.Left, sum-root.Val)
		if flag == true {
			return true
		}
	}

	if root.Right != nil {
		flag := hasPathSum(root.Right, sum-root.Val)
		if flag == true {
			return true
		}
	}
	return false
}
