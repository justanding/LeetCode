package sumNumbers_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	total := 0
	sumNumbersRecu(root, &total, 0)
	return total
}

func sumNumbersRecu(root *TreeNode, total *int, num int) {
	val := num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*total += val
		return
	}

	if root.Left != nil {
		sumNumbersRecu(root.Left, total, val)
	}

	if root.Right != nil {
		sumNumbersRecu(root.Right, total, val)
	}
}
