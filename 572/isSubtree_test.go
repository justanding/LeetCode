package isSubtree_test

import (
    "testing"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func TestIsSubtree(t *testing.T) {

}

func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {
        return false
    }
    if isSub(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t){
       return true
    }
    return false
}

func isSub(s *TreeNode, t *TreeNode) bool {
   if s == nil && t == nil {
       return true
   }
   if s == nil || t == nil || s.Val != t.Val {
       return false
   }

   return isSub(s.Left, t.Left) &&  isSub(s.Right, t.Right)
}