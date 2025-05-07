package main

func main() {

}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// 统计前缀乘积
	preTimes := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(preTimes) != 0 {
			preTimes = append(preTimes, nums[i]*preTimes[i-1])
		} else {
			preTimes = append(preTimes, nums[i])
		}
	}

	// 统计后缀乘积
	sufTimes := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		if len(sufTimes) != 0 {
			sufTimes = append(sufTimes, nums[i]*sufTimes[n-i-2])
		} else {
			sufTimes = append(sufTimes, nums[i])
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i != 0 && i != n-1 {
			// 将每个乘积分解为 preTimes[i-1] * sufTimes[n-2-i]
			ans[i] = preTimes[i-1] * sufTimes[n-2-i]
		} else if i == 0 {
			ans[0] = sufTimes[n-2]
		} else {
			ans[n-1] = preTimes[n-2]
		}
	}

	return ans
}
