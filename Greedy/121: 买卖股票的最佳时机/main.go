package main

import "math"

func main() {

}

func maxProfit(prices []int) int {
	profit, cost := 0, math.MaxInt // cost 表示买入

	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}

	return profit
}
