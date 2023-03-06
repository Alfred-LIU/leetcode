package Math

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min, maxSale := prices[0], 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if (price - min) > maxSale {
			maxSale = price - min
		}
	}

	return maxSale
}
