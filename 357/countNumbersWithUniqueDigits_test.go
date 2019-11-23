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
