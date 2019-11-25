package climbStairs_test

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs(0))
	t.Log(climbStairs(1))
	t.Log(climbStairs(2))
	t.Log(climbStairs(10))
	t.Log(climbStairs(100))
}

/*
 * https://leetcode-cn.com/problems/climbing-stairs/
 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	n2 := 1
	n1 := 2
	for i := 3; i <= n; i++ {
		n2, n1 = n1, n2+n1
	}

	return n1
}
