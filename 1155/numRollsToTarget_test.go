package numRollsToTarget_test

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	t.Log(numRollsToTarget(1, 1, 1))
	t.Log(numRollsToTarget(4, 6, 20))
	t.Log(numRollsToTarget(30, 20, 500))
}

/*
 *dp(n) = dp(n-f)+dp(n-(f-1)).....+dp(n-1)
 */
func numRollsToTarget(d int, f int, target int) int {

	dp := make([]int, d)

	for i := 0; i < d; i++ {
		for j := 1; j <= f; j++ {
			if i-j < 0 {
				break
			}
			dp[i] += dp[i-j]
		}
	}
	return dp[target]
}
