package pathSum_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * https://leetcode-cn.com/problems/path-sum-ii/
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	var allPath [][]int
	pathSumRecu(root, sum, &allPath, []int{})
	return allPath
}

func pathSumRecu(root *TreeNode, sum int, allPath *[][]int, path []int) {
	sum -= root.Val
	newPath := make([]int, len(path)+1)
	copy(newPath, path)
	newPath[len(path)] = root.Val
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			*allPath = append(*allPath, newPath)
		}
		return
	}

	if root.Left != nil {
		pathSumRecu(root.Left, sum, allPath, newPath)
	}

	if root.Right != nil {
		pathSumRecu(root.Right, sum, allPath, newPath)
	}
}
