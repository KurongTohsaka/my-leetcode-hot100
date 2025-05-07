package main

import (
	"math"
	"slices"
)

func main() {

}

func maxSubArray(nums []int) int {
	preSum, minPreSum, ans := 0, 0, math.MinInt

	for _, num := range nums {
		preSum += num
		// 每个子数组都可以表示为前缀和-最小前缀和的形式
		// 前缀和：从数组首部开始计算
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}

	return ans
}

func maxSubArrayDP(nums []int) int {
	// 一维 DP
	// 存储到 i 个元素为止的最大子数组和
	cache := make([]int, len(nums)+1)
	cache[0] = math.MinInt

	for i, num := range nums {
		// 要么选择前者结果进行计算，要么不选从自己开始计算
		cache[i+1] = max(cache[i], 0) + num
	}

	return slices.Max(cache)
}
