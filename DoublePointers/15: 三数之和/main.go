package main

import "slices"

func main() {}

func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	// 用哈希表存储结果，去重
	hashMap := make(map[[3]int]bool)
	slices.Sort(nums)

	for first := 0; first < n; first++ {
		second, third := first+1, n-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				if v := hashMap[[3]int{nums[first], nums[second], nums[third]}]; !v {
					ans = append(ans, []int{nums[first], nums[second], nums[third]})
					hashMap[[3]int{nums[first], nums[second], nums[third]}] = true
				}
				second++
				third--
			} else if sum < 0 {
				second++
			} else {
				third--
			}
		}
	}

	return ans
}
