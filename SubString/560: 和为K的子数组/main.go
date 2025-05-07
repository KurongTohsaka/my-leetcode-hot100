package main

func main() {

}

func subarraySum(nums []int, k int) int {
	ans := 0
	// 统计前缀和
	preSum := make([]int, len(nums)+1)
	for i, n := range nums {
		preSum[i+1] = preSum[i] + n
	}

	// 通过检索 preSum - k 前缀和的个数来计算子数组数量
	cnt := make(map[int]int, len(preSum))
	for _, sum := range preSum {
		ans += cnt[sum-k]
		cnt[sum]++
	}

	return ans
}
