package main

import "slices"

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	var dfs func(cur, left int)
	ans := make([][]int, 0)
	path := make([]int, 0)

	dfs = func(cur, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
		}

		if left < 0 || cur >= len(candidates) {
			return
		}

		for i := cur; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, left-candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)

	return ans
}
