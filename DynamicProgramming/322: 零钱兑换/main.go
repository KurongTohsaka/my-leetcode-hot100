package main

import "math"

func main() {

}

func coinChange(coins []int, amount int) int {
	cache := make([]int, amount+1) // 金额 i 由 coins 组成的最小数量
	cache[0] = 0

	for i := 1; i <= amount; i++ {
		cache[i] = math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && cache[i-coin] != math.MaxInt {
				cache[i] = min(cache[i], cache[i-coin]+1)
			}
		}
	}

	if cache[amount] == math.MaxInt {
		return -1
	}

	return cache[amount]
}
