package maxProfit_test

import "testing"

func TestMaxProfile(t *testing.T) {
	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	profile := 0

	l := len(prices)
	if l < 2 {
		return profile
	}

	buy := prices[0]
	for i := 1; i < l; i++ {
		if prices[i] < prices[i-1] {
			p := prices[i-1] - buy
			if p > 0 {
				profile += p
			}
			buy = prices[i]
		}
	}

	if prices[l-1]-buy > 0 {
		profile += prices[l-1] - buy
	}

	return profile
}
