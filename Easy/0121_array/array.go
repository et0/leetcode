package l_e_0121_array

func maxProfit(prices []int) int {
	min, max := 0, 0
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min, max = i, i
		}
		if prices[i] > prices[max] {
			max = i
		}

		newProfit := prices[max] - prices[min]
		if profit < newProfit {
			profit = newProfit
		}
	}
	return profit
}

func Wrapper(prices []int) int {
	return maxProfit(prices)
}
