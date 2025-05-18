package main

import "slices"

func main() {

}

func subsets(nums []int) [][]int {
	n := len(nums)
	subset := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(cur int)

	dfs = func(cur int) {
		ans = append(ans, slices.Clone(subset))
		// 枚举数字
		for i := cur; i < n; i++ {
			subset = append(subset, nums[i])
			// 注意这里是 i + 1，不是 cur + 1
			dfs(i + 1)
			subset = subset[:len(subset)-1]
		}
	}

	dfs(0)

	return ans
}
