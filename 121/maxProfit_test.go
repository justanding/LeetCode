package maxProfit_test

/*
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 * 标记以前的最小值，计算最大值
 */
func maxProfit(prices []int) int {
	max := 0
	l := len(prices)
	if l < 2 {
		return max
	}
	for i, minCur := 1, prices[0]; i < l; i++ {
		if prices[i]-minCur > max {
			max = prices[i] - minCur
		}

		if minCur > prices[i] {
			minCur = prices[i]
		}
	}
	return max
}
