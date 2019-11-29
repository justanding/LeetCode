package isValidBST_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * 采用二叉树中序排序，得到一个升序slice即是二叉搜索树
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sortSlice := midOrder(root)
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
