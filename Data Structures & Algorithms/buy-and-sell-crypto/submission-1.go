func maxProfit(prices []int) int {
	minBuy := prices[0]
	maxProfit := 0
	for i:=0;i<len(prices);i++{
		buy := prices[i]
		if buy < minBuy{
			minBuy = buy
		}
		profit := buy - minBuy
		if maxProfit < profit {
			maxProfit = profit
		}
	}
	return maxProfit
}
