package countNumbersWithUniqueDigits_test

import "testing"

func TestCountNumbersWithUniqueDigits(t *testing.T) {
	t.Log(countNumbersWithUniqueDigits(0))
	t.Log(countNumbersWithUniqueDigits(1))
	t.Log(countNumbersWithUniqueDigits(2))
	t.Log(countNumbersWithUniqueDigits(3))
	t.Log(countNumbersWithUniqueDigits(10))
	t.Log(countNumbersWithUniqueDigits(12))
}

/*
 * 采用动态规划思路，
 * 当n=0，return 1
 * 当n=1，return 10
 * 当n=2, return 9*9+f(1)
 * 当n=3, return 9*(9*8)+f(2)
 * 以此类推到n=10
 * 当n>10，肯定有重复，等于n=10的情况
 */
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	ret := 10
	for i, tmp := 1, 9; i < n; i++ {
		tmp *= 10 - i
		ret += tmp
	}
	return ret
}
