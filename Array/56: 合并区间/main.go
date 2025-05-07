package main

import "slices"

func main() {

}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	ans := make([][]int, 0)

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for _, interval := range intervals {
		n := len(ans)
		// 只分为区间相交、区间不相交两情况
		if n > 0 && interval[0] <= ans[n-1][1] {
			ans[n-1][1] = max(ans[n-1][1], interval[1])
		} else {
			ans = append(ans, interval)
		}
	}

	return ans
}
